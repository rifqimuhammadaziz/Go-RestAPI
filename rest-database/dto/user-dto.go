package dto

// UserUpdateDTO is used by client when request PUT method profile
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"name" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"min=6"`
}

// UserCreateDTO is used by client when register a user
//type UserCreateDTO struct {
//	Name     string `json:"name" form:"name" binding:"required"`
//	Email    string `json:"email" form:"name" binding:"required" validate:"email"`
//	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
//}
