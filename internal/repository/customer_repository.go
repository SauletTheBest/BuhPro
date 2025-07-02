package repository

import (
	"BuhPro+/internal/domain" // Обновлен импорт

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *domain.Customer) error
	GetByEmail(email string) (*domain.Customer, error)
	GetByID(id string) (*domain.Customer, error)
	CreateRefreshToken(token *domain.RefreshToken) error
	GetRefreshToken(token string) (*domain.RefreshToken, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) Create(customer *domain.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) GetByEmail(email string) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.db.Where("email = ?", email).First(&customer).Error
	return &customer, err
}

func (r *customerRepository) GetByID(id string) (*domain.Customer, error) {
	var customer domain.Customer
	err := r.db.First(&customer, "id = ?", id).Error
	return &customer, err
}

func (r *customerRepository) CreateRefreshToken(token *domain.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *customerRepository) GetRefreshToken(token string) (*domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken
	err := r.db.Where("token = ?", token).First(&refreshToken).Error
	return &refreshToken, err
}
