package partydto

type CreatePartyRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	PaslonID int    `json:"paslon_id" form:"paslon_id" validate:"required"`
}

type UpdatePartyRequest struct {
	Name     string `json:"name" form:"name"`
	PaslonID int    `json:"paslon_id" form:"paslon_id"`
}

type PartyResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	PaslonID int    `json:"paslon_id" form:"paslon_id" validate:"required"`
}