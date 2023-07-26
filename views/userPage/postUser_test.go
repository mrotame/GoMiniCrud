package userPage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
	"github.com/mrotame/GoCrudApi/test"
)

var mockedGetRequestMap = map[string]interface{}{
	"name":        "Test User",
	"age":         21,
	"accessLevel": 1,
	"email":       "test@test",
	"password":    "123456",
}

func TestPostUser_send_correct_request_and_check_user_in_db(t *testing.T) {
	test.SetupTestDB(&models.User{})

	var mockedGetJson, _ = json.Marshal(mockedGetRequestMap)
	var mockedGetBody = bytes.NewReader(mockedGetJson)
	var user models.User
	var w = *httptest.NewRecorder()
	var r, _ = http.NewRequest(http.MethodPost, "/user", mockedGetBody)

	UserPage(&w, r)

	err := database.GetOne(&user, "email = ?", "test@test").Error

	if err != nil || user.ID == 0 {
		t.Error(fmt.Sprintf("Error, user creation was requested to user view, but user was not found in database"))
	}

	if w.Result().StatusCode != http.StatusCreated {
		t.Error(fmt.Sprintf("Error, user was created but view returned status `%v` instead `%v`", w.Result().StatusCode, http.StatusCreated))
	}

	test.TeardownTestDB()
}

func TestPostUser_send_request_with_invalid_email_and_check_response(t *testing.T) {
	test.SetupTestDB(&models.User{})

	var _mockedGetRequestMap = mockedGetRequestMap
	_mockedGetRequestMap["email"] = "test"
	var mockedGetJson, _ = json.Marshal(_mockedGetRequestMap)
	var mockedGetBody = bytes.NewReader(mockedGetJson)
	var w = *httptest.NewRecorder()
	var r, _ = http.NewRequest(http.MethodPost, "/user", mockedGetBody)

	UserPage(&w, r)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Error(fmt.Sprintf("user creation request with invalid email was sent, but got status `%v` instead `%v`", w.Result().StatusCode, http.StatusBadRequest))
	}

	test.TeardownTestDB()
}

func TestPostUser_send_request_with_missing_required_data_and_check_response(t *testing.T) {
	test.SetupTestDB(&models.User{})

	var _mockedGetRequestMap = mockedGetRequestMap
	delete(_mockedGetRequestMap, "email")
	var mockedGetJson, _ = json.Marshal(_mockedGetRequestMap)
	var mockedGetBody = bytes.NewReader(mockedGetJson)
	var w = *httptest.NewRecorder()
	var r, _ = http.NewRequest(http.MethodPost, "/user", mockedGetBody)

	UserPage(&w, r)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Error(fmt.Sprintf("user creation request with missing email was sent, but got status `%v` instead `%v`", w.Result().StatusCode, http.StatusBadRequest))
	}

	test.TeardownTestDB()
}
