package services

import (
	"bytes"
	"encoding/json"
	"github.com/edgar-care/diagnostic/cmd/main/lib"
	"net/http"
	"os"
)

type examRequestBody struct {
	Context []Symptom `json:"context"`
}

type examResponseBody struct {
	Context  []Symptom `json:"context"`
	Done     bool      `json:"done"`
	Question string    `json:"question"`
	Symptoms []string  `json:"symptoms"`
}

func CallExam(context []Symptom) examResponseBody {
	var rBody = examRequestBody{
		Context: context,
	}

	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(rBody)
	lib.CheckError(err)

	resp, err := http.Post(os.Getenv("EXAM_URL"), "application/json", buf)
	lib.CheckError(err)

	var respBody examResponseBody
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	lib.CheckError(err)

	return respBody
}
