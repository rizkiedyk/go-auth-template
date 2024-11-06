package service

import (
	"errors"
	"fmt"
	"go-auth/domain/model"
	"go-auth/repository"
	"go-auth/utils/jwt"
	"go-auth/utils/security"

	"github.com/go-playground/validator/v10"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("main")

type IAuthService interface {
	Register(user model.User) error
    Login(username, password string) (string, error)
}

type authService struct {
	repo repository.IAuthRepo
}

func NewAuthService(repo repository.IAuthRepo) *authService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(user model.User) error {
	fmt.Printf("registering user: %v\n", user)
	exist, err := s.repo.GetUserExisting(user.Username)
	if err != nil {
		fmt.Println("testing")
		return err
	}

	if exist.Username != "" {
		return errors.New("user already exist")
	}

	validation := validator.New()
	err = validation.Struct(user)
	if err != nil {
		return err
	}

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	// logging.INFO("registering user service: %v\n", user)
	err = s.repo.RegisterRepo(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.repo.CheckUserExistingForLogin(username)
	if err != nil {
		return "", err
	}

	if user.Username == "" || security.CheckPassword(user.Password, password) != nil {
		return "", errors.New("invalid credentials")
	}

	return jwt.GenerateToken(user)
}