package userPage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
	"github.com/mrotame/GoCrudApi/test"
)

func TestPutUser_byAuthToken(t *testing.T) {
	test.SetupTestDB(&models.User{}, models.Auth{})
	var (
		user_from_db  models.User
		user, auth1   = GetTestUser()
		user_new_name = "New Test Name"
		user_old_name = user.Name

		w1 = *httptest.NewRecorder()
		w2 = *httptest.NewRecorder()

		r1 = httptest.NewRequest(http.MethodPut, "/user", nil)
	)
	patch_body_map := map[string]interface{}{
		"name":  user_new_name,
		"age":   user.Age,
		"email": user.Email,
	}

	patch_request_body_json, _ := json.Marshal(patch_body_map)
	ioreader := bytes.NewBuffer(patch_request_body_json)

	r2 := httptest.NewRequest(http.MethodPut, "/user", ioreader)

	r2.Header.Add("Authorization", "Bearer "+auth1.Token)

	UserPage(&w1, r1)
	UserPage(&w2, r2)

	var w2_result_body map[string]interface{}

	json.NewDecoder(w2.Body).Decode(&w2_result_body)

	if w1.Result().StatusCode == http.StatusOK {
		t.Error("Error, request without Auth token returned", http.StatusOK)
	}

	if w2.Result().StatusCode != http.StatusOK {
		t.Error(fmt.Sprintf("Put request from authorized user %v received status %v instead status %v", user.Name, w1.Result().StatusCode, http.StatusOK))
	}

	if w2_result_body["name"] != user_new_name {
		t.Error(fmt.Sprintf("Error. Method Put responded user (originally named as `%v`) as name `%v` instead new name `%v`", user_old_name, w2_result_body["name"], user_new_name))
	}

	database.GetOne(&user_from_db, user.ID)
	if user_from_db.Name != user_new_name {
		t.Error(fmt.Sprintf("Error. User originally named as `%v` has name `%v` instead new one `%v`", user_old_name, w2_result_body["name"], user_new_name))
	}

	test.TeardownTestDB()
}

func TestPutUser_by_id(t *testing.T) {
	test.SetupTestDB(&models.User{}, models.Auth{})
	var (
		user1_from_db models.User
		user2_from_db models.User
		user3_from_db models.User

		user1, _     = GetTestUser()
		user2, auth2 = GetTestUser()
		user3, auth3 = GetTestUser()

		user_new_name = "New Test Name"
		user_old_name = user1.Name

		w1 = *httptest.NewRecorder()
		w2 = *httptest.NewRecorder()
	)
	user3.AccessLevel = 2
	database.Save(&user3)

	patch_body_map := map[string]interface{}{
		"name":  user_new_name,
		"age":   user2.Age,
		"email": user2.Email,
	}

	patch_request_body_json, _ := json.Marshal(patch_body_map)
	ioreader := bytes.NewBuffer(patch_request_body_json)

	r1 := httptest.NewRequest(http.MethodPut, "/user", ioreader)
	r1.Header.Add("Authorization", "Bearer "+auth2.Token)
	r1 = mux.SetURLVars(r1, map[string]string{"id": fmt.Sprint(user3.ID)})

	ioreader = bytes.NewBuffer(patch_request_body_json)
	r2 := httptest.NewRequest(http.MethodPut, "/user", ioreader)
	r2.Header.Add("Authorization", "Bearer "+auth3.Token)
	r2 = mux.SetURLVars(r2, map[string]string{"id": fmt.Sprint(user2.ID)})

	UserPage(&w1, r1)
	UserPage(&w2, r2)

	if w1.Result().StatusCode == http.StatusOK {
		t.Error(fmt.Sprintf("Error. request from user `%v` with lower access level `%v` to modify data from user `%v` with higher access level %v returned status %v", user2.Name, user2.AccessLevel, user3.Name, user3.AccessLevel, w1.Result().StatusCode))
	}

	database.GetOne(&user3_from_db, user3.ID)
	if user3_from_db.Name == user_new_name {
		t.Error(fmt.Sprintf("Error. user `%v` with lower access level `%v` was able to modify data from user `%v` with higher access level %v", user2.Name, user2.AccessLevel, user3.Name, user3.AccessLevel))
	}
	if w2.Result().StatusCode != http.StatusOK {
		t.Error(fmt.Sprintf("Error. request from user `%v` with higher access level `%v` to modify data from user `%v` with lower access level %v returned status %v instead status %v", user3.Name, user3.AccessLevel, user2.Name, user2.AccessLevel, w2.Result().StatusCode, http.StatusOK))
	}

	database.GetOne(&user2_from_db, user2.ID)
	if user2_from_db.Name != user_new_name {
		t.Error(fmt.Sprintf("Error. user `%v` with higher access level `%v` failed to modify data from user `%v` with lower access level %v", user3.Name, user3.AccessLevel, user2.Name, user2.AccessLevel))
	}

	database.GetOne(&user1_from_db, user1.ID)
	if user1_from_db.Name == user_new_name {
		t.Error(fmt.Sprintf("Error. User originally named as `%v` had his name changed to `%v` without any direct request", user_old_name, user_new_name))
	}

	test.TeardownTestDB()
}
