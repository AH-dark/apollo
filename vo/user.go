package vo

import (
	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/hashids"
	"time"
)

type UserVO struct {
	ID        string    `json:"id" xml:"ID"`
	CreatedAt time.Time `json:"created_at" xml:"CreatedAt"`
	Username  string    `json:"username" xml:"Username"`
	Email     string    `json:"email" xml:"Email"`
	Role      int       `json:"role" xml:"Role"`
}

func BuildUserVO(user *model.User) UserVO {
	return UserVO{
		ID:        hashids.Encode(user.ID, hashids.UserHash),
		CreatedAt: user.CreatedAt,
		Username:  user.Username,
		Email:     user.Email,
		Role:      int(user.Role),
	}
}
