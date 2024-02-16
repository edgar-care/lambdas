package services

import "fmt"

type Symptom struct {
	Name    string `json:"symptom"`
	Present *bool  `json:"present"`
}

func SymptomsToString(symptoms []SessionSymptom) []string {
	var strings = []string{}
	for _, symptom := range symptoms {
		if symptom.Presence == nil {
			strings = append(strings, fmt.Sprintf("?%s", symptom.Name))
		} else if *symptom.Presence == true {
			strings = append(strings, fmt.Sprintf("*%s", symptom.Name))
		} else if *symptom.Presence == false {
			strings = append(strings, fmt.Sprintf("!%s", symptom.Name))
		}
	}
	return strings
}

func pointerToBool(val bool) *bool {
	return &val
}

func StringToSymptoms(strings []string) []SessionSymptom {
	var newSymptoms = []SessionSymptom{}
	for _, symptom := range strings {
		if symptom[0] == '*' {
			newSymptoms = append(newSymptoms, SessionSymptom{Name: symptom[1:], Presence: pointerToBool(true), Duration: nil})
		} else if symptom[0] == '!' {
			newSymptoms = append(newSymptoms, SessionSymptom{Name: symptom[1:], Presence: pointerToBool(false), Duration: nil})
		} else if symptom[0] == '?' {
			newSymptoms = append(newSymptoms, SessionSymptom{Name: symptom[1:], Presence: nil, Duration: nil})
		}
	}
	return newSymptoms
}
