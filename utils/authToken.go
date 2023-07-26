package utils

import (
	"encoding/hex"
	"math/rand"
	"net/http"
	"strings"
)

var RANDOM_TOKEN_LENGTH int = 512

func GenerateRandomToken() (string, error) {
	b := make([]byte, RANDOM_TOKEN_LENGTH/2)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func GetTokenFromRequest(r *http.Request) string {
	token := r.Header.Get("Authorization")
	token_slice := strings.Split(token, "Bearer")

	token = token_slice[len(token_slice)-1]
	token = strings.TrimSpace(token)

	return token
}
