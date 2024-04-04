package main

import (
	"log"
	"net/http"

	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/config"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/controller"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/repository"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/service"
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

	db, err := config.DatabaseConnection(env)
	if err != nil {
		log.Fatalf("Error getting database connection: %v", err)
	}

	dbRepository := repository.NewDbRepository(db)
	fileRepository := repository.NewFileRepository()

	cdnService := service.NewCDNService(dbRepository, fileRepository)

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{env.AllowOrigins}
	config.AllowMethods = []string{"POST"}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.New(config))

	// isLocalEnv := os.Getenv("GO_ENV")
	// if isLocalEnv != "" && isLocalEnv != "dev" {
	// 	r.Use(controllerMiddleware(env))
	// }

	_ = controller.NewCDNController(r, cdnService, env)

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
		secFetchSite := c.Request.Header.Get("Sec-Fetch-Site")

		log.Println("Origin: ", origin)
		log.Println("Sec-Fetch-Site: ", secFetchSite)

		if origin != allowedOrigin && secFetchSite != "same-origin" && secFetchSite != "same-site" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
	}
}
