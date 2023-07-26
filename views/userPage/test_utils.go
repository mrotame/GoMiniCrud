package userPage

import (
	"fmt"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
)

func GetTestUser() (models.User, models.Auth) {
	var err error
	var user models.User

	for i := 1; i < 10; i++ {

		email := "test" + fmt.Sprint(i)

		user = models.User{
			Name:     "Test user " + fmt.Sprint(i),
			Email:    email + "@test.com",
			Password: "123",
		}
		err = database.Save(&user).Error
		if err == nil {
			break
		}
	}

	var auth = models.Auth{
		User_id: user.ID,
	}
	database.Save(&auth)
	return user, auth
}
