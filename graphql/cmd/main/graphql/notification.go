package graphql

import (
	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/edgar-care/graphql/cmd/main/lib"
	"github.com/graph-gophers/graphql-go"
)

type notificationResolver struct {
	p *models.Notification
}

func (u *notificationResolver) ID() graphql.ID {
	return graphql.ID(u.p.ID.Hex())
}

func (u *notificationResolver) Token() string {
	return u.p.Token
}


func (u *notificationResolver) Title() string {
	return u.p.Title
}

func (u *notificationResolver) Message() string {
	return u.p.Message
}

func resolverFromNotification(p *models.Notification) notificationResolver {
	return notificationResolver{p: p}
}

func (*Resolver) GetNotifications() (*[]*notificationResolver, error) {
	notifications, err := db.GetNotifications()
	lib.CheckError(err)

	var entities []*notificationResolver
	for i := range *notifications {
		resolv := resolverFromNotification(&(*notifications)[i])
		entities = append(entities, &resolv)
	}
	return &entities, nil
}

func (*Resolver) GetNotificationById(args struct{ Id string }) (*notificationResolver, error) {
	notification, err := db.GetNotificationByID(args.Id)
	lib.CheckError(err)

	resolv := resolverFromNotification(notification)
	return &resolv, nil
}

func (*Resolver) CreateNotification(notification models.NotificationCreateInput) (*notificationResolver, error) {
	entity, err := db.InsertNotification(&notification)
	lib.CheckError(err)

	return &notificationResolver{entity}, nil
}

func (*Resolver) UpdateNotification(input models.NotificationUpdateInput) (*notificationResolver, error) {
	res, err := db.UpdateNotification(&input)
	lib.CheckError(err)

	return &notificationResolver{res}, nil
}

func (*Resolver) DeleteNotification(args struct{ Id string }) (*bool, error) {
	result, err := db.DeleteNotification(args.Id)
	lib.CheckError(err)

	return &result, err
}