package usecase

import (
	"errors"
	"time"

	"BuhPro+/internal/domain"
	"BuhPro+/internal/repository"
	"BuhPro+/internal/utils" // Обновлен импорт

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type CoachUsecase struct {
	coachRepo repository.CoachRepository
	jwtSecret string
	logger    *logrus.Logger
}

func NewCoachUsecase(repo repository.CoachRepository, secret string, logger *logrus.Logger) *CoachUsecase {
	return &CoachUsecase{repo, secret, logger}
}

// internal/usecase/coach_usecase.go
// ... (остальная часть файла)

func (s *CoachUsecase) RegisterCoach(coach *domain.Coach) error { // Изменена сигнатура
	s.logger.WithFields(logrus.Fields{
		"email": coach.Email,
	}).Info("Attempting to register coach")

	_, err := s.coachRepo.GetByEmail(coach.Email)
	if err == nil {
		s.logger.Warn("Coach already exists")
		return errors.New("coach with this email already exists")
	}

	if !utils.IsPasswordComplex(coach.PasswordHash) {
		s.logger.Warn("Password does not meet complexity requirements")
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(coach.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		s.logger.WithError(err).Error("Failed to hash password")
		return err
	}
	coach.PasswordHash = string(hashed) // Хешируем пароль

	if err := s.coachRepo.Create(coach); err != nil {
		s.logger.WithError(err).Error("Failed to create coach")
		return err
	}

	s.logger.Info("Coach registered successfully")
	return nil
}

// ... (остальная часть файла)

func (s *CoachUsecase) LoginCoach(email, password string) (string, string, error) {
	s.logger.WithFields(logrus.Fields{
		"email": email,
	}).Info("Attempting to login coach")

	coach, err := s.coachRepo.GetByEmail(email)
	if err != nil {
		s.logger.Warn("Invalid email or password for coach login")
		return "", "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(coach.PasswordHash), []byte(password)); err != nil {
		s.logger.Warn("Invalid email or password for coach login")
		return "", "", errors.New("invalid email or password")
	}

	accessToken, err := utils.GenerateToken(coach.ID, s.jwtSecret)
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate access token for coach")
		return "", "", err
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate refresh token for coach")
		return "", "", err
	}

	refreshTokenModel := &domain.RefreshToken{
		ID:        uuid.New().String(),
		UserID:    coach.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.coachRepo.CreateRefreshToken(refreshTokenModel); err != nil {
		s.logger.WithError(err).Error("Failed to save refresh token for coach")
		return "", "", err
	}

	s.logger.Info("Coach logged in successfully")
	return accessToken, refreshToken, nil
}

func (s *CoachUsecase) RefreshCoachToken(refreshToken string) (string, error) {
	s.logger.Info("Attempting to refresh coach token")

	token, err := s.coachRepo.GetRefreshToken(refreshToken)
	if err != nil {
		s.logger.WithError(err).Warn("Invalid refresh token for coach")
		return "", errors.New("invalid refresh token")
	}

	if time.Now().After(token.ExpiresAt) {
		s.logger.Warn("Refresh token expired for coach")
		return "", errors.New("refresh token expired")
	}

	coach, err := s.coachRepo.GetByID(token.UserID)
	if err != nil {
		s.logger.WithError(err).Error("Coach not found for refresh token")
		return "", errors.New("coach not found")
	}

	accessToken, err := utils.GenerateToken(coach.ID, s.jwtSecret)
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate new access token for coach")
		return "", err
	}

	s.logger.Info("Coach token refreshed successfully")
	return accessToken, nil
}

func (s *CoachUsecase) GetCoachByID(id string) (*domain.Coach, error) {
	coach, err := s.coachRepo.GetByID(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get coach by ID")
		return nil, err
	}
	return coach, nil
}
