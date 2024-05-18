package database

import (
	config "slanj/src/common/config"
	"slanj/src/common/database/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Init() {
	dsn := config.GetConfig().Params.DataBase.Url

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	connection = conn
}

func RunMigrates() {
	connection.AutoMigrate(&schema.User{}, &schema.Score{})
}

func GetConnection() *gorm.DB {
	return connection
}
