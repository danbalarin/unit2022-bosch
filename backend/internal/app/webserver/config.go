package webserver

import (
	"github.com/pkg/errors"
	"os"
	"strconv"
)

type webConfig struct {
	httpPort int
}

func NewWebConfig() (*webConfig, error) {
	httpPortRaw := os.Getenv("PORT")
	if httpPortRaw == "" {
		httpPortRaw = "3001"
	}
	httpPort, err := strconv.Atoi(httpPortRaw)
	if err != nil {
		return nil, errors.Wrap(err, "HTTP_PORT")
	}

	return &webConfig{
		httpPort,
	}, nil
}
