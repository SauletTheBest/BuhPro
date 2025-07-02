// cmd/web/main.go
package main

import (
    "log"
    "net/http"

    "buhpro/internal/delivery/gin/routes"
    "buhpro/internal/delivery/http/handler"
    "buhpro/internal/repository"
    "buhpro/internal/models"
    "buhpro/internal/services"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    // Подключение к БД
    dsn := "host=localhost user=postgres password=0000 dbname=buhpro port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Автомиграция
    err = db.AutoMigrate(&models.Order{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    // Инициализация репозиториев
    orderRepo := repository.NewOrderRepository(db)

    // Инициализация Service
    orderService := services.NewOrderService(orderRepo)

    // Инициализация Handler'ов
    orderHandler := handler.NewOrderHandler(orderService)

    // Настройка маршрутов
    r := gin.Default()
    api := r.Group("/api/v1")

    routes.OrderRoutes(api, orderHandler)

    // Запуск сервера
    log.Println("Server is running on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}