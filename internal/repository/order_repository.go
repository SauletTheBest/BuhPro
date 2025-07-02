package repository

import (
    "context"
    "buhpro/internal/models"
)

type OrderRepository interface {
    Create(ctx context.Context, order *models.Order) error
    GetAll(ctx context.Context) ([]*models.Order, error)
    GetByID(ctx context.Context, id string) (*models.Order, error)
    Update(ctx context.Context, order *models.Order) error
    Delete(ctx context.Context, id string) error
}

