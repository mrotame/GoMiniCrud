package database

import (
	"fmt"
	"testing"
)

func TestCreateInstanceAndSaveInDB(t *testing.T) {
	SetupTestDB(&TestInstance{})

	var instanceFromDb TestInstance
	u := TestInstance{
		Name:         "Testing instance",
		Unique_field: "Test",
	}

	Save(&u)

	err := GetOne(&instanceFromDb, u.ID).Error

	if err != nil {
		t.Error(fmt.Sprintf("Error, saving instance named `%v`.\n err: %v", u.Name, err))
	}

	if u.ID == 0 {
		t.Error(fmt.Sprintf("Error, saved instance named `%v` not found in database after save", u.Name))
	}

	DropDatabase()
}

func TestDropDatabaseAndSetupAgain(t *testing.T) {
	SetupTestDB(&TestInstance{})

	var instanceFromDb TestInstance
	u := TestInstance{
		Name:         "Testing instance",
		Unique_field: "Test",
	}
	Save(&u)

	DropDatabase()
	SetupTestDB(&TestInstance{})

	err := GetOne(&instanceFromDb, u.ID).Error

	if err == nil {
		t.Error("Error, new setuped database returned data from before the drop")
	}

	DropDatabase()
}
