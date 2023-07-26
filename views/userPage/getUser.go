package userPage

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
	"github.com/mrotame/GoCrudApi/utils"
	"github.com/mrotame/GoCrudApi/validators"
)

func getUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, idInRequest := vars["id"]

	authToken := utils.GetTokenFromRequest(r)

	if !validators.ValidateAuthToken(w, authToken) {
		return
	}

	if idInRequest {
		_getUserWithId(w, id, authToken)
	} else {
		_getUserWithAuthToken(w, authToken)
	}
}

func _getUserWithId(w http.ResponseWriter, requested_id string, requester_token string) {
	var requester_user models.User
	var requested_user models.User

	models.GetUserBy_AuthToken(&requester_user, requester_token)

	err := database.GetOne(&requested_user, requested_id).Error

	if !validators.ValidateModelFound(w, err, "User not found") {
		return
	}

	if !validators.ValidateHigherAccessLevel(w, uint(requester_user.AccessLevel), uint(requested_user.AccessLevel)) {
		return
	}

	jData, _ := requested_user.As_json()
	utils.Respond(w, jData)
}

func _getUserWithAuthToken(w http.ResponseWriter, authToken string) {
	var user models.User
	err := models.GetUserBy_AuthToken(&user, authToken).Error
	if !validators.ValidateModelFound(w, err, "User not found") {
		return
	}

	jData, _ := user.As_json()
	utils.Respond(w, jData)
}
