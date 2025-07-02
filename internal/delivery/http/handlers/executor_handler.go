package handlers

import (
	"net/http"

	"BuhPro+/internal/delivery/http/requests"
	responses "BuhPro+/internal/delivery/http/response"
	"BuhPro+/internal/domain"
	"BuhPro+/internal/usecase"
	"BuhPro+/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ExecutorHandler struct {
	usecase  *usecase.ExecutorUsecase
	validate *validator.Validate
	logger   *logrus.Logger
}

func NewExecutorHandler(u *usecase.ExecutorUsecase, logger *logrus.Logger) *ExecutorHandler {
	return &ExecutorHandler{
		usecase:  u,
		validate: validator.New(),
		logger:   logger,
	}
}

func (h *ExecutorHandler) RegisterExecutor(c *gin.Context) {
	var req requests.ExecutorRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request format for executor registration")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.logger.WithError(err).Warn("Validation failed for executor registration")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "validation failed", Details: utils.CustomValidationErrors(validationErrors)})
		return
	}

	executor := &domain.Executor{
		Email:           req.Email,
		PasswordHash:    req.Password, // Пароль будет хеширован в usecase
		Name:            req.Name,
		Surname:         req.Surname,
		Patronymic:      req.Patronymic,
		IIN:             req.IIN,
		PhoneNumber:     req.PhoneNumber,
		City:            req.City,
		ExpWork:         req.ExpWork,
		Specializations: req.Specializations,
		Education:       req.Education,
		WorkFormat:      req.WorkFormat,
		HourlyRate:      req.HourlyRate,
		AboutExecutor:   req.AboutExecutor,
	}

	err := h.usecase.RegisterExecutor(executor)
	if err != nil {
		h.logger.WithError(err).Warn("Executor registration failed")
		c.JSON(http.StatusConflict, responses.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("Executor registered successfully")
	c.JSON(http.StatusCreated, responses.AuthSuccessResponse{
		Status:  "success",
		Message: "executor registered successfully",
	})
}

func (h *ExecutorHandler) LoginExecutor(c *gin.Context) {
	var req requests.AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request format for executor login")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.logger.WithError(err).Warn("Validation failed for executor login")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "validation failed", Details: utils.CustomValidationErrors(validationErrors)})
		return
	}

	accessToken, refreshToken, err := h.usecase.LoginExecutor(req.Email, req.Password)
	if err != nil {
		h.logger.WithError(err).Warn("Executor login failed")
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("Executor logged in successfully")
	c.JSON(http.StatusOK, responses.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *ExecutorHandler) RefreshExecutor(c *gin.Context) {
	var req requests.RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request format for executor token refresh")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.logger.WithError(err).Warn("Validation failed for executor token refresh")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "validation failed", Details: utils.CustomValidationErrors(validationErrors)})
		return
	}

	accessToken, err := h.usecase.RefreshExecutorToken(req.RefreshToken)
	if err != nil {
		h.logger.WithError(err).Warn("Executor token refresh failed")
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("Executor token refreshed successfully")
	c.JSON(http.StatusOK, responses.TokenRefreshResponse{
		AccessToken: accessToken,
	})
}

func (h *ExecutorHandler) GetExecutorProfile(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		h.logger.Warn("Executor not authenticated")
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "Executor not authenticated"})
		return
	}

	executor, err := h.usecase.GetExecutorByID(userID.(string))
	if err != nil {
		h.logger.WithError(err).Warn("Executor not found")
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Executor not found"})
		return
	}

	h.logger.Info("Executor profile retrieved successfully")
	c.JSON(http.StatusOK, responses.ExecutorProfileResponse{
		ID:              executor.ID,
		Name:            executor.Name,
		Surname:         executor.Surname,
		Patronymic:      executor.Patronymic,
		IIN:             executor.IIN,
		PhoneNumber:     executor.PhoneNumber,
		Email:           executor.Email,
		City:            executor.City,
		ExpWork:         executor.ExpWork,
		Specializations: executor.Specializations,
		Education:       executor.Education,
		WorkFormat:      executor.WorkFormat,
		HourlyRate:      executor.HourlyRate,
		AboutExecutor:   executor.AboutExecutor,
	})
}
