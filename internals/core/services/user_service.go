package services

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
	"github.com/n0o01lh/llp/internals/utils"
)

type UserService struct {
	UserRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

const (
	ROLE_ADMIN   = 1
	ROLE_USER    = 2
	ROLE_TEACHER = 3
)

var _ ports.UserService = (*UserService)(nil)

func (s *UserService) Register(user *domain.RegisterRequest) (*domain.RegisterResponse, error) {

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	user.Password = hashedPassword
	user.RoleId = ROLE_USER
	userRegistered, err := s.UserRepository.Register(user)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return userRegistered, nil
}

func (s *UserService) Login(loginRequest *domain.LoginRequest) (*domain.LoginResponse, error) {

	user, err := s.UserRepository.Login(loginRequest)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = utils.ComparePassword([]byte(user.Password), []byte(loginRequest.Password))

	if err != nil {
		err = fmt.Errorf("Password incorrect for email %s", loginRequest.Email)
		log.Error(err)
		return nil, err
	}

	//Generate a session or token for authentication

	return &domain.LoginResponse{
		Username: user.Username,
		Token:    "token",
	}, nil
}
