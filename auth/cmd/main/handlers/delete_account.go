package handlers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/edgar-care/auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/edgar-care/edgarlib/v2/redis"
	"net/http"
	"strconv"
)

func DeleteAccount(w http.ResponseWriter, req *http.Request) {
	accountID := edgarlib.AuthMiddlewareAccount(w, req)
	if accountID.ID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	err := deleteAccountLogic(accountID.ID)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": err.Error(),
		}, 500)
		return
	}

	lib.WriteResponse(w, map[string]string{
		"message": "Account deletion scheduled",
	}, 200)
}

func deleteAccountLogic(ownerID string) error {
	expire := 30 * 24 * 60 * 60
	_, err := redis.SetKey(ownerID+"_delete_request", "pending", &expire)
	if err != nil {
		return fmt.Errorf("failed to set delete request in Redis: %v", err)
	}

	email, accountType, err := edgarlib.getEmailByOwnerID(ownerID)
	if err != nil {
		return fmt.Errorf("failed to get email: %v", err)
	}

	err = edgarlib.sendDeleteConfirmationEmail(email)
	if err != nil {
		return fmt.Errorf("failed to send confirmation email: %v", err)
	}

	for i := 1; i < 30; i++ {
		err = scheduleReminderEmailWithEventBridge(ownerID, email, 30-i)
		if err != nil {
			return fmt.Errorf("failed to schedule reminder email for day %d: %v", i, err)
		}
	}

	err = scheduleFinalAccountDeletionWithEventBridge(ownerID, accountType, 30)
	if err != nil {
		return fmt.Errorf("failed to schedule final account deletion: %v", err)
	}

	return nil
}

func scheduleFinalAccountDeletionWithEventBridge(ownerID, accountType string, days int) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-west-3"),
	}))

	eb := eventbridge.New(sess)

	ruleName := fmt.Sprintf("FinalAccountDeletion-%s", ownerID)

	_, err := eb.PutRule(&eventbridge.PutRuleInput{
		Name:               aws.String(ruleName),
		ScheduleExpression: aws.String(fmt.Sprintf("rate(%d days)", days)),
		State:              aws.String("ENABLED"),
	})
	if err != nil {
		return fmt.Errorf("failed to create EventBridge rule: %v", err)
	}

	lambdaTarget := &eventbridge.Target{
		Id:  aws.String("1"),
		Arn: aws.String("arn:aws:lambda:eu-west-3:146778342232:function:auth"),
		Input: aws.String(fmt.Sprintf(`{
			"ownerID": "%s",
			"accountType": "%s"
		}`, ownerID, accountType)),
	}

	_, err = eb.PutTargets(&eventbridge.PutTargetsInput{
		Rule:    aws.String(ruleName),
		Targets: []*eventbridge.Target{lambdaTarget},
	})
	if err != nil {
		return fmt.Errorf("failed to set EventBridge target: %v", err)
	}

	return nil
}

func scheduleReminderEmailWithEventBridge(ownerID, email string, daysLeft int) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-west-3"),
	}))

	eb := eventbridge.New(sess)

	ruleName := fmt.Sprintf("ReminderEmail-%s-%d-days-left", ownerID, daysLeft)

	_, err := eb.PutRule(&eventbridge.PutRuleInput{
		Name:               aws.String(ruleName),
		ScheduleExpression: aws.String("rate(1 day)"),
		State:              aws.String("ENABLED"),
	})
	if err != nil {
		return fmt.Errorf("failed to create EventBridge rule: %v", err)
	}

	lambdaTarget := &eventbridge.Target{
		Id:  aws.String("1"),
		Arn: aws.String("arn:aws:lambda:eu-west-3:146778342232:function:auth"),
		Input: aws.String(fmt.Sprintf(`{
			"ownerID": "%s",
			"email": "%s",
			"daysLeft": %d
		}`, ownerID, email, daysLeft)),
	}

	_, err = eb.PutTargets(&eventbridge.PutTargetsInput{
		Rule:    aws.String(ruleName),
		Targets: []*eventbridge.Target{lambdaTarget},
	})
	if err != nil {
		return fmt.Errorf("failed to set EventBridge target: %v", err)
	}

	return nil
}

func DeleteAccountFinal(w http.ResponseWriter, req *http.Request) {
	accountID := edgarlib.AuthMiddlewareAccount(w, req)
	if accountID.ID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	status, err := redis.GetKey(accountID.ID + "_delete_request")
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Failed to get delete request status",
		}, 500)
		return
	}

	if status == "pending" {
		_, accountType, err := edgarlib.getEmailByOwnerID(accountID.ID)
		if err != nil {
			lib.WriteResponse(w, map[string]string{
				"message": "Failed to get account type",
			}, 500)
			return
		}

		err = edgarlib.deleteDefinitlyAccount(accountID.ID, accountType)
		if err != nil {
			lib.WriteResponse(w, map[string]string{
				"message": "Failed to delete account",
			}, 500)
			return
		}

		lib.WriteResponse(w, map[string]string{
			"message": "Account successfully deleted",
		}, 200)
	} else {
		lib.WriteResponse(w, map[string]string{
			"message": "Account deletion is no longer pending",
		}, 200)
	}
}

func SendReminderEmail(w http.ResponseWriter, req *http.Request) {
	accountID := edgarlib.AuthMiddlewareAccount(w, req)
	if accountID.ID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	daysLeftStr := req.URL.Query().Get("daysLeft")
	if daysLeftStr == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Missing 'daysLeft' parameter",
		}, 400)
		return
	}

	daysLeft, err := strconv.Atoi(daysLeftStr)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid 'daysLeft' parameter",
		}, 400)
		return
	}

	status, err := redis.GetKey(accountID.ID + "_delete_request")
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Failed to get delete request status",
		}, 500)
		return
	}

	if status == "pending" {
		email, _, err := edgarlib.getEmailByOwnerID(accountID.ID)
		if err != nil {
			lib.WriteResponse(w, map[string]string{
				"message": "Failed to get email for account",
			}, 500)
			return
		}

		err = edgarlib.sendReminderEmail(email, daysLeft)
		if err != nil {
			lib.WriteResponse(w, map[string]string{
				"message": "Failed to send reminder email",
			}, 500)
			return
		}

		lib.WriteResponse(w, map[string]string{
			"message": fmt.Sprintf("Reminder email sent with %d days left", daysLeft),
		}, 200)
	} else {
		lib.WriteResponse(w, map[string]string{
			"message": "Account deletion is no longer pending",
		}, 200)
	}
}
