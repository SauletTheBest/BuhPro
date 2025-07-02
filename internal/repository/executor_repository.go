package repository

import (
	"BuhPro+/internal/domain" // Обновлен импорт

	"gorm.io/gorm"
)

type ExecutorRepository interface {
	Create(executor *domain.Executor) error
	GetByEmail(email string) (*domain.Executor, error)
	GetByID(id string) (*domain.Executor, error)
	CreateRefreshToken(token *domain.RefreshToken) error
	GetRefreshToken(token string) (*domain.RefreshToken, error)
}

type executorRepository struct {
	db *gorm.DB
}

func NewExecutorRepository(db *gorm.DB) ExecutorRepository {
	return &executorRepository{db}
}

func (r *executorRepository) Create(executor *domain.Executor) error {
	return r.db.Create(executor).Error
}

func (r *executorRepository) GetByEmail(email string) (*domain.Executor, error) {
	var executor domain.Executor
	err := r.db.Where("email = ?", email).First(&executor).Error
	return &executor, err
}

func (r *executorRepository) GetByID(id string) (*domain.Executor, error) {
	var executor domain.Executor
	err := r.db.First(&executor, "id = ?", id).Error
	return &executor, err
}

func (r *executorRepository) CreateRefreshToken(token *domain.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *executorRepository) GetRefreshToken(token string) (*domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken
	err := r.db.Where("token = ?", token).First(&refreshToken).Error
	return &refreshToken, err
}
