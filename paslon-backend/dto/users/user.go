package userdto

type CreateUserRequest struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name" form:"full_name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	// VoterID  int64  `json:"voter_id" form:"voter_id" validate:"required"`
}

type UpdateUserRequest struct {
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	// VoterID  int64  `json:"voter_id" form:"voter_id"`
}

type UserResponse struct {
	ID int `json:"id"`
	// VoterID  int64  `json:"voter_id" form:"voter_id" validate:"required"`
	FullName string `json:"full_name" form:"full_name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	// Password string `json:"password" form:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}