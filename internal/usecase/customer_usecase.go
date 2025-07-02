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

type CustomerUsecase struct {
	customerRepo repository.CustomerRepository
	jwtSecret    string
	logger       *logrus.Logger
}

func NewCustomerUsecase(repo repository.CustomerRepository, secret string, logger *logrus.Logger) *CustomerUsecase {
	return &CustomerUsecase{repo, secret, logger}
}

// internal/usecase/customer_usecase.go
// ... (остальная часть файла)

func (s *CustomerUsecase) RegisterCustomer(customer *domain.Customer) error { // Изменена сигнатура
	s.logger.WithFields(logrus.Fields{
		"email": customer.Email,
	}).Info("Attempting to register customer")

	_, err := s.customerRepo.GetByEmail(customer.Email)
	if err == nil {
		s.logger.Warn("Customer already exists")
		return errors.New("customer with this email already exists")
	}

	if !utils.IsPasswordComplex(customer.PasswordHash) { // Проверяем переданный пароль
		s.logger.Warn("Password does not meet complexity requirements")
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(customer.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		s.logger.WithError(err).Error("Failed to hash password")
		return err
	}
	customer.PasswordHash = string(hashed) // Хешируем пароль и сохраняем в модели

	if err := s.customerRepo.Create(customer); err != nil {
		s.logger.WithError(err).Error("Failed to create customer")
		return err
	}

	s.logger.Info("Customer registered successfully")
	return nil
}

// ... (остальная часть файла)
func (s *CustomerUsecase) LoginCustomer(email, password string) (string, string, error) {
	s.logger.WithFields(logrus.Fields{
		"email": email,
	}).Info("Attempting to login customer")

	customer, err := s.customerRepo.GetByEmail(email)
	if err != nil {
		s.logger.Warn("Invalid email or password for customer login")
		return "", "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.PasswordHash), []byte(password)); err != nil {
		s.logger.Warn("Invalid email or password for customer login")
		return "", "", errors.New("invalid email or password")
	}

	accessToken, err := utils.GenerateToken(customer.ID, s.jwtSecret)
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate access token for customer")
		return "", "", err
	}

	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate refresh token for customer")
		return "", "", err
	}

	refreshTokenModel := &domain.RefreshToken{
		ID:        uuid.New().String(),
		UserID:    customer.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.customerRepo.CreateRefreshToken(refreshTokenModel); err != nil {
		s.logger.WithError(err).Error("Failed to save refresh token for customer")
		return "", "", err
	}

	s.logger.Info("Customer logged in successfully")
	return accessToken, refreshToken, nil
}

func (s *CustomerUsecase) RefreshCustomerToken(refreshToken string) (string, error) {
	s.logger.Info("Attempting to refresh customer token")

	token, err := s.customerRepo.GetRefreshToken(refreshToken)
	if err != nil {
		s.logger.WithError(err).Warn("Invalid refresh token for customer")
		return "", errors.New("invalid refresh token")
	}

	if time.Now().After(token.ExpiresAt) {
		s.logger.Warn("Refresh token expired for customer")
		return "", errors.New("refresh token expired")
	}

	customer, err := s.customerRepo.GetByID(token.UserID)
	if err != nil {
		s.logger.WithError(err).Error("Customer not found for refresh token")
		return "", errors.New("customer not found")
	}

	accessToken, err := utils.GenerateToken(customer.ID, s.jwtSecret)
	if err != nil {
		s.logger.WithError(err).Error("Failed to generate new access token for customer")
		return "", err
	}

	s.logger.Info("Customer token refreshed successfully")
	return accessToken, nil
}

func (s *CustomerUsecase) GetCustomerByID(id string) (*domain.Customer, error) {
	customer, err := s.customerRepo.GetByID(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to get customer by ID")
		return nil, err
	}
	return customer, nil
}
