package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	Namespace      string
	AllowOrigins   string
	DeployURL      string
	ServerPort     string
	RepositoryPath string
}

func GetEnvironment() (*Environment, error) {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "local"
	}

	envFile := fmt.Sprintf(".env.%s", env)
	if err := godotenv.Load(envFile); err != nil {
		return &Environment{}, fmt.Errorf("error loading %s file: %w", envFile, err)
	}

	return &Environment{
		Namespace:      os.Getenv("NAMESPACE"),
		AllowOrigins:   os.Getenv("ALLOWED_ORIGIN"),
		DeployURL:      os.Getenv("DEPLOY_URL"),
		ServerPort:     os.Getenv("SERVER_PORT"),
		RepositoryPath: os.Getenv("REPOSITORY_PATH"),
	}, nil
}
