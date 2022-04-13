package config

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"log"
	"os"
)

func LoadEnvVariables() error {
	envFiles := []string{}

	env := os.Getenv("ENV")
	if env == "" {
		env = "production"
	}

	envFiles = append(envFiles, "configs/.env."+env)
	envFiles = append(envFiles, "configs/.env")

	log.Println("Loading ENV variables from files:", envFiles)

	err := godotenv.Load(envFiles...)
	if err != nil {
		return errors.Wrap(err, "Loading environment variables")
	}

	return nil
}
