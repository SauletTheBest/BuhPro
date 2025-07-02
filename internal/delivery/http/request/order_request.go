package request

import (
	"time"

	"github.com/gin-gonic/gin"
    "errors"
)

type CreateOrderRequest struct {
    Title       string    `json:"title" binding:"required"`
    Description string    `json:"description" binding:"required"`
    Deadline    time.Time `json:"deadline" binding:"required"`
    Category    string    `json:"category" binding:"required"`
    Region      string    `json:"region" binding:"required"`
}
type UpdateOrderRequest struct {
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Deadline    time.Time `json:"deadline"`
    Category    string    `json:"category"`
    Region      string    `json:"region"`
    Status      string    `json:"status"` // active, in_progress, closed
}

// GetUserID — извлечение ID пользователя из контекста
func GetUserID(c *gin.Context) (string, error) {
    userID, exists := c.Get("user_id")
    if !exists {
        return "", errors.New("user_id not found in context")
    }
    if userID == nil {
        return "", errors.New("user_id is nil")
    }
    userIDStr, ok := userID.(string)
    if !ok {
        return "", errors.New("user_id is not a string")
    }
    return userIDStr, nil
}