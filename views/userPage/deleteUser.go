package userPage

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
	"github.com/mrotame/GoCrudApi/utils"
	"github.com/mrotame/GoCrudApi/validators"
)

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user models.User
	id, idInRequest := vars["id"]
	authToken := utils.GetTokenFromRequest(r)

	if !validators.ValidateAuthToken(w, authToken) {
		return
	}
	if idInRequest {
		_deleteWithId(w, r, id, authToken)
	} else {
		_deleteWithAuthToken(w, r, authToken)
	}

	database.Delete(&user)
}

func _deleteWithId(w http.ResponseWriter, _ *http.Request, id string, authToken string) {
	var user models.User
	var requester models.User

	err := database.GetOne(&user, id).Error
	if !validators.ValidateModelFound(w, err, "User not found") {
		return
	}

	models.GetUserBy_AuthToken(&requester, authToken)

	if !validators.ValidateHigherAccessLevel(w, uint(requester.AccessLevel), uint(user.AccessLevel)) {
		return
	}

	database.Delete(&user)
	jData, _ := user.As_json()
	utils.Respond(w, jData)
}

func _deleteWithAuthToken(w http.ResponseWriter, _ *http.Request, authToken string) {
	var user models.User
	models.GetUserBy_AuthToken(&user, authToken)
	database.Delete(&user)
	jData, _ := user.As_json()
	utils.Respond(w, jData)
}
