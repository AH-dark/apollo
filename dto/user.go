package dto

type AdminCreateUserDTO struct {
	Username string `json:"username" xml:"Username" form:"username" binding:"required,max=32"`
	Password string `json:"password" xml:"Password" form:"password" binding:"required,max=32"`
	Email    string `json:"email" xml:"Email" form:"email" binding:"required,email,max=64"`
	Role     int    `json:"role" xml:"Role" form:"role" binding:"required,oneof=0 1"`
}
