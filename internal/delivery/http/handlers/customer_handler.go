package handlers

import (
	"net/http"

	"BuhPro+/internal/delivery/http/requests"
	responses "BuhPro+/internal/delivery/http/response"
	"BuhPro+/internal/domain" // Добавлено для создания полной модели
	"BuhPro+/internal/usecase"
	"BuhPro+/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type CustomerHandler struct {
	usecase  *usecase.CustomerUsecase
	validate *validator.Validate
	logger   *logrus.Logger
}

func NewCustomerHandler(u *usecase.CustomerUsecase, logger *logrus.Logger) *CustomerHandler {
	return &CustomerHandler{
		usecase:  u,
		validate: validator.New(),
		logger:   logger,
	}
}

func (h *CustomerHandler) RegisterCustomer(c *gin.Context) {
	var req requests.CustomerRegisterRequest // Используем CustomerRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request format for customer registration")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.logger.WithError(err).Warn("Validation failed for customer registration")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "validation failed", Details: utils.CustomValidationErrors(validationErrors)})
		return
	}

	// Здесь мы передаем все поля в usecase
	customer := &domain.Customer{
		Email:           req.Email,
		PasswordHash:    req.Password, // Пароль будет хеширован в usecase
		ClientType:      req.ClientType,
		CompanyName:     req.CompanyName,
		IIN:             req.IIN,
		Name:            req.Name,
		JobPosition:     req.JobPosition,
		PhoneNumber:     req.PhoneNumber,
		Address:         req.Address,
		WorkDescription: req.WorkDescription,
	}

	err := h.usecase.RegisterCustomer(customer) // Передаем полную модель Customer
	if err != nil {
		h.logger.WithError(err).Warn("Customer registration failed")
		c.JSON(http.StatusConflict, responses.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("Customer registered successfully")
	c.JSON(http.StatusCreated, responses.AuthSuccessResponse{
		Status:  "success",
		Message: "customer registered successfully",
	})
}

func (h *CustomerHandler) LoginCustomer(c *gin.Context) {
	var req requests.AuthRequest // Для логина достаточно email и password

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request format for customer login")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.logger.WithError(err).Warn("Validation failed for customer login")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "validation failed", Details: utils.CustomValidationErrors(validationErrors)})
		return
	}

	accessToken, refreshToken, err := h.usecase.LoginCustomer(req.Email, req.Password)
	if err != nil {
		h.logger.WithError(err).Warn("Customer login failed")
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("Customer logged in successfully")
	c.JSON(http.StatusOK, responses.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *CustomerHandler) RefreshCustomer(c *gin.Context) {
	var req requests.RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request format for customer token refresh")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.logger.WithError(err).Warn("Validation failed for customer token refresh")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "validation failed", Details: utils.CustomValidationErrors(validationErrors)})
		return
	}

	accessToken, err := h.usecase.RefreshCustomerToken(req.RefreshToken)
	if err != nil {
		h.logger.WithError(err).Warn("Customer token refresh failed")
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("Customer token refreshed successfully")
	c.JSON(http.StatusOK, responses.TokenRefreshResponse{
		AccessToken: accessToken,
	})
}

func (h *CustomerHandler) GetCustomerProfile(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		h.logger.Warn("Customer not authenticated")
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "Customer not authenticated"})
		return
	}

	customer, err := h.usecase.GetCustomerByID(userID.(string))
	if err != nil {
		h.logger.WithError(err).Warn("Customer not found")
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Customer not found"})
		return
	}

	h.logger.Info("Customer profile retrieved successfully")
	c.JSON(http.StatusOK, responses.CustomerProfileResponse{
		ID:              customer.ID,
		ClientType:      customer.ClientType,
		CompanyName:     customer.CompanyName,
		IIN:             customer.IIN,
		Name:            customer.Name,
		JobPosition:     customer.JobPosition,
		PhoneNumber:     customer.PhoneNumber,
		Email:           customer.Email,
		Address:         customer.Address,
		WorkDescription: customer.WorkDescription,
	})
}
