package utils

import (
	"net/http"
)

func Respond(w http.ResponseWriter, jData []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}
