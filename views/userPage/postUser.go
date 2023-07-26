package userPage

import (
	"encoding/json"
	"net/http"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
	"github.com/mrotame/GoCrudApi/validators"
)

func postUser(w http.ResponseWriter, r *http.Request) {
	var p models.User

	err := json.NewDecoder(r.Body).Decode(&p)
	if !validators.ValidateJsonDecode(w, err) {
		return
	}

	if !validators.ValidateEmail(w, p.Email) {
		return
	}

	err = database.Save(&p).Error

	if !validators.ValidateSavedModel(w, err) {
		return
	}

	jData, _ := p.As_json()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jData)
}
