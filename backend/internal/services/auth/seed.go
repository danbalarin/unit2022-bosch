package auth

import (
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"log"
)

type IAuthDbSeeder interface {
	Seed() error
}

type authDbSeeder struct {
	svc IAuthService
}

func NewAuthDbSeeder(svc IAuthService) IAuthDbSeeder {
	return &authDbSeeder{
		svc: svc,
	}
}

func (s *authDbSeeder) Seed() error {
	log.Println("Seeding auth database")
	err := s.svc.createUser(&entity.User{
		Email:    "admin@unit.cz",
		Password: s.svc.hashPassword("admin"),
		Role:     entity.Role_Admin,
	})
	if err != nil {
		if err != ErrUserAlreadyExists {
			return err
		} else {
			log.Println("User admin already exists")
		}
	}
	err = s.svc.createUser(&entity.User{
		Email:    "user@unit.cz",
		Password: s.svc.hashPassword("user"),
		Role:     entity.Role_User,
	})
	if err != nil {
		if err != ErrUserAlreadyExists {
			return err
		} else {
			log.Println("User basic already exists")
		}
	}
	log.Println("Seeding auth database complete!")

	return nil
}
