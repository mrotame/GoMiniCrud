package test

import (
	"log"
	"os"
	"os/exec"

	"github.com/mrotame/GoCrudApi/database"
)

var Packages_to_test = []string{
	"database",
	"models",
	"utils",
	"validators",
	"views/userPage",
}

func RunTests() {
	database.IS_TESTING = true
	database.Open()

	for i := 0; i < len(Packages_to_test); i++ {
		cmd := exec.Command("go", "test", "-v", "./"+Packages_to_test[i])
		cmd.Env = os.Environ()

		cmd.Stdout = os.Stdout

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
	}
}
