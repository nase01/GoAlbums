package config

import (
	"embed"
	"encoding/json"
	"log"
	"os"

	"GoAlbums/internal/db"

	"github.com/spf13/viper"
)

var (
	//go:embed db.json
	res     embed.FS
	UseDB   = true
	DBCreds = db.DBCredentials{}
	jwtKey  []byte
)

func LoadCreds() {
	data, _ := res.ReadFile("db.json")
	err := json.Unmarshal(data, &DBCreds)
	if err != nil {
		log.Printf("Credentials Error: %v", err)
		os.Exit(1)
	}
}

func LoadENV() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading .env file: %v", err)
		os.Exit(1)
	}
}

func InitializeConfig() {
	LoadENV()
	LoadCreds()
	jwtKey = []byte(viper.GetString("JWT_SECRET_KEY"))
}

func GetJWTKey() []byte {
	return jwtKey
}
