package services

import (
    "context"
    "buhpro/internal/models"
    "buhpro/internal/repository"
)

// Интерфейс
type OrderService interface {
    CreateOrder(ctx context.Context, order *models.Order) error
    GetAllOrders(ctx context.Context) ([]*models.Order, error)
    GetOrderById(ctx context.Context, id string) (*models.Order, error)
    UpdateOrder(ctx context.Context, order *models.Order) error
    DeleteOrder(ctx context.Context, id string) error
}

// Структура
type OrderServiceImpl struct {
    repo repository.OrderRepository
}

// Конструктор
func NewOrderService(repo repository.OrderRepository) OrderService {
    return &OrderServiceImpl{repo: repo}
}

// Методы реализации
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, order *models.Order) error {
    return s.repo.Create(ctx, order)
}

func (s *OrderServiceImpl) GetAllOrders(ctx context.Context) ([]*models.Order, error) {
    return s.repo.GetAll(ctx)
}

func (s *OrderServiceImpl) GetOrderById(ctx context.Context, id string) (*models.Order, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, order *models.Order) error {
    return s.repo.Update(ctx, order)
}

func (s *OrderServiceImpl) DeleteOrder(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}