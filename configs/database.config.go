package configs

import (
	"be_todo_app/database"
	"be_todo_app/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db_name = ""
var db_port = "3306"
var db_user = "root"
var db_password = ""
var db_host = "127.0.0.1"

func BootDatabase() {
	if dbNameEnv := os.Getenv("DB_NAME"); dbNameEnv != "" {
		db_name = dbNameEnv
	}

	if dbPortEnv := os.Getenv("DB_PORT"); dbPortEnv != "" {
		db_port = dbPortEnv
	}

	if dbUserEnv := os.Getenv("DB_USER"); dbUserEnv != "" {
		db_user = dbUserEnv
	}

	if dbPasswordEnv := os.Getenv("DB_PASSWORD"); dbPasswordEnv != "" {
		db_password = dbPasswordEnv
	}

	if dbHostEnv := os.Getenv("DB_HOST"); dbHostEnv != "" {
		db_host = dbHostEnv
	}
}

func ConnectDatabase() {
	var errConnection error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)
	database.DB, errConnection = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errConnection != nil {
		panic("Can't connect to database")
	} else {
		log.Println("Connected to database")
	}
}

func RunMigration() {
	err := database.DB.AutoMigrate(
		models.Todo{},
	)

	if err != nil {
		fmt.Println(err)
		log.Println("Failed to migrate schema")
	} else {
		log.Println("schemas migrated")
	}
}
