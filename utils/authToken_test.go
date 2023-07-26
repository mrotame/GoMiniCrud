package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthToken(t *testing.T) {
	var random_token string
	var err error

	random_token, err = GenerateRandomToken()

	if err != nil {
		t.Error("Error while hashing password", err)
	}

	if len(random_token) < RANDOM_TOKEN_LENGTH {
		t.Error(fmt.Sprintf("Error, random token length too short. `%s` has less than 60 characters", random_token))
	}
}

func TestGetTokenFromRequest(t *testing.T) {
	token, _ := GenerateRandomToken()

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Add("Authorization", token)

	extracted_token := GetTokenFromRequest(r)

	if extracted_token != token {
		t.Error(fmt.Sprintf("Error, extracted token from request `%s` is different from original token `%s`", extracted_token, token))
	}
}
