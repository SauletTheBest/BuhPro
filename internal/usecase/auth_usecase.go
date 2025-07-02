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

type AuthUsecase struct {
	userRepo  repository.UserRepository
	jwtSecret string
	logger    *logrus.Logger
}

func NewAuthUsecase(repo repository.UserRepository, secret string, logger *logrus.Logger) *AuthUsecase {
	return &AuthUsecase{repo, secret, logger}
}

func (s *AuthUsecase) Register(email, password string) error {
	s.logger.WithFields(logrus.Fields{
		"email": email,
	}).Info("Attempting to register user")

	_, err := s.userRepo.GetByEmail(email)
	if err == nil {
		s.logger.Warn("User already exists")
		return errors.New("user already exists")
	}

	if !utils.IsPasswordComplex(password) {
		s.logger.Warn("Password does not meet complexity requirements")
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.WithError(err).Error("Failed to hash password")
		return err
	}

	user := &domain.User{
		Email:        email,
		PasswordHash: string(hashed),
	}

	if err := s.userRepo.Create(user); err != nil {
		s.logger.WithError(err).Error("Failed to create user")
		return err
	}

	s.logger.Info("User registered successfully")
	return nil
}

func (s *AuthUsecase) Login(email, password string) (string, string, error) {
	s.logger.WithFields(logrus.Fields{
		"email": email,
	}).Info("Attempting to login user")

	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		s.logger.Warn("Invalid email or password")
		return "", "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		s.logger.Warn("Invalid email or password")
		return "", "", errors.New("invalid email or password")
	}

	accessToken, err := utils.GenerateToken(user.ID, s.jwtSecret)
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate access token")
		return "", "", err
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate refresh token")
		return "", "", err
	}

	refreshTokenModel := &domain.RefreshToken{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.userRepo.CreateRefreshToken(refreshTokenModel); err != nil {
		s.logger.WithError(err).Error("Failed to save refresh token")
		return "", "", err
	}

	s.logger.Info("User logged in successfully")
	return accessToken, refreshToken, nil
}

func (s *AuthUsecase) RefreshToken(refreshToken string) (string, error) {
	s.logger.Info("Attempting to refresh token")

	token, err := s.userRepo.GetRefreshToken(refreshToken)
	if err != nil {
		s.logger.WithError(err).Warn("Invalid refresh token")
		return "", errors.New("invalid refresh token")
	}

	if time.Now().After(token.ExpiresAt) {
		s.logger.Warn("Refresh token expired")
		return "", errors.New("refresh token expired")
	}

	user, err := s.userRepo.GetByID(token.UserID)
	if err != nil {
		s.logger.WithError(err).Error("User not found for refresh token")
		return "", errors.New("user not found")
	}

	accessToken, err := utils.GenerateToken(user.ID, s.jwtSecret)
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate new access token")
		return "", err
	}

	s.logger.Info("Token refreshed successfully")
	return accessToken, nil
}

func (s *AuthUsecase) GetUserByID(id string) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get user by ID")
		return nil, err
	}
	return user, nil
}
