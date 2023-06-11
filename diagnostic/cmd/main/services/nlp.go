package services

import (
	"bytes"
	"encoding/json"
	"github.com/edgar-care/diagnostic/cmd/main/lib"
	"net/http"
	"os"
)

type nlpRequestBody struct {
	Input    string   `json:"input"`
	Symptoms []string `json:"symptoms"`
}

type nlpResponseBody struct {
	Context []Symptom `json:"context"`
}

func CallNlp(sentence string, symptoms []string) nlpResponseBody {
	var rBody = nlpRequestBody{
		Input:    sentence,
		Symptoms: symptoms,
	}

	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(rBody)
	lib.CheckError(err)

	resp, err := http.Post(os.Getenv("NLP_URL"), "application/json", buf)
	lib.CheckError(err)

	var respBody nlpResponseBody
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	lib.CheckError(err)

	return respBody
}
