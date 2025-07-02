package repository

import (
    "context"
    "buhpro/internal/models"
    "gorm.io/gorm"
)

// OrderRepositoryImpl - реализация интерфейса OrderRepository
type OrderRepositoryImpl struct {
    db *gorm.DB
}

// NewOrderRepository - конструктор для OrderRepository
func NewOrderRepository(db *gorm.DB) OrderRepository {
    return &OrderRepositoryImpl{db: db}
}

// Create - создание нового заказа
func (r *OrderRepositoryImpl) Create(ctx context.Context, order *models.Order) error {
    return r.db.WithContext(ctx).Create(order).Error
}

// GetAll - получение всех заказов
func (r *OrderRepositoryImpl) GetAll(ctx context.Context) ([]*models.Order, error) {
    var orders []*models.Order
    err := r.db.WithContext(ctx).Find(&orders).Error
    return orders, err
}

// GetByID - получение заказа по ID
func (r *OrderRepositoryImpl) GetByID(ctx context.Context, id string) (*models.Order, error) {
    var order models.Order
    err := r.db.WithContext(ctx).First(&order, "id = ?", id).Error
    return &order, err
}

// Update - обновление заказа
func (r *OrderRepositoryImpl) Update(ctx context.Context, order *models.Order) error {
    return r.db.WithContext(ctx).Save(order).Error
}

// Delete - удаление заказа по ID
func (r *OrderRepositoryImpl) Delete(ctx context.Context, id string) error {
    return r.db.WithContext(ctx).Delete(&models.Order{}, "id = ?", id).Error
}