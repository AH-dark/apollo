package model

import (
	"github.com/AH-dark/apollo/pkg/crypto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string     `gorm:"uniqueIndex;not null"`
	Password string     `gorm:"not null"`
	Email    string     `gorm:"uniqueIndex;not null"`
	Role     UserStatus `gorm:"not null;default:0"`

	Comments []*Comment `gorm:"foreignKey:ReplyUser"`
}

type UserStatus int

const (
	UserStatusWatcher UserStatus = iota
	UserStatusAdmin
)

func (u *User) ComparePassword(password string) bool {
	return u.Password == crypto.Password(password)
}
