package auth

import (
	"log"
	"math/rand"
	"os"
)

type authConfig struct {
	jwtSecret string
}

func NewAuthConfig() *authConfig {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = randStringBytesRmndr(32)
		log.Println("Generated random JWT_SECRET")
	}

	return &authConfig{
		jwtSecret,
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
