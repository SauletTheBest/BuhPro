package handler

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"

    "buhpro/internal/delivery/http/request"
    "buhpro/internal/delivery/http/response"
    "buhpro/internal/models"
    "buhpro/internal/services"
)

type OrderHandler struct {
    service services.OrderService // используем интерфейс
}

func NewOrderHandler(service services.OrderService) *OrderHandler {
    return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
    var req request.CreateOrderRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
        return
    }

    userID, err := request.GetUserID(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to get user ID: " + err.Error()})
        return
    }

    order := &models.Order{
        ID:          "", // будет сгенерирован в UseCase
        Title:       req.Title,
        Description: req.Description,
        Deadline:    req.Deadline,
        Category:    req.Category,
        Region:      req.Region,
        Status:      "active",
        ClientID:    userID,
        CreatedAt:   time.Now(),
    }

    if err := h.service.CreateOrder(c.Request.Context(), order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order: " + err.Error()})
        return
    }

    c.JSON(http.StatusCreated, response.MapToOrderResponse(order))
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
    orders, err := h.service.GetAllOrders(context.Background())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var responses []*response.OrderResponse
    for _, o := range orders {
        responses = append(responses, response.MapToOrderResponse(o))
    }

    c.JSON(http.StatusOK, responses)
}

func (h *OrderHandler) GetOrderById(c *gin.Context) {
    id := c.Param("id")
    order, err := h.service.GetOrderById(context.Background(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    c.JSON(http.StatusOK, response.MapToOrderResponse(order))
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
    id := c.Param("id")
    var req request.UpdateOrderRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    existing, err := h.service.GetOrderById(context.Background(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    existing.Title = req.Title
    existing.Description = req.Description
    existing.Deadline = req.Deadline
    existing.Category = req.Category
    existing.Region = req.Region
    existing.Status = req.Status

    if err := h.service.UpdateOrder(context.Background(), existing); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response.MapToOrderResponse(existing))
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
    id := c.Param("id")
    if err := h.service.DeleteOrder(context.Background(), id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    c.Status(http.StatusNoContent)
}