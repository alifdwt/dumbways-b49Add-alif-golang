package handlers

import (
	"fmt"
	dto "myapp/dto/result"
	userdto "myapp/dto/users"
	"myapp/models"
	"myapp/repositories"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type handlerU struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerU {
	return &handlerU{UserRepository}
}

func (h *handlerU) Register(c echo.Context) error {
	request := new(userdto.CreateUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	checkEmail, err := h.UserRepository.GetUserByEmail(request.Email)
	if checkEmail.Email != "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Email is already registered"})
	}
	if err != nil {
		fmt.Println("Gak taulah")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		// VoterID: request.VoterID,
		FullName: request.FullName,
		Email: request.Email,
		Password: string(password),
		PostedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data, err := h.UserRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertUserResponse(data)})
}

func (h *handlerU) Login(c echo.Context) error {
	request := new(userdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	if err := validation.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	checkEmail, err := h.UserRepository.GetUserByEmail(request.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Email / password is wrong!"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkEmail.Password), []byte(request.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Your email / password is wrong!"})
	}
	
	user := models.User{
		ID:       checkEmail.ID,
		FullName: checkEmail.FullName,
		Email:    checkEmail.Email,
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Failed to create token"})
	}

	return c.JSON(http.StatusOK, dto.LoginResult{Code: http.StatusOK, Data: convertUserResponse(user), Token: tokenString})
}

func (h *handlerU) Check(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "unauthorized",
		})
	}

	// Pastikan token memiliki format "Bearer <token>"
	tokenParts := strings.Split(tokenString, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "unauthorized",
		})
	}

	tokenString = tokenParts[1]

	// Validasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ganti dengan kunci rahasia yang sesuai
		return []byte("secret-jwt-token"), nil
	})

	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "unauthorized",
		})
	}

	// Jika token valid, Anda dapat mengakses informasi yang ada di dalamnya
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "internal server error",
		})
	}

	userID, ok := claims["user_id"].(float64) // Sesuaikan dengan nama klaim yang Anda gunakan dalam token
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "internal server error",
		})
	}
	fmt.Println(userID)

	// Lakukan operasi yang sesuai dengan informasi yang Anda dapatkan dari token
	// Contoh: Cari pengguna berdasarkan userID
	// user, err := UserRepository.FindUserByID(int(userID))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Token is valid!",
	})
}


func convertUserResponse(p models.User) *userdto.UserResponse {
	return &userdto.UserResponse{
		ID: p.ID,
		// VoterID: p.VoterID,
		FullName: p.FullName,
		Email: p.Email,
	}
}