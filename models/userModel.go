package models

import (
	"encoding/json"

	"gorm.io/gorm"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/utils"
)

var db *gorm.DB = database.DB

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Age         int    `json:"age"`
	AccessLevel int    `gorm:"not null; default:0" json:"accessLevel"`
	Email       string `gorm:"unique;not null; default:null" json:"email"`
	Password    string `gorm:"not null; default:null; size:256" json:"password"`
}

type UserUpdate struct {
	Name  *string `json:"name"`
	Age   *int    `json:"age"`
	Email *string `json:"email"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	u.Password, err = utils.HashPassword(u.Password)

	return
}

func (u User) As_json() ([]byte, error) {
	type user User // prevent recursion
	x := user(u)
	x.Password = ""
	return json.Marshal(x)
}

func (u User) As_map() (map[string]interface{}, error) {
	type user User // prevent recursion
	var userAsMap map[string]interface{}
	var err error

	x := user(u)
	x.Password = ""

	jData, err := json.Marshal(x)
	if err == nil {
		err = json.Unmarshal(jData, &userAsMap)
	}
	return userAsMap, err
}

func (u User) GetBy_id(id uint) User {
	var user User
	database.GetOne(&user, id)
	return user
}

func GetUserBy_Credentials(u User, email string, password string) *gorm.DB {
	var user User
	result := database.GetOne(&user, "email = ?", email)
	if utils.CheckPasswordHash(password, user.Password) {
		u = user
	} else {
		u = User{}
	}

	return result
}

func GetUserBy_AuthToken(u *User, token string) *gorm.DB {
	var auth Auth

	result := database.GetOne(&auth, "Token = ?", token)

	if result.Error != nil {
		return result
	}

	result = database.GetOne(u, auth.User_id)
	return result
}
