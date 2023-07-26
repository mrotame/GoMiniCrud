package utils

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "test"

	hashed, err := HashPassword(password)

	if err != nil {
		t.Error("Error while hashing password", err)
	}

	if hashed == password {
		t.Error(fmt.Sprintf("Error, password `%s` is equal to hashed password `%s`", password, hashed))
	}
}

func TestComparePassword(t *testing.T) {
	password := "test"
	hashed, err := HashPassword(password)

	if err != nil {
		t.Error("Error while hashing password", err)
	}

	if !CheckPasswordHash(password, hashed) {
		t.Error(fmt.Sprintf("Error comparing, password `%s` with hashed password `%s`", password, hashed))
	}
}
