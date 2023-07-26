package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/mrotame/GoCrudApi/utils"
)

type Auth struct {
	gorm.Model
	User_id         uint      `json:"user_id"`
	Token           string    `json:"token"`
	Expiration_date time.Time `json:"expiration_date"`
}

func (a *Auth) BeforeCreate(db *gorm.DB) (err error) {
	a.Expiration_date = time.Now().AddDate(0, 0, 1)
	a.Token, _ = utils.GenerateRandomToken()
	return
}

func (a *Auth) IsValid() bool {
	if a.Expiration_date.Before(time.Now()) {
		return false
	}

	return true
}
