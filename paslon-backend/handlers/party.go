package handlers

import (
	partydto "myapp/dto/parties"
	dto "myapp/dto/result"
	"myapp/models"
	"myapp/repositories"
	"strconv"
	"time"

	// voterdto "myapp/dto/parties"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerP struct {
	PartyRepository repositories.PartyRepository
}

func HandlerParty(PartyRepository repositories.PartyRepository) *handlerP {
	return &handlerP{PartyRepository}
}

func (h *handlerP) FindParties(c echo.Context) error {
	parties, err := h.PartyRepository.FindParties()
	if err != nil {
		return c.JSON(http.StatusBadRequest,
		dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: parties})
}

func (h *handlerP) GetParty(c echo.Context) error {
	partyId, _ := strconv.Atoi(c.Param("partyId"))

	party, err := h.PartyRepository.GetParty(partyId)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: party})
}

func (h *handlerP) CreateParty(c echo.Context) error {
	request := new(partydto.CreatePartyRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	party := models.Party{
		// ID: request.ID,
		Name: request.Name,
		PaslonID: request.PaslonID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, err := h.PartyRepository.CreateParty(party)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}