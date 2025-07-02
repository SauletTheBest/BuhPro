package main

import (
	"BuhPro+/internal/config"
	"BuhPro+/internal/delivery/gin/middleware"
	"BuhPro+/internal/delivery/gin/routes"
	"BuhPro+/internal/delivery/http/handlers"
	"BuhPro+/internal/repository"
	"BuhPro+/internal/usecase"
	"BuhPro+/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Загрузка конфигурации
	cfg := config.LoadConfig()

	// 2. Настройка логгеров
	appLogger := utils.SetupLogger(cfg.AppLogFile)
	serviceLogger := utils.SetupLogger(cfg.ServiceLogFile)
	handlerLogger := utils.SetupLogger(cfg.HandlerLogFile)

	// 3. Подключение к базе данных
	database := config.Connect(cfg.DBURL)

	// 4. Инициализация репозиториев
	userRepo := repository.NewUserRepository(database)
	customerRepo := repository.NewCustomerRepository(database)
	coachRepo := repository.NewCoachRepository(database)
	executorRepo := repository.NewExecutorRepository(database)

	// Пустые репозитории для будущих функций
	// orderRepo := repository.NewOrderRepository(database)
	// responseRepo := repository.NewResponseRepository(database)
	// ratingRepo := repository.NewRatingRepository(database)
	// courseRepo := repository.NewCourseRepository(database)
	// paymentRepo := repository.NewPaymentRepository(database)

	// 5. Инициализация UseCase (бизнес-логика)
	authUsecase := usecase.NewAuthUsecase(userRepo, cfg.JWTSecret, serviceLogger)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo, cfg.JWTSecret, serviceLogger)
	coachUsecase := usecase.NewCoachUsecase(coachRepo, cfg.JWTSecret, serviceLogger)
	executorUsecase := usecase.NewExecutorUsecase(executorRepo, cfg.JWTSecret, serviceLogger)

	// Пустые UseCase для будущих функций
	// orderUsecase := usecase.NewOrderUsecase(orderRepo, serviceLogger)
	// responseUsecase := usecase.NewResponseUsecase(responseRepo, serviceLogger)
	// ratingUsecase := usecase.NewRatingUsecase(ratingRepo, serviceLogger)
	// courseUsecase := usecase.NewCourseUsecase(courseRepo, serviceLogger)
	// paymentUsecase := usecase.NewPaymentUsecase(paymentRepo, serviceLogger)

	// 6. Инициализация HTTP-обработчиков
	authHandler := handlers.NewAuthHandler(authUsecase, handlerLogger)
	customerHandler := handlers.NewCustomerHandler(customerUsecase, handlerLogger)
	coachHandler := handlers.NewCoachHandler(coachUsecase, handlerLogger)
	executorHandler := handlers.NewExecutorHandler(executorUsecase, handlerLogger)

	// Пустые обработчики для будущих функций
	// orderHandler := handlers.NewOrderHandler(/* dependencies */)
	// responseHandler := handlers.NewResponseHandler(/* dependencies */)
	// ratingHandler := handlers.NewRatingHandler(/* dependencies */)
	// courseHandler := handlers.NewCourseHandler(/* dependencies */)
	// paymentHandler := handlers.NewPaymentHandler(/* dependencies */)

	// 7. Инициализация Gin роутера
	r := gin.Default()

	// 8. Инициализация общего JWT Middleware
	authMiddleware := middleware.JWTAuth(cfg.JWTSecret)

	// 9. Настройка маршрутов
	routes.AuthRoutes(r, authHandler, authMiddleware)
	routes.CustomerAuthRoutes(r, customerHandler, authMiddleware)
	routes.CoachAuthRoutes(r, coachHandler, authMiddleware)
	routes.ExecutorAuthRoutes(r, executorHandler, authMiddleware)

	// Пустые маршруты для будущих функций
	// routes.OrderRoutes(r, orderHandler, authMiddleware)
	// routes.ResponseRoutes(r, responseHandler, authMiddleware)
	// routes.RatingRoutes(r, ratingHandler, authMiddleware)
	// routes.CourseRoutes(r, courseHandler, authMiddleware)
	// routes.PaymentRoutes(r, paymentHandler, authMiddleware)

	// 10. Запуск сервера
	if err := r.Run(":" + cfg.Port); err != nil {
		appLogger.Fatalf("Failed to start server: %v", err)
	}
}
