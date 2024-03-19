package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Dev4w4n/cdn.e-masjid.my/api/image-store-api/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server ...")

	env, err := utils.GetEnvironment()
	if err != nil {
		log.Fatalf("Error getting environment: %v", err)
	}

	// Initialize image folders
	log.Println("Initializing image folders ...")
	initializeImageFolders(env)

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

func initializeImageFolders(env *utils.Environment) error {
	namespace := env.Namespace

	repositoryPath := env.RepositoryPath

	// Change directory to the repository path
	if err := os.Chdir(repositoryPath); err != nil {
		log.Printf("Error changing directory: %v\n", err)
		return err
	}

	// Split the namespace by comma to get individual folder names
	folders := strings.Split(namespace, ",")

	// Iterate through each folder name and check if it exists
	for _, folder := range folders {
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			// If folder doesn't exist, create it
			err := os.Mkdir(folder, 0755)
			if err != nil {
				log.Printf("Error creating folder %s: %v\n", folder, err)
				return err
			} else {
				log.Printf("Folder %s created successfully.\n", folder)
			}
		} else if err != nil {
			log.Printf("Error checking folder %s: %v\n", folder, err)
			return err
		} else {
			log.Printf("Folder %s already exists.\n", folder)
		}
	}
	return nil
}
