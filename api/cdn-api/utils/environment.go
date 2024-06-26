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
	Workdir        string
	RepositoryPath string
	DbHost         string
	DbPort         string
	DbUser         string
	DbPassword     string
	DbName         string
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
		DbHost:         os.Getenv("DB_HOST"),
		DbPort:         os.Getenv("DB_PORT"),
		DbUser:         os.Getenv("DB_USER"),
		DbPassword:     os.Getenv("DB_PASSWORD"),
		DbName:         os.Getenv("DB_NAME"),
		Namespace:      os.Getenv("NAMESPACE"),
		AllowOrigins:   os.Getenv("ALLOWED_ORIGIN"),
		DeployURL:      os.Getenv("DEPLOY_URL"),
		ServerPort:     os.Getenv("SERVER_PORT"),
		Workdir:        os.Getenv("WORKDIR"),
		RepositoryPath: os.Getenv("REPOSITORY_PATH"),
	}, nil
}
