package voterdto

type CreateVoterRequest struct {
	ID        int64  `json:"id" form:"id"`
	VoterName string `json:"voter_name" form:"voter_name" validate:"required"`
	PaslonID  int    `json:"paslon_id" form:"paslon_id" validate:"required"`
	UserID    int    `json:"user_id" form:"user_id" validate:"required"`
}

type VoterResponse struct {
	ID        int64  `json:"id"`
	VoterName string `json:"voter_name" form:"voter_name" validate:"required"`
	PaslonID  int    `json:"paslon_id" form:"paslon_id" validate:"required"`
	UserID    uint   `json:"user_id" form:"user_id" validate:"required"`
}