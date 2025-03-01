package service

import (
	"errors"
	"go-auth/domain/dto"
	"go-auth/domain/model"
	"go-auth/repository"
	"time"
)

type IUserService interface {
	GetAllUsers(role string) ([]model.User, error)
	GetUserByID(userID string, role string) (model.User, error)
	SetRole(req dto.SetRoleRequest) (model.User, error)
}

type userService struct {
	repo repository.IUserRepo
}

func NewUserService(repo repository.IUserRepo) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAllUsers(role string) ([]model.User, error) {
	if role == "super_admin" {
		return s.repo.GetAllUsers()
	}

	if role == "admin" {
		return s.repo.GetAllUsersByRole([]string{"admin", "user"})
	}

	return nil, errors.New("unauthorized to view users")
}

func (s *userService) GetUserByID(userID string, role string) (model.User, error) {
    user, err := s.repo.GetUserByID(userID)
    if err != nil {
        return model.User{}, err
    }

    // Superadmin can access all users, admins can access only admins and users
    if role != "superadmin" && !(role == "admin" && (user.Role == "admin" || user.Role == "user")) {
        return model.User{}, errors.New("unauthorized to view this user")
    }

    return user, nil
}

func (s *userService) SetRole(req dto.SetRoleRequest) (model.User, error) {
	user, err := s.repo.GetUserByID(req.UserID)
    if err != nil {
        return model.User{}, err
    }

    user.Role = req.Role
	user.UpdatedAt = int(time.Now().Unix())

	userUpdate, err := s.repo.UpdateUser(user)
	if err != nil {
		return model.User{}, err
	}
    return userUpdate, nil
}