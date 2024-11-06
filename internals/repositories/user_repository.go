package repositories

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{Database: database}
}

var _ ports.UserRepository = (*UserRepository)(nil)

const (
	USERS_TABLE = "users"
)

func (r *UserRepository) Register(user *domain.RegisterRequest) (*domain.RegisterResponse, error) {

	var newUser *domain.RegisterResponse

	row := r.Database.Table(USERS_TABLE).Create(&user)
	r.Database.Table(USERS_TABLE).Where("id = ?", user.Id).Find(&newUser)

	if row.Error != nil {
		log.Error(row.Error)
		return nil, row.Error
	}

	return newUser, nil
}

func (r *UserRepository) Login(loginRequest *domain.LoginRequest) (*domain.User, error) {

	var user *domain.User
	row := r.Database.Where("email = ?", loginRequest.Email).Find(&user)

	if row.Error != nil {
		log.Error(row.Error)
		return nil, row.Error
	}

	return user, nil
}
