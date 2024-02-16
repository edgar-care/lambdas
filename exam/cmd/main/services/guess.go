package services

import (
	"fmt"
	"sort"
)

type diseaseCoverage struct {
	disease           string
	coverage          int
	present           int
	absent            int
	potentialQuestion string
}

type ByCoverage []diseaseCoverage

func (a ByCoverage) Len() int           { return len(a) }
func (a ByCoverage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCoverage) Less(i, j int) bool { return a[i].coverage > a[j].coverage }

func findInContext(context []ExamContextItem, symptom string) *ExamContextItem {
	for _, item := range context {
		if item.Name == symptom {
			return &item
		}
	}
	return nil
}

func isPresent(context []ExamContextItem, symptom string) *bool {
	item := findInContext(context, symptom)
	if item != nil {
		return item.Presence
	}
	return nil
}

func calculCoverage(context []ExamContextItem, disease Disease) diseaseCoverage {
	var coverage int
	var present int
	var absent int
	var potentialQuestionSymptom string
	total := len(disease.Symptoms)

	for _, symptom := range disease.Symptoms {
		presence := isPresent(context, symptom)
		if presence != nil {
			coverage += 1
			if *presence == true {
				present += 1
			} else {
				absent += 1
			}
		} else {
			potentialQuestionSymptom = symptom
		}
	}
	return diseaseCoverage{disease: disease.Code, coverage: coverage * 100 / total, present: present * 100 / total, absent: absent * 100 / total, potentialQuestion: potentialQuestionSymptom}
}

func getTheQuestion(symptomName string) string {
	//	var symptoms []Symptom
	fmt.Println("hey")
	symptoms, _ := GetSymptoms()
	fmt.Println(symptoms)
	for _, symptom := range symptoms {
		if symptomName == symptom.Name {
			fmt.Println(symptom.Question)
			return symptom.Question
		}
	}
	return symptomName
}

func GuessQuestion(context []ExamContextItem) (string, []string, bool) {
	diseases, _ := GetDiseases()
	//symptoms := getPossibleSymptoms()
	mapped := make([]diseaseCoverage, len(diseases))
	for i, e := range diseases {
		mapped[i] = calculCoverage(context, e)
	}
	//fmt.Println(mapped)
	if len(context) == 0 {
		return "Pourriez-vous décrire vos symptomes ?", []string{}, false
	}

	sort.Sort(ByCoverage(mapped))
	//fmt.Println(mapped)

	for _, disease := range mapped {
		if disease.absent >= 40 {
			continue
		}
		if disease.present >= 70 {
			return "", []string{}, true
		}
		return getTheQuestion(disease.potentialQuestion), []string{disease.potentialQuestion}, false
	}
	return "", []string{}, true
}
