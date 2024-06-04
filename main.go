package main

import (
	"GoAlbums/api/v1/routes"
	"GoAlbums/config"
	"GoAlbums/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load database credentials
	config.LoadCreds()

	// Connect to the database
	if config.UseDB {
		passed := db.ConnectDB(config.DBCreds.Username, config.DBCreds.Password, config.DBCreds.HostPath, config.DBCreds.Database, config.DBCreds.IsPrivatePath)
		config.UseDB = passed
	}

	// Create a new Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupAlbumRoutes(router)
	// You can add more route setups here for other resources

	// Start the server
	router.Run("localhost:8080")
}
