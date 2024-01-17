package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/edgar-care/edgarlib"
)

type examRequestBody struct {
	Context []Symptom `json:"context"`
}

type examResponseBody struct {
	Context  []Symptom `json:"context"`
	Done     bool      `json:"done"`
	Question string    `json:"question"`
	Symptoms []string  `json:"symptoms"`
	Alert    []string  `json:"alert"`
}

func CallExam(context []Symptom) examResponseBody {
	var rBody = examRequestBody{
		Context: context,
	}

	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(rBody)
	edgarlib.CheckError(err)

	resp, err := http.Post(os.Getenv("EXAM_URL"), "application/json", buf)
	edgarlib.CheckError(err)

	var respBody examResponseBody
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	edgarlib.CheckError(err)

	return respBody
}
