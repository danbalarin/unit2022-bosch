package orm

import (
	"github.com/pkg/errors"
	"os"
)

type dbConfig struct {
	url string
}

func NewDbConfig() (*dbConfig, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}

	return &dbConfig{
		dbUrl,
	}, nil
}
