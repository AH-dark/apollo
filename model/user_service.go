package model

import (
	"github.com/AH-dark/apollo/dto"
	"github.com/AH-dark/apollo/pkg/crypto"
	"gorm.io/gorm"
)

type UserService interface {
	// GetUserByID get user by id
	GetUserByID(id uint) (*User, error)
	// GetUserByUsername get user by username
	GetUserByUsername(username string) (*User, error)
	// GetUserByEmail get user by email
	GetUserByEmail(email string) (*User, error)
	// ComparePassword compare password
	ComparePassword(login, password string) bool
	// CreateUser create user
	CreateUser(user *User) error
	// CreateUserByDTO create user by dto
	CreateUserByDTO(dto dto.AdminCreateUserDTO) error
	// UpdateUser update user
	UpdateUser(user *User) error
	// DeleteUser delete user
	DeleteUser(id uint) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

func (s *userService) GetUserByID(id uint) (*User, error) {
	var user User
	err := s.db.Model(&User{}).Where("id = ?", id).First(&user).Error
	return &user, err
}

func (s *userService) GetUserByUsername(username string) (*User, error) {
	var user User
	err := s.db.Model(&User{}).Where("username = ?", username).First(&user).Error
	return &user, err
}

func (s *userService) GetUserByEmail(email string) (*User, error) {
	var user User
	err := s.db.Model(&User{}).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (s *userService) ComparePassword(login, password string) bool {
	var user User
	err := s.db.Model(&User{}).Where("username = ? OR email = ?", login, login).First(&user).Error
	if err != nil {
		return false
	}

	return user.ComparePassword(password)
}

func (s *userService) CreateUser(user *User) error {
	user.Password = crypto.Password(user.Password)

	return s.db.Model(&User{}).Create(user).Error
}

func (s *userService) CreateUserByDTO(dto dto.AdminCreateUserDTO) error {
	user := &User{
		Username: dto.Username,
		Password: dto.Password,
		Email:    dto.Email,
		Role:     UserStatus(dto.Role),
	}

	return s.CreateUser(user)
}

func (s *userService) UpdateUser(user *User) error {
	return s.db.Model(&User{}).Save(user).Error
}

func (s *userService) DeleteUser(id uint) error {
	return s.db.Model(&User{}).Where("id = ?", id).Delete(&User{}).Error
}
