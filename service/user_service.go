package service

import (
	"go-auth/domain/dto"
	"go-auth/domain/model"
	"go-auth/repository"
	"time"
)

type IUserService interface {
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