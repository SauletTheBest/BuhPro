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

type CoachHandler struct {
	usecase  *usecase.CoachUsecase
	validate *validator.Validate
	logger   *logrus.Logger
}

func NewCoachHandler(u *usecase.CoachUsecase, logger *logrus.Logger) *CoachHandler {
	return &CoachHandler{
		usecase:  u,
		validate: validator.New(),
		logger:   logger,
	}
}

func (h *CoachHandler) RegisterCoach(c *gin.Context) {
	var req requests.CoachRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request format for coach registration")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.logger.WithError(err).Warn("Validation failed for coach registration")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "validation failed", Details: utils.CustomValidationErrors(validationErrors)})
		return
	}

	coach := &domain.Coach{
		Email:                  req.Email,
		PasswordHash:           req.Password, // Пароль будет хеширован в usecase
		Name:                   req.Name,
		Surname:                req.Surname,
		PhoneNumber:            req.PhoneNumber,
		ExpCoach:               req.ExpCoach,
		Specializations:        req.Specializations,
		EducationCertificates:  req.EducationCertificates,
		AchievementsExperience: req.AchievementsExperience,
		Methodology:            req.Methodology,
		AboutCoach:             req.AboutCoach,
	}

	err := h.usecase.RegisterCoach(coach)
	if err != nil {
		h.logger.WithError(err).Warn("Coach registration failed")
		c.JSON(http.StatusConflict, responses.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("Coach registered successfully")
	c.JSON(http.StatusCreated, responses.AuthSuccessResponse{
		Status:  "success",
		Message: "coach registered successfully",
	})
}

func (h *CoachHandler) LoginCoach(c *gin.Context) {
	var req requests.AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request format for coach login")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.logger.WithError(err).Warn("Validation failed for coach login")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "validation failed", Details: utils.CustomValidationErrors(validationErrors)})
		return
	}

	accessToken, refreshToken, err := h.usecase.LoginCoach(req.Email, req.Password)
	if err != nil {
		h.logger.WithError(err).Warn("Coach login failed")
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("Coach logged in successfully")
	c.JSON(http.StatusOK, responses.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *CoachHandler) RefreshCoach(c *gin.Context) {
	var req requests.RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Invalid request format for coach token refresh")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "invalid request format"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		h.logger.WithError(err).Warn("Validation failed for coach token refresh")
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Error: "validation failed", Details: utils.CustomValidationErrors(validationErrors)})
		return
	}

	accessToken, err := h.usecase.RefreshCoachToken(req.RefreshToken)
	if err != nil {
		h.logger.WithError(err).Warn("Coach token refresh failed")
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: err.Error()})
		return
	}

	h.logger.Info("Coach token refreshed successfully")
	c.JSON(http.StatusOK, responses.TokenRefreshResponse{
		AccessToken: accessToken,
	})
}

func (h *CoachHandler) GetCoachProfile(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		h.logger.Warn("Coach not authenticated")
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Error: "Coach not authenticated"})
		return
	}

	coach, err := h.usecase.GetCoachByID(userID.(string))
	if err != nil {
		h.logger.WithError(err).Warn("Coach not found")
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Error: "Coach not found"})
		return
	}

	h.logger.Info("Coach profile retrieved successfully")
	c.JSON(http.StatusOK, responses.CoachProfileResponse{
		ID:                     coach.ID,
		Name:                   coach.Name,
		Surname:                coach.Surname,
		PhoneNumber:            coach.PhoneNumber,
		Email:                  coach.Email,
		ExpCoach:               coach.ExpCoach,
		Specializations:        coach.Specializations,
		EducationCertificates:  coach.EducationCertificates,
		AchievementsExperience: coach.AchievementsExperience,
		Methodology:            coach.Methodology,
		AboutCoach:             coach.AboutCoach,
	})
}
