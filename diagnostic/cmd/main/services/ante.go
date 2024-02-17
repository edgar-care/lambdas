package services

func CheckAnteDiseaseInSymptoms(session Session) string {
	var question string

	//for _, anteId := range session.AnteDiseases {
	//	ante := getAnteDiseaseById(anteId)
	//	if ante.StillRelevant == true && len(Symptoms) > 0 {
	//		for _, anteSymptomId := ante.Symptoms {
	//			anteSymptom := getSymptomById(anteSymptomID)
	//			question = "Ressentez-vous " + anteSymptom.Code + " plus intensément récemment ?"
	//			for _, sessionSymptom := range session.Symptoms {
	//				if anteSymptom.Code == sessionSymptom.Name {
	//					question = ""
	//				}
	//			}
	//			if question != "" {
	//				return question
	//			}
	//		}
	//	}
	//}
	return question
}
