package handlers

import (
	"net/http"

	"BuhPro+/internal/delivery/http/requests"           // Обновлен импорт
	responses "BuhPro+/internal/delivery/http/response" // Обновлен импорт
	"BuhPro+/internal/usecase"                          // Обновлен импорт
	"BuhPro+/internal/utils"                            // Обновлен импорт

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	usecase  *usecase.AuthUsecase
	validate *validator.Validate
	logger   *logrus.Logger
}

func NewAuthHandler(u *usecase.AuthUsecase, logger *logrus.Logger) *AuthHandler {
	return &AuthHandler{
		usecase:  u,
		validate: validator.New(),
		logger:   logger,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req requests.AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request format")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.logger.WithError(err).Warn("Validation failed")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "validation failed", Details: utils.CustomValidationErrors(validationErrors)})
		return
	}

	err := h.usecase.Register(req.Email, req.Password)
	if err != nil {
		h.logger.WithError(err).Warn("Registration failed")
		c.JSON(http.StatusConflict, responses.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("User registered successfully")
	c.JSON(http.StatusCreated, responses.AuthSuccessResponse{
		Status:  "success",
		Message: "user registered successfully",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req requests.AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request format")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.logger.WithError(err).Warn("Validation failed")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "validation failed", Details: utils.CustomValidationErrors(validationErrors)})
		return
	}

	accessToken, refreshToken, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		h.logger.WithError(err).Warn("Login failed")
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("User logged in successfully")
	c.JSON(http.StatusOK, responses.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var req requests.RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request format")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.logger.WithError(err).Warn("Validation failed")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "validation failed", Details: utils.CustomValidationErrors(validationErrors)})
		return
	}

	accessToken, err := h.usecase.RefreshToken(req.RefreshToken)
	if err != nil {
		h.logger.WithError(err).Warn("Token refresh failed")
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("Token refreshed successfully")
	c.JSON(http.StatusOK, responses.TokenRefreshResponse{
		AccessToken: accessToken,
	})
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		h.logger.Warn("User not authenticated")
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "User not authenticated"})
		return
	}

	user, err := h.usecase.GetUserByID(userID.(string))
	if err != nil {
		h.logger.WithError(err).Warn("User not found")
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "User not found"})
		return
	}

	h.logger.Info("User profile retrieved successfully")
	c.JSON(http.StatusOK, responses.UserProfileResponse{
		ID:    user.ID,
		Email: user.Email,
	})
}
