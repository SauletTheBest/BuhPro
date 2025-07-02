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

type ExecutorUsecase struct {
	executorRepo repository.ExecutorRepository
	jwtSecret    string
	logger       *logrus.Logger
}

func NewExecutorUsecase(repo repository.ExecutorRepository, secret string, logger *logrus.Logger) *ExecutorUsecase {
	return &ExecutorUsecase{repo, secret, logger}
}

// internal/usecase/executor_usecase.go
// ... (остальная часть файла)

func (s *ExecutorUsecase) RegisterExecutor(executor *domain.Executor) error { // Изменена сигнатура
	s.logger.WithFields(logrus.Fields{
		"email": executor.Email,
	}).Info("Attempting to register executor")

	_, err := s.executorRepo.GetByEmail(executor.Email)
	if err == nil {
		s.logger.Warn("Executor already exists")
		return errors.New("executor with this email already exists")
	}

	if !utils.IsPasswordComplex(executor.PasswordHash) {
		s.logger.Warn("Password does not meet complexity requirements")
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(executor.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		s.logger.WithError(err).Error("Failed to hash password")
		return err
	}
	executor.PasswordHash = string(hashed) // Хешируем пароль

	if err := s.executorRepo.Create(executor); err != nil {
		s.logger.WithError(err).Error("Failed to create executor")
		return err
	}

	s.logger.Info("Executor registered successfully")
	return nil
}

// ... (остальная часть файла)
func (s *ExecutorUsecase) LoginExecutor(email, password string) (string, string, error) {
	s.logger.WithFields(logrus.Fields{
		"email": email,
	}).Info("Attempting to login executor")

	executor, err := s.executorRepo.GetByEmail(email)
	if err != nil {
		s.logger.Warn("Invalid email or password for executor login")
		return "", "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(executor.PasswordHash), []byte(password)); err != nil {
		s.logger.Warn("Invalid email or password for executor login")
		return "", "", errors.New("invalid email or password")
	}

	accessToken, err := utils.GenerateToken(executor.ID, s.jwtSecret)
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate access token for executor")
		return "", "", err
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate refresh token for executor")
		return "", "", err
	}

	refreshTokenModel := &domain.RefreshToken{
		ID:        uuid.New().String(),
		UserID:    executor.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.executorRepo.CreateRefreshToken(refreshTokenModel); err != nil {
		s.logger.WithError(err).Error("Failed to save refresh token for executor")
		return "", "", err
	}

	s.logger.Info("Executor logged in successfully")
	return accessToken, refreshToken, nil
}

func (s *ExecutorUsecase) RefreshExecutorToken(refreshToken string) (string, error) {
	s.logger.Info("Attempting to refresh executor token")

	token, err := s.executorRepo.GetRefreshToken(refreshToken)
	if err != nil {
		s.logger.WithError(err).Warn("Invalid refresh token for executor")
		return "", errors.New("invalid refresh token")
	}

	if time.Now().After(token.ExpiresAt) {
		s.logger.Warn("Refresh token expired for executor")
		return "", errors.New("refresh token expired")
	}

	executor, err := s.executorRepo.GetByID(token.UserID)
	if err != nil {
		s.logger.WithError(err).Error("Executor not found for refresh token")
		return "", errors.New("executor not found")
	}

	accessToken, err := utils.GenerateToken(executor.ID, s.jwtSecret)
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate new access token for executor")
		return "", err
	}

	s.logger.Info("Executor token refreshed successfully")
	return accessToken, nil
}

func (s *ExecutorUsecase) GetExecutorByID(id string) (*domain.Executor, error) {
	executor, err := s.executorRepo.GetByID(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get executor by ID")
		return nil, err
	}
	return executor, nil
}
