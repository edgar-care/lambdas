package handlers

import (
	"net/http"

	"github.com/edgar-care/Lambda_test/cmd/main/lib"
)


func AuthUser(w http.ResponseWriter, req *http.Request) {


	token := lib.GetAuthenticatedUser(w, req)

	lib.WriteResponse(w, map[string]interface{}{
		"token": token,
	}, 201)
}