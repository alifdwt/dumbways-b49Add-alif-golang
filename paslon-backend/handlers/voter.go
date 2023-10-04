package handlers

import (
	dto "myapp/dto/result"
	voterdto "myapp/dto/voters"
	"myapp/models"
	"myapp/repositories"
	"net/http"
	"strconv"
	"time"

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
	voterId, _ := strconv.Atoi(c.Param("voterId"))

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
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	voterName := form.Value["voter_name"][0]
	voterPaslonId, _ := strconv.ParseInt(form.Value["paslon_id"][0], 10, 10)

	voter := models.Voter {
		VoterName: voterName,
		PaslonID: int(voterPaslonId),
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
	voterId, _ := strconv.Atoi(c.Param("voterId"))

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