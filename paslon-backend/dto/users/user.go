package userdto

type CreateUserRequest struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name" form:"full_name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name" form:"full_name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}