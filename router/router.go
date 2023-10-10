package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initiallize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
