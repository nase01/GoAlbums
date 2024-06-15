package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"GoAlbums/internal/db"

	"github.com/spf13/viper"
)

var (
	UseDB   = true
	DBCreds = db.DBCredentials{}
	jwtKey  []byte
)

func LoadCreds() {
	IsPrivatePath, err := strconv.ParseBool(viper.GetString("DB_IS_PRIVATE_PATH"))
	if err != nil {
		log.Printf("Error parsing DB_IS_PRIVATE: %v", err)
		os.Exit(1)
	}

	dbConfig := map[string]interface{}{
		"HostPath":      viper.GetString("DB_HOST_PATH"),
		"Username":      viper.GetString("DB_USERNAME"),
		"Password":      viper.GetString("DB_PASSWORD"),
		"Database":      viper.GetString("DB_NAME"),
		"IsPrivatePath": IsPrivatePath,
	}

	data, err := json.Marshal(dbConfig)
	if err != nil {
		log.Printf("Error marshaling dbConfig to JSON: %v", err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &DBCreds)
	if err != nil {
		log.Printf("Error unmarshaling JSON to DBCreds: %v", err)
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
