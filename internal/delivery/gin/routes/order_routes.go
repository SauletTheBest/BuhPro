// internal/delivery/gin/routes/order_routes.go
package routes

import (
    "github.com/gin-gonic/gin"
    "buhpro/internal/delivery/http/handler"
)

func OrderRoutes(router *gin.RouterGroup, handler *handler.OrderHandler) {
    router.POST("/orders", handler.CreateOrder)
    router.GET("/orders", handler.GetAllOrders)
    router.GET("/orders/:id", handler.GetOrderById)
    router.PUT("/orders/:id", handler.UpdateOrder)
    router.DELETE("/orders/:id", handler.DeleteOrder)
} 