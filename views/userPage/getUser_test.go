package userPage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
	"github.com/mrotame/GoCrudApi/test"
)

func TestGetUser_from_auth_token(t *testing.T) {
	test.SetupTestDB(&models.User{}, models.Auth{})

	w := *httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/user", nil)

	_, auth := GetTestUser()

	authToken := "Bearer " + auth.Token
	r.Header.Add("Authorization", authToken)

	UserPage(&w, r)
	var bodyMap map[string]interface{}
	json.NewDecoder(w.Body).Decode(&bodyMap)

	if w.Result().StatusCode != http.StatusOK {
		t.Error(fmt.Sprintf("Get method from user's view returned status `%v` instead `%v`", w.Result().StatusCode, http.StatusOK))
	}

	if _, keyInMap := bodyMap["ID"]; !keyInMap {
		t.Error("Error, user ID not in response body")
	}

	test.TeardownTestDB()
}

func TestGetUser_from_id(t *testing.T) {
	test.SetupTestDB(&models.User{}, models.Auth{})

	var (
		user1, auth1 = GetTestUser()
		user2, auth2 = GetTestUser()

		mux_vars1 = map[string]string{
			"id": fmt.Sprint(user1.ID),
		}
		mux_vars2 = map[string]string{
			"id": fmt.Sprint(user2.ID),
		}

		w1 = *httptest.NewRecorder()
		w2 = *httptest.NewRecorder()

		r1 = httptest.NewRequest(http.MethodGet, "/user", nil)
		r2 = httptest.NewRequest(http.MethodGet, "/user", nil)

		expected_response, _ = user1.As_map()
	)

	r1 = mux.SetURLVars(r1, mux_vars1)
	r2 = mux.SetURLVars(r2, mux_vars2)

	r1.Header.Add("Authorization", "Bearer"+auth2.Token)
	r2.Header.Add("Authorization", "Bearer"+auth1.Token)

	database.Update(&user2, map[string]interface{}{"AccessLevel": 1})

	UserPage(&w1, r1)
	UserPage(&w2, r2)

	if w1.Result().StatusCode != http.StatusOK {
		t.Error(fmt.Sprintf("Error. user %v with access level %v failed to get data from user %v with access level %v", user2.Name, user2.AccessLevel, user1.Name, user1.AccessLevel))
	}

	var bodyMap map[string]interface{}
	json.NewDecoder(w1.Body).Decode(&bodyMap)

	delete(bodyMap, "CreatedAt")
	delete(bodyMap, "UpdatedAt")
	delete(expected_response, "CreatedAt")
	delete(expected_response, "UpdatedAt")

	if !reflect.DeepEqual(bodyMap, expected_response) {
		t.Error(fmt.Sprint(
			"Error, received response body is different than expected. \n \n",
			"Received: \n",
			bodyMap, "\n \n",
			"Expected: \n",
			expected_response,
		))
	}

	if w2.Result().StatusCode == http.StatusOK {
		t.Error(fmt.Sprintf("Error. user %v with access level %v got access to data from user %v with access level %v", user1.Name, user1.AccessLevel, user2.Name, user2.AccessLevel))
	}

	test.TeardownTestDB()
}
