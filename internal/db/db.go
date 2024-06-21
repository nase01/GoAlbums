package db

import (
	"log"

	"GoAlbums/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		log.Fatalf("Failed to connect database at %v: %v", unixPath, err)
		return false
	} else {
		log.Printf("Connected to database: %v", unixPath)
	}

	log.Print("ORM Migrating DB Objects Begin")

	if err := DB.DB.AutoMigrate(&models.Album{}, &models.User{}, &models.UserLogs{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Print("ORM Migrating DB Objects End")

	return true
}
