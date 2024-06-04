package config

import (
	"embed"
	"encoding/json"
	"log"
	"os"

	"GoAlbums/internal/db"
)

var (
	//go:embed db.json
	res     embed.FS
	UseDB   = true
	DBCreds = db.DBCredentials{}
)

func LoadCreds() {
	data, _ := res.ReadFile("db.json")
	err := json.Unmarshal(data, &DBCreds)
	if err != nil {
		log.Printf("Credentials Error: %v", err)
		os.Exit(1)
	}
}
