package validators

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
	"github.com/mrotame/GoCrudApi/test"
	"github.com/mrotame/GoCrudApi/utils"
)

type TestModel struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `gorm:"not null; default:null" json:"age"`
}

func TestValidateAuthToken(t *testing.T) {
	test.SetupTestDB(&models.Auth{})
	w := *httptest.NewRecorder()
	auth := models.Auth{
		User_id: 1,
	}

	database.Save(&auth)

	if !ValidateAuthToken(&w, auth.Token) {
		t.Error("Error, Auth Token validation returned false for valid token")
	}

	if ValidateAuthToken(&w, "123") {
		t.Error("Error, Auth token validation returned true for invalid token")
	}

	database.Delete(&auth)
	if ValidateAuthToken(&w, auth.Token) {
		t.Error("Error, Auth token validation returned true for valid but deleted token")
	}

	test.TeardownTestDB()
}

func TestValidatePassword(t *testing.T) {
	w := *httptest.NewRecorder()
	password := "Testingpassword123"
	hashed_password, _ := utils.HashPassword(password)

	if !ValidatePassword(&w, password, hashed_password) {
		t.Error("Error, valid password failed when compared to hash")
	}

	if ValidatePassword(&w, "testpassword123", hashed_password) {
		t.Error("Error, invalid password accepted when compared to hash")
	}

	if w.Result().StatusCode != http.StatusUnauthorized {
		t.Error(fmt.Sprintf("Error, status code from failed validation is not `%v`", http.StatusForbidden))
	}
}

func TestValidateIdInRequest(t *testing.T) {
	w := *httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r = mux.SetURLVars(r, map[string]string{"testIdParam": "32"})

	if !ValidateIdInRequest(&w, r, "testIdParam") {
		t.Error("Error, valid request with id returned false by validator")
	}

	if ValidateIdInRequest(&w, r, "hello") {
		t.Error("Error, invalid request with no id returned true by validator")
	}

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Error(fmt.Sprintf("Error, status code from failed validation is not `%v`", http.StatusBadRequest))
	}
}

func TestValidateBody(t *testing.T) {
	w := *httptest.NewRecorder()
	var testModel_validRequest TestModel
	var testModel_invalidRequest TestModel

	requestBodyJson_valid, _ := json.Marshal(
		map[string]interface{}{
			"name": "Test",
			"age":  38,
		},
	)

	requestBodyJson_invalid, _ := json.Marshal(
		map[string]interface{}{
			"name": "Test",
			"age":  "55",
		},
	)

	if !ValidateBody(&w, bytes.NewReader(requestBodyJson_valid), &testModel_validRequest) {
		t.Error("Error, validator rejected valid request")
	}

	if ValidateBody(&w, bytes.NewReader(requestBodyJson_invalid), &testModel_invalidRequest) {
		t.Error("Error, validator approved an invalid request")
	}

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Error(fmt.Sprintf("Error, status code defined by validator is not `%v`", http.StatusBadRequest))
	}
}

func TestValidateHigherAccessLevel(t *testing.T) {
	w := *httptest.NewRecorder()

	if !ValidateHigherAccessLevel(&w, 4, 3) {
		t.Error("Error, validator failed with higher level requester")
	}

	if ValidateHigherAccessLevel(&w, 3, 3) {
		t.Error("Error, validator returned true with requester access level as same as requested")
	}

	if ValidateHigherAccessLevel(&w, 2, 3) {
		t.Error("Error, validator returned true with lower level requeste")
	}
}
