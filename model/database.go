package model

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var envVars environmentVariables

func init() {
	godotenv.Load(".env")
	envVars.Load()
	connectDatabase()
}

func connectDatabase() {
	databaseStringConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", envVars.User, envVars.Password, envVars.Host, envVars.Port, envVars.Database)
	db, err := gorm.Open(mysql.Open(databaseStringConfig), &gorm.Config{})

	migrate(db)

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Bookshelf{})
	db.AutoMigrate(&Category{})
}
