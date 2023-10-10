package handlers

import (
	paslonsdto "myapp/dto/paslons"
	dto "myapp/dto/result"
	"myapp/models"
	"myapp/repositories"
	cloudinaryconfig "myapp/utils/cloudinary-upload"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type handler struct {
	PaslonRepository repositories.PaslonRepository
}

func HandlerPaslon(PaslonRepository repositories.PaslonRepository) *handler {
	return &handler{PaslonRepository}
}

func (h *handler) FindPaslons(c echo.Context) error {
	paslons, err := h.PaslonRepository.FindPaslons()
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: paslons})
}

func (h *handler) GetPaslon(c echo.Context) error {
	paslonId, _ := strconv.Atoi(c.Param("paslonId"))

	paslon, err := h.PaslonRepository.GetPaslon(paslonId)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.ErrorResult{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: paslon})
}

func (h* handler) CreatePaslon(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	fileHeader, ok := form.File["image"]
	if !ok || len(fileHeader) == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "File 'image' harus diunggah"})
	}
	imageFile, err := fileHeader[0].Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	defer imageFile.Close()

	paslonName := form.Value["name"][0]
	paslonVision := form.Value["vision"][0]
	folderName := "Pemilu"
	imageUrl, err := cloudinaryconfig.UploadToCloudinary(imageFile, folderName, paslonName)
	if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
    }

	paslon := models.Paslon {
		Name: paslonName,
		Vision: paslonVision,
		Image: imageUrl,
		PostedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, err := h.PaslonRepository.CreatePaslon(paslon)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handler) UpdatePaslon(c echo.Context) error {

	paslonId, _ := strconv.Atoi(c.Param("paslonId"))
	paslon, err := h.PaslonRepository.GetPaslon(paslonId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	fileHeader, ok := form.File["image"]
	if !ok || len(fileHeader) == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "File 'image' harus diunggah"})
	}
	imageFile, err := fileHeader[0].Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	defer imageFile.Close()

	paslonName := form.Value["name"][0]
	paslonVision := form.Value["vision"][0]
	folderName := "Pemilu"
	imageUrl, err := cloudinaryconfig.UploadToCloudinary(imageFile, folderName, paslonName)
	if err != nil {
        return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
    }

	if paslonName != "" {
		paslon.Name = paslonName
	}

	if paslonVision != "" {
		paslon.Vision = paslonVision
	}

	if imageUrl != "" {
		paslon.Image = imageUrl
	}

	paslon.UpdatedAt = time.Now()

	data, err := h.PaslonRepository.UpdatePaslon(paslon, paslonId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}) 
}

func (h *handler) DeletePaslon(c echo.Context) error {
	paslonId, _ := strconv.Atoi(c.Param("paslonId"))
	
	paslon, err := h.PaslonRepository.GetPaslon(paslonId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	cloudinaryconfig.DeleteCloudinary("Pemilu", paslon.Name)

	data, err := h.PaslonRepository.DeletePaslon(paslon, paslonId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func convertResponse(p models.Paslon) *paslonsdto.PaslonResponse {
	return &paslonsdto.PaslonResponse{
		ID: p.ID,
		Name: p.Name,
		Vision: p.Vision,
		Image: p.Image,
	}
}