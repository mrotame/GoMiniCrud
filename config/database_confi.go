package config

import (
	"os"
)

type database_settings struct {
	Dbname   string
	Host     string
	User     string
	Password string
	Port     string
	Sslmode  string
}

func Get_prod_db() *database_settings {
	return &database_settings{
		Dbname:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		Sslmode:  os.Getenv("DB_SSLMODE"),
	}
}

func Get_test_db() *database_settings {
	return &database_settings{
		Dbname:   os.Getenv("DB_TESTNAME"),
		Host:     os.Getenv("DB_TESTHOST"),
		User:     os.Getenv("DB_TESTUSERNAME"),
		Password: os.Getenv("DB_TESTPASSWORD"),
		Port:     os.Getenv("DB_TESTPORT"),
		Sslmode:  os.Getenv("DB_TESTSSLMODE"),
	}
}
