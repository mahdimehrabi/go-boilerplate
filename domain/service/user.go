package service

import (
	models "bbdk/domain/entity"
	userRepo "bbdk/domain/repository/user"
	logger "bbdk/infrastructure/log"
	"errors"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}
type userService struct {
	userRepo userRepo.Repository
	logger   logger.Logger
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo userRepo.Repository, logger logger.Logger) UserService {
	return &userService{userRepo: userRepo, logger: logger}
}

func (s *userService) CreateUser(user *models.User) error {
	if err := s.userRepo.CreateUser(user); err != nil {
		s.logger.Errorf("failed to create user:%s", err.Error())
		return err
	}
	return nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			return nil, err
		}
		s.logger.Errorf("failed to get user:%s", err.Error())
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(user *models.User) error {
	if err := s.userRepo.UpdateUser(user); err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			return err
		}
		s.logger.Errorf("failed to update user:%s", err.Error())
		return err
	}
	return nil
}

func (s *userService) DeleteUser(id uint) error {
	if err := s.userRepo.DeleteUser(id); err != nil {
		if errors.Is(err, userRepo.ErrNotFound) {
			return err
		}
		s.logger.Errorf("failed to delete user:%s", err.Error())
	}
	return nil
}