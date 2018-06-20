package rest

import (
	"net/http"
)

func HandlePostMessage(w http.ResponseWriter, r *http.Request) (err error){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return err
}

