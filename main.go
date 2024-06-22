package main

import (
	"GoAlbums/config"
	"GoAlbums/internal/db"
	apiRoutes "GoAlbums/routes/api/v1/routers"
	webRoutes "GoAlbums/routes/web"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load Configs
	config.InitializeConfig()

	// Connect to the database
	if config.UseDB {
		passed := db.ConnectDB(config.DBCreds.Username, config.DBCreds.Password, config.DBCreds.HostPath, config.DBCreds.Database, config.DBCreds.IsPrivatePath)
		config.UseDB = passed
	}

	// Create a new Gin router
	router := gin.Default()

	// Setup routes
	apiRoutes.SetupAlbumRoutes(router)
	apiRoutes.SetupAuthRoutes(router)
	apiRoutes.SetupAccountRoutes(router)
	apiRoutes.SetupUserLogsRoutes(router)
	webRoutes.SetupPublicRoutes(router)
	// You can add more route setups here for other resources

	// Start the server
	router.Run("localhost:8080")
}
