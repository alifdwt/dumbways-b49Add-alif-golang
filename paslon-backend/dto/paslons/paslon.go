package paslonsdto

type CreatePaslonRequest struct {
	Name   string `json:"name" form:"name" validate:"required"`
	Vision string `json:"vision" form:"vision" validate:"required"`
	Image  string `json:"image" form:"image" validate:"required"`
}

type UpdatePaslonRequest struct {
	Name   string `json:"name" form:"name"`
	Vision string `json:"vision" form:"vision"`
	Image  string `json:"image" form:"image"`
}

type PaslonResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name" form:"name" validate:"required"`
	Vision string `json:"vision" form:"vision" validate:"required"`
	Image  string `json:"image" form:"image" validate:"required"`
}
