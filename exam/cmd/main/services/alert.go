package services

func isAlertPresent(context []ExamContextItem, symptom string) bool {
	for _, e := range context {
		if e.Present != nil {
			if e.Symptom == symptom && *e.Present == true {
				return true
			}
		}
	}
	return false
}

func coverAlert(context []ExamContextItem, alert Alert) string {
	present := true
	for _, symptom := range alert.Symptoms {
		presence := isAlertPresent(context, symptom)
		if presence == false {
			present = false
		}
	}
	if present == true {
		return alert.ID
	} else {
		return ""
	}
}

func CheckAlerts(context []ExamContextItem) []string {
	//alerts, _ := GetAlerts()
	alerts := GetAlertsHotFix()
	var present []string
	for _, alert := range alerts {
		tmp := coverAlert(context, alert)
		if tmp != "" {
			present = append(present, tmp)
		}
	}
	return present
}
