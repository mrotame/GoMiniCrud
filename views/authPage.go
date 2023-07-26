package views

import (
	"encoding/json"
	"net/http"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
	"github.com/mrotame/GoCrudApi/validators"
)

type Login = struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response = struct {
	Auth_token string                 `json:"auth_token"`
	User_data  map[string]interface{} `json:"user_data"`
}

func AuthPage(w http.ResponseWriter, r *http.Request) {
	var login Login

	if !validators.ValidateBody(w, r.Body, &login) {
		return
	}

	var user models.User
	err := database.GetOne(&user, "Email = ?", login.Email).Error

	if !validators.ValidateModelFound(w, err, "invalid email or password") {
		return
	}

	if !validators.ValidatePassword(w, login.Password, user.Password) {
		return
	}

	auth := models.Auth{
		User_id: user.ID,
	}

	database.Save(&auth)

	user_json, _ := user.As_map()

	response := Response{
		auth.Token,
		user_json,
	}

	jData, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(jData)
}
