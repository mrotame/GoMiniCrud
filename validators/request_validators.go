package validators

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
	"github.com/mrotame/GoCrudApi/utils"
)

func ValidateAuthToken(w http.ResponseWriter, token string) bool {
	var authModel models.Auth
	err := database.GetOne(&authModel, "token = ?", token).Error

	if !ValidateModelFound(w, err, "Auth token invalid or expired") {
		return false
	}

	if !authModel.IsValid() {
		http.Error(w, "Auth token invalid or expired", http.StatusNotFound)
		return false
	}
	return true
}

func ValidatePassword(w http.ResponseWriter, password string, hashedPassword string) bool {

	if !utils.CheckPasswordHash(password, hashedPassword) {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return false
	}
	return true
}

func ValidateIdInRequest(w http.ResponseWriter, r *http.Request, keyName string) bool {
	vars := mux.Vars(r)
	_, id_in_request := vars[keyName]
	if !id_in_request {
		http.Error(w, fmt.Sprintf("Error, %s is a required field", keyName), http.StatusBadRequest)
		return false
	}
	return true
}

func ValidateBody(w http.ResponseWriter, requestBody io.Reader, model interface{}) bool {
	err := json.NewDecoder(requestBody).Decode(model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

func ValidateHigherAccessLevel(w http.ResponseWriter, requester_accessLevel uint, requested_accessLevel uint) bool {
	if requester_accessLevel <= requested_accessLevel {
		http.Error(w, "Requester has no access to requested data", http.StatusUnauthorized)
		return false
	}
	return true
}
