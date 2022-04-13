package auth

import (
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
)

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrInvalidCredentials = errors.New("invalid credentials")

type LoginResponse struct {
	Token string `json:"token"`
	Role int `json:"role"`
}

type IAuthRepository interface {
	insertUser(user *entity.User) error
	findUser(id uint) (*entity.User, error)
	findUserByEmail(email string) (*entity.User, error)
}

type IAuthService interface {
	Login(email, password string) (*LoginResponse, error)
	GetUserRoleFromToken(token string) (*entity.UserRole, error)
	GetUserByID(id uint) (*entity.User, error)

	createUser(user *entity.User) error
	hashPassword(password string) string
}

type authService struct {
	repo      IAuthRepository
	jwtSecret string
}

func NewAuthService(repo IAuthRepository, config *authConfig) IAuthService {
	return &authService{
		repo:      repo,
		jwtSecret: config.jwtSecret,
	}
}

func (s *authService) Login(email, password string) (*LoginResponse, error) {
	user, err := s.repo.findUserByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find user")
	}

	if !s.compareHashAndPassword(user.Password, password) {
		return nil, ErrInvalidCredentials
	}

	token, err := s.createJwtToken(user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create jwt token")
	}

	res := &LoginResponse{
		Token: token,
		Role: int(user.Role),
	}

	return res, nil
}

func (s *authService) GetUserByID(id uint) (*entity.User, error) {
	return s.repo.findUser(id)
}

func (s *authService) GetUserRoleFromToken(token string) (*entity.UserRole, error) {
	userID, err := s.parseJwtToken(token)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user, err := s.GetUserByID(userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// User doesn't exist
	if user == nil {
		return nil, nil
	}

	// User exists
	return &user.Role, nil
}

func (s *authService) createUser(user *entity.User) error {
	return s.repo.insertUser(user)
}
