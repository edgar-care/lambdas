package services

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

func GuessQuestion(context []ExamContextItem) (string, []string, bool) {
	if len(context) == 0 {
		return "Pourriez-vous décrire vos symptomes ?", []string{}, false
	}
	if isPresent(context, "maux_de_tetes") != nil {
		if isPresent(context, "vision_trouble") == nil {
			return "Avez vous la vision trouble ?", []string{"vision_trouble"}, false
		}
		if isPresent(context, "fievre") == nil {
			return "Avez vous de la fievre ?", []string{"fievre"}, false
		}
	}
	if isPresent(context, "vision_trouble") != nil {
		if isPresent(context, "maux_de_tetes") == nil {
			return "Avez vous des maux de têtes ?", []string{"maux_de_tetes"}, false
		}
		if isPresent(context, "fievre") == nil {
			return "Avez vous de la fievre ?", []string{"fievre"}, false
		}
	}
	if isPresent(context, "maux_de_ventre") != nil {
		if isPresent(context, "vomissement") == nil {
			return "Avez vous des vomissement ?", []string{"vomissement"}, false
		}
		if isPresent(context, "fievre") == nil {
			return "Avez vous de la fievre ?", []string{"fievre"}, false
		}
	}
	if isPresent(context, "fievre") != nil {
		if isPresent(context, "vomissement") == nil {
			return "Avez vous des vomissement ?", []string{"vomissement"}, false
		}
		if isPresent(context, "vision_trouble") == nil {
			return "Avez vous la vision trouble ?", []string{"vision_trouble"}, false
		}
	}
	return "Comment ca va mon reuf ?", []string{}, true
}
