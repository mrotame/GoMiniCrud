package database

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/mrotame/GoCrudApi/config"
)

var IS_TESTING = strings.HasSuffix(os.Args[0], ".test")
var DB *gorm.DB

type IModel interface {
}

func get_dsn(test_env bool) string {
	var dbsettings = config.Get_prod_db()

	if test_env {
		dbsettings = config.Get_test_db()
	}

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
		dbsettings.Host, dbsettings.User, dbsettings.Password, dbsettings.Dbname, dbsettings.Port, dbsettings.Sslmode,
	)
	log.Println(dsn)
	return dsn
}

func Open() {
	var err error
	DB, err = gorm.Open(postgres.Open(get_dsn(IS_TESTING)), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		panic("Failed to connect database")
	}
}

func SetupTestDB(modelList ...interface{}) {
	if DB == nil {
		Open()
	}
	MakeMigrations(modelList)
}

func DropDatabase() {
	tables, _ := DB.Migrator().GetTables()

	s := make([]interface{}, len(tables))

	for i, v := range tables {
		s[i] = v
	}

	DB.Migrator().DropTable(s...)
}

func Migrate(modelList ...interface{}) {
	err := MakeMigrations(modelList)

	for _, err := range err {
		if err != nil {
			log.Fatalln("Error migrating the database", err)
		}
	}

}

func DropTables() {
	tables, _ := DB.Migrator().GetTables()

	s := make([]interface{}, len(tables))

	for i, v := range tables {
		s[i] = v
	}

	DB.Migrator().DropTable(s...)
}
