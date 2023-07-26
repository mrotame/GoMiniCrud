package userPage

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
	"github.com/mrotame/GoCrudApi/test"
)

func createUser(email string, accessLevel int) models.User {
	user := models.User{
		Name:        "Test User",
		Age:         21,
		AccessLevel: accessLevel,
		Email:       email,
		Password:    "123",
	}
	database.Save(&user)
	return user
}

func createAuth(user models.User) models.Auth {
	auth := models.Auth{
		User_id: user.ID,
	}
	database.Save(&auth)
	return auth
}

func TestDeleteUser_with_auth_token(t *testing.T) {
	test.SetupTestDB(&models.User{}, &models.Auth{})

	var (
		non_deleted_user1 models.User
		deleted_user2     models.User

		user1 = createUser("test@test.com", 0)
		user2 = createUser("test2@test.com", 0)

		authToken = "Bearer " + createAuth(user2).Token

		w = *httptest.NewRecorder()
		r = httptest.NewRequest(http.MethodDelete, "/user", nil)
	)
	r.Header.Add("Authorization", "Bearer "+authToken)
	UserPage(&w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Error(fmt.Sprintf("Error, Delete method returned status %v instead %v", w.Result(), http.StatusOK))
	}

	err := database.GetOne(&deleted_user2, user2.ID).Error
	if err == nil {
		t.Error(fmt.Sprintf("Error, user %v (ID %v) was found in the database. user %v (ID %v) is supposed to be deleted.", deleted_user2.Name, deleted_user2.ID, user2.Name, user2.ID))
	}

	err = database.GetOne(&non_deleted_user1, user1.ID).Error
	if err != nil {
		t.Error(fmt.Sprintf("Error, user %v (ID %v) was not found in the database. It was not supposed to be deleted.", non_deleted_user1.Name, non_deleted_user1.ID))
	}

	test.TeardownTestDB()
}

func TestDeleteUser_with_id(t *testing.T) {
	test.SetupTestDB(&models.User{}, &models.Auth{})

	var (
		user1 = createUser("test@test.com", 0)
		user2 = createUser("test2@test.com", 1)
		user3 = createUser("test3@test.com", 2)

		authToken_user2 = "Bearer " + createAuth(user2).Token
		authToken_user3 = "Bearer " + createAuth(user3).Token

		mux_vars1 = map[string]string{
			"id": fmt.Sprint(user3.ID),
		}
		mux_vars2 = map[string]string{
			"id": fmt.Sprint(user2.ID),
		}

		w1 = *httptest.NewRecorder()
		w2 = *httptest.NewRecorder()

		r1 = httptest.NewRequest(http.MethodDelete, "/user", nil)
		r2 = httptest.NewRequest(http.MethodDelete, "/user", nil)
	)

	r1 = mux.SetURLVars(r1, mux_vars1)
	r2 = mux.SetURLVars(r2, mux_vars2)

	r1.Header.Add("Authorization", "Bearer "+authToken_user2)
	r2.Header.Add("Authorization", "Bearer "+authToken_user3)

	UserPage(&w1, r1)
	UserPage(&w2, r2)

	if w1.Result().StatusCode == http.StatusOK {
		t.Error(fmt.Sprintf("Error, request to delete a user with higher access level returned status %v", w1.Result().StatusCode))
	}

	if database.GetOne(&models.User{}, user3.ID).Error != nil {
		t.Error(fmt.Sprintf("Error. User %v with access level %v was able to delete user %v with higher access level %v", user2.Name, user2.AccessLevel, user3.Name, user3.AccessLevel))
	}

	if w2.Result().StatusCode != http.StatusOK {
		t.Error(fmt.Sprintf("Error, request to delete a user with lower access level did not returned status %v", http.StatusOK))
	}

	err := database.GetOne(&models.User{}, user2.ID).Error
	if err == nil {
		t.Error(fmt.Sprintf("Error. user %v with access level %v was not able to delete user %v with lower access level %v", user3.Name, user3.AccessLevel, user2.Name, user2.AccessLevel))
	}
	err = database.GetOne(&models.User{}, user1.ID).Error
	if err != nil {
		t.Error(fmt.Sprintf("Error. user %v was deleted without a request", user1.Name))
	}

	test.TeardownTestDB()
}

func TestDeleteUser_without_auth_token(t *testing.T) {
	test.SetupTestDB(&models.User{})

	var (
		user, _ = GetTestUser()

		mux_vars = map[string]string{
			"id": fmt.Sprint(user.ID),
		}

		w = *httptest.NewRecorder()
		r = httptest.NewRequest(http.MethodDelete, "/user", nil)
	)

	r = mux.SetURLVars(r, mux_vars)

	UserPage(&w, r)

	if w.Result().StatusCode == http.StatusOK {
		t.Error("Error. Delete request without Auth token returned", w.Result().StatusCode)
	}

	if database.GetOne(&models.User{}, user.ID).Error != nil {
		t.Error("Error, delete request without auth token deleted the user")
	}

	test.TeardownTestDB()
}
