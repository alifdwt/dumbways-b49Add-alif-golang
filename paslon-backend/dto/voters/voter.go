package voterdto

type CreateVoterRequest struct {
	VoterName string `json:"voter_name" form:"voter_name" validate:"required"`
	PaslonID  string `json:"paslon_id" form:"paslon_id" validate:"required"`
}

type UpdateVoterRequest struct {
	VoterName string `json:"voter_name" form:"voter_name"`
	PaslonID  string `json:"paslon_id" form:"paslon_id"`
}

type VoterResponse struct {
	ID        int    `json:"id"`
	VoterName string `json:"voter_name" form:"voter_name" validate:"required"`
	PaslonID  int    `json:"paslon_id" form:"paslon_id" validate:"required"`
}