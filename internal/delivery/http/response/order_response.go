// internal/delivery/http/response/order_response.go
package response

import (
    "time"
	"buhpro/internal/models"
)

// OrderResponse — DTO для ответа клиенту
type OrderResponse struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Deadline    string `json:"deadline"`
    Category    string `json:"category"`
    Region      string `json:"region"`
    Status      string `json:"status"`
    ClientID    string `json:"client_id"`
    CreatedAt   string `json:"created_at"`
}

// MapToOrderResponse — преобразует models.Order в JSON DTO
func MapToOrderResponse(order *models.Order) *OrderResponse {
    return &OrderResponse{
        ID:          order.ID,
        Title:       order.Title,
        Description: order.Description,
        Deadline:    order.Deadline.Format(time.RFC3339),
        Category:    order.Category,
        Region:      order.Region,
        Status:      order.Status,
        ClientID:    order.ClientID,
        CreatedAt:   order.CreatedAt.Format(time.RFC3339),
    }
}