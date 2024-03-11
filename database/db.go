package database

import (
	"github/batuhanzorbeyzengin/insider-messaging-system/configs"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	config, err := configs.LoadConfig()
	if err != nil {
		return err
	}

	connectionString := configs.GetMySQLConnectionString(config.Database)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Connected to MySQL database")
	DB = db

	return nil
}
