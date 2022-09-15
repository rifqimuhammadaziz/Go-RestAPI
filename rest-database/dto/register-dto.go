package dto

// RegisterDTO is used when client request POST method from /register URL
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required,min=5"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}
