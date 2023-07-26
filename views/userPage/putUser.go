package userPage

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
	"github.com/mrotame/GoCrudApi/utils"
	"github.com/mrotame/GoCrudApi/validators"
)

func putUser(w http.ResponseWriter, r *http.Request) {
	var userUpdated models.UserUpdate

	token := utils.GetTokenFromRequest(r)
	vars := mux.Vars(r)
	body, _ := io.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	id, idInRequest := vars["id"]

	if !validators.ValidateBody(w, bytes.NewBuffer(body), &userUpdated) {
		return
	}

	if idInRequest {
		_updateUserWithId(w, r, body, id, token)
	} else {
		_updateUserWithAuthToken(w, r, body, token)
	}

}

func _updateUserWithId(w http.ResponseWriter, r *http.Request, body []byte, user_id string, requester_token string) {
	var user models.User
	var requester models.User

	err := database.GetOne(&user, user_id).Error
	_ = models.GetUserBy_AuthToken(&requester, requester_token)

	if !validators.ValidateModelFound(w, err, "User not found") {
		return
	}

	if !validators.ValidateHigherAccessLevel(w, uint(requester.AccessLevel), uint(user.AccessLevel)) {
		return
	}

	_updateModel(w, r, body, &user)

	jData, _ := user.As_json()
	utils.Respond(w, jData)
}

func _updateUserWithAuthToken(w http.ResponseWriter, r *http.Request, body []byte, requester_token string) {
	var user models.User
	err := models.GetUserBy_AuthToken(&user, requester_token).Error

	if !validators.ValidateModelFound(w, err, "User not found") {
		return
	}

	_updateModel(w, r, body, &user)

	jData, _ := user.As_json()
	utils.Respond(w, jData)
}

func _updateModel(w http.ResponseWriter, r *http.Request, body []byte, user *models.User) {
	var userUpdatedMap map[string]interface{}
	var err error

	r.Body = io.NopCloser(bytes.NewBuffer(body))
	err = json.NewDecoder(r.Body).Decode(&userUpdatedMap)
	if !validators.GenericValidator(w, err, http.StatusInternalServerError) {
		return
	}

	err = database.Update(&user, userUpdatedMap).Error

	if !validators.GenericValidator(w, err, http.StatusInternalServerError) {
		return
	}
}
