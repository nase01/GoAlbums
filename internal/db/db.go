package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"GoAlbums/internal/models"
)

type DBLogger struct {
	DB *gorm.DB
}

type DBCredentials struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	HostPath      string `json:"hostpath"`
	Database      string `json:"database"`
	IsPrivatePath bool   `json:"isprivatepath"`
}

var DB DBLogger

func ConnectDB(userName string, password string, unixPath string, database string, isPrivate bool) bool {
	var err error

	tp := "tcp"
	if isPrivate {
		tp = "unix"
	}

	dsn := userName + ":" + password + "@" + tp + "(" + unixPath + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=UTC"
	DB.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect database: %v", unixPath)
		log.Printf("Failed to connect database: %v", err)
		return false
	} else {
		log.Printf("Connected to database: %v", unixPath)
	}

	log.Print("ORM Migrating DB Objects Begin")

	DB.DB.AutoMigrate(&models.Album{})

	log.Print("ORM Migrating DB Objects End")

	return true
}
