package test

import (
	"os"
	"path"
	"runtime"

	"github.com/mrotame/GoCrudApi/database"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func SetupTestDB(modelList ...interface{}) {
	// time.Sleep(4 * time.Second)
	database.SetupTestDB(modelList...)
	// time.Sleep(4 * time.Second)
}

func TeardownTestDB() {
	// time.Sleep(4 * time.Second)
	database.DropDatabase()
	// time.Sleep(4 * time.Second)
}
