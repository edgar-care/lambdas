package services

import "fmt"

type Symptom struct {
	Name    string `json:"symptom"`
	Present *bool  `json:"present"`
}

func SymptomsToString(symptoms []Symptom) []string {
	var strings = []string{}
	for _, symptom := range symptoms {
		if symptom.Present == nil {
			strings = append(strings, fmt.Sprintf("?%s", symptom.Name))
		} else if *symptom.Present == true {
			strings = append(strings, fmt.Sprintf("*%s", symptom.Name))
		} else if *symptom.Present == false {
			strings = append(strings, fmt.Sprintf("!%s", symptom.Name))
		}
	}
	return strings
}

func pointerToBool(val bool) *bool {
	return &val
}

func StringToSymptoms(strings []string) []Symptom {
	var newSymptoms = []Symptom{}
	for _, symptom := range strings {
		if symptom[0] == '*' {
			newSymptoms = append(newSymptoms, Symptom{Name: symptom[1:], Present: pointerToBool(true)})
		} else if symptom[0] == '!' {
			newSymptoms = append(newSymptoms, Symptom{Name: symptom[1:], Present: pointerToBool(false)})
		} else if symptom[0] == '?' {
			newSymptoms = append(newSymptoms, Symptom{Name: symptom[1:], Present: nil})
		}
	}
	return newSymptoms
}
