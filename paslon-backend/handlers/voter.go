package handlers

import (
	"fmt"
	dto "myapp/dto/result"

	voterdto "myapp/dto/voters"
	"myapp/models"
	"myapp/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerV struct {
	VoterRepository repositories.VoterRepository
}

func HandlerVoter(VoterRepository repositories.VoterRepository) *handlerV {
	return &handlerV{VoterRepository}
}

func (h *handlerV) FindVoter(c echo.Context) error {
	voters, err := h.VoterRepository.FindVoters()
	if err != nil {
		return c.JSON(http.StatusBadRequest,
		dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: voters})
}

func (h *handlerV) GetVoter(c echo.Context) error {
	voterId, _ := strconv.ParseInt(c.Param("voterId"), 10, 64)

	voter, err := h.VoterRepository.GetVoter(voterId)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: voter})
}

func (h *handlerV) CreateVoter(c echo.Context) error {

	request := new(voterdto.CreateVoterRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	fmt.Print(request)

	voter := models.Voter{
		ID: request.ID,
		VoterName: request.VoterName,
		PaslonID: request.PaslonID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, err := h.VoterRepository.CreateVoter(voter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertVoterResponse(data)})
}

func (h *handlerV) DeleteVoter(c echo.Context) error {
	voterId, _ := strconv.ParseInt(c.Param("voterId"), 10, 64)

	voter, err := h.VoterRepository.GetVoter(voterId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	
	data, err := h.VoterRepository.DeleteVoter(voter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func convertVoterResponse(p models.Voter) *voterdto.VoterResponse {
	return &voterdto.VoterResponse{
		ID: p.ID,
		VoterName: p.VoterName,
		PaslonID: p.PaslonID,
	}
}