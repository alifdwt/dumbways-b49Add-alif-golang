package voterdto

type CreateVoterRequest struct {
	ID        int64  `json:"id" form:"id" validate:"required"`
	VoterName string `json:"voter_name" form:"voter_name" validate:"required"`
	PaslonID  int    `json:"paslon_id" form:"paslon_id" validate:"required"`
}

// type UpdateVoterRequest struct {
// 	VoterName string `json:"voter_name" form:"voter_name"`
// 	PaslonID  string `json:"paslon_id" form:"paslon_id"`
// }

type VoterResponse struct {
	ID        int64  `json:"id"`
	VoterName string `json:"voter_name" form:"voter_name" validate:"required"`
	PaslonID  int    `json:"paslon_id" form:"paslon_id" validate:"required"`
}