package userPage

import (
	"net/http"
)

var viewMethods = map[string]func(http.ResponseWriter, *http.Request){
	http.MethodGet:    getUser,
	http.MethodPost:   postUser,
	http.MethodPut:    putUser,
	http.MethodDelete: deleteUser,
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	viewMethods[r.Method](w, r)
}
