package dto

// BookUpdateDTO is used by client to update the book
type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

// BookCreateDTO is used by client to create new the book
type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
