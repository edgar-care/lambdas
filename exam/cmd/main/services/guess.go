package services

import (
	"sort"
)

type diseaseCoverage struct {
	coverage int
	present  int
	absent   int
}

type ByCoverage []diseaseCoverage

func (a ByCoverage) Len() int           { return len(a) }
func (a ByCoverage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCoverage) Less(i, j int) bool { return a[i].coverage > a[j].coverage }

func findInContext(context []ExamContextItem, symptom string) *ExamContextItem {
	for _, item := range context {
		if item.Symptom == symptom {
			return &item
		}
	}
	return nil
}

func isPresent(context []ExamContextItem, symptom string) *bool {
	item := findInContext(context, symptom)
	if item != nil {
		return item.Present
	}
	return nil
}

func calculCoverage(context []ExamContextItem, disease Disease) diseaseCoverage {
	var coverage int
	var present int
	var absent int
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
		}
	}
	return diseaseCoverage{coverage: coverage * 100 / total, present: present * 100 / total, absent: absent * 100 / total}
}

func GuessQuestion(context []ExamContextItem) (string, []string, bool) {
	diseases, _ := GetDiseases()

	mapped := make([]diseaseCoverage, len(diseases))
	for i, e := range diseases {
		mapped[i] = calculCoverage(context, e)
	}

	if len(context) == 0 {
		return "Pourriez-vous décrire vos symptomes ?", []string{}, false
	}

	sort.Sort(ByCoverage(mapped))

	for _, disease := range mapped {
		if disease.absent >= 40 {
			continue
		}
		if disease.present >= 70 {
			return "", []string{}, true
		}
		return "next question", []string{"question"}, false
	}
	return "", []string{}, true
}
