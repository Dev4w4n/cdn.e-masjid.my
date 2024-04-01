package main

import (
	"log"
	"net/http"

	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server ...")

	env, err := utils.GetEnvironment()
	if err != nil {
		log.Fatalf("Error getting environment: %v", err)
	}

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{env.AllowOrigins}
	config.AllowMethods = []string{"POST"}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.New(config))
	r.Use(controllerMiddleware(env))

	go func() {
		err := r.Run(":" + env.ServerPort)
		if err != nil {
			log.Fatal("Error starting the server:", err)
		}
	}()

	log.Println("Server listening on port ", env.ServerPort)

	select {} // Block indefinitely to keep the program running
}

// Strictly allow from allowedOrigin
func controllerMiddleware(env *utils.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the request origin is allowed
		allowedOrigin := env.AllowOrigins
		origin := c.GetHeader("Origin")

		log.Println("Origin: ", origin)
		if origin != allowedOrigin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
	}
}
