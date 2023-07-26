package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/mrotame/GoCrudApi/database"
	"github.com/mrotame/GoCrudApi/models"
	"github.com/mrotame/GoCrudApi/test"
	"github.com/mrotame/GoCrudApi/views"
	"github.com/mrotame/GoCrudApi/views/userPage"
)

var ip = "0.0.0.0"
var port = "8001"

func main() {
	load_dot_env()

	run_tests_if_case()

	r := registerViews()
	configure_database()

	log.Printf("Server started and listening in port %v \n", port)
	http.ListenAndServe(fmt.Sprintf("%v:%v", ip, port), r)
}

func load_dot_env() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading the `.env` file:\n%v", err)
	}
}

func configure_database() {
	database.Open()
	database.DB.AutoMigrate(models.ModelList...)
}

func registerViews() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", views.HomePage)
	r.HandleFunc("/auth", views.AuthPage)
	r.HandleFunc("/user", userPage.UserPage)
	r.HandleFunc("/user/{id}", userPage.UserPage)
	http.Handle("/", r)
	return r
}

func run_tests_if_case() {
	if os.Getenv("TEST_ONLY") == "true" {
		test.RunTests()
		os.Exit(0)
	}
	if len(os.Args) > 1 {
		args := os.Args[1:]
		if args[0] == "-t" || args[0] == "--test" {
			test.RunTests()
			os.Exit(0)
		}
	}
}
