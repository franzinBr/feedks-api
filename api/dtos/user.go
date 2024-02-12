package dtos

type CreateUserRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	UserName  string `json:"userName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}
