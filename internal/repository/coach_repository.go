package repository

import (
	"BuhPro+/internal/domain" // Обновлен импорт

	"gorm.io/gorm"
)

type CoachRepository interface {
	Create(coach *domain.Coach) error
	GetByEmail(email string) (*domain.Coach, error)
	GetByID(id string) (*domain.Coach, error)
	CreateRefreshToken(token *domain.RefreshToken) error
	GetRefreshToken(token string) (*domain.RefreshToken, error)
}

type coachRepository struct {
	db *gorm.DB
}

func NewCoachRepository(db *gorm.DB) CoachRepository {
	return &coachRepository{db}
}

func (r *coachRepository) Create(coach *domain.Coach) error {
	return r.db.Create(coach).Error
}

func (r *coachRepository) GetByEmail(email string) (*domain.Coach, error) {
	var coach domain.Coach
	err := r.db.Where("email = ?", email).First(&coach).Error
	return &coach, err
}

func (r *coachRepository) GetByID(id string) (*domain.Coach, error) {
	var coach domain.Coach
	err := r.db.First(&coach, "id = ?", id).Error
	return &coach, err
}

func (r *coachRepository) CreateRefreshToken(token *domain.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *coachRepository) GetRefreshToken(token string) (*domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken
	err := r.db.Where("token = ?", token).First(&refreshToken).Error
	return &refreshToken, err
}
