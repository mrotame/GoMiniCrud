package models

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/test"
	"github.com/mrotame/GoCrudApi/utils"
)

var USER_PASSWORD = "123"

var u = User{
	Name:     "Testing user",
	Age:      0,
	Email:    "Test@test.com",
	Password: USER_PASSWORD,
}

func TestUser(t *testing.T) {
	test.SetupTestDB(&User{})

	database.Save(&u)

	if u.ID == 0 {
		t.Error(fmt.Sprintf("Error, saved user named `%v` not saved properly in database", u.Name))
	}

	if !utils.CheckPasswordHash("123", u.Password) {
		t.Error(fmt.Sprintf("Error, Password match for saved user named `%v` did not matched. \n original password `%v` did not match user password `%v`", u.Name, USER_PASSWORD, u.Password))
	}

	test.TeardownTestDB()
}

func TestAsJson(t *testing.T) {
	test.SetupTestDB(&User{})

	database.Save(&u)

	jData, err := u.As_json()

	if err != nil {
		t.Error(fmt.Sprintf("Error converting user named `%v` to json.\n err: %v", u.Name, err))
	}

	var userMap map[string]interface{}

	err = json.Unmarshal(jData, &userMap)

	if err != nil {
		t.Error(fmt.Sprintf("Error converting jData from user named `%v` to map.\n err: %v", u.Name, err))
	}

	if userMap["password"] != "" {
		t.Error(fmt.Sprintf("Error. converted jData from user named `%v` contains the password `%v`", u.Name, userMap["password"]))
	}

	test.TeardownTestDB()
}
