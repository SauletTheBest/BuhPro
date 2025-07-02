package routes

import (
	"BuhPro+/internal/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

// AuthRoutes настраивает маршруты для аутентификации обычного пользователя.
func AuthRoutes(router *gin.Engine, authHandler *handlers.AuthHandler, authMiddleware gin.HandlerFunc) {
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
	router.POST("/refresh", authHandler.Refresh)
	router.GET("/me", authMiddleware, authHandler.GetProfile)
}

// CustomerAuthRoutes настраивает маршруты для аутентификации клиента.
func CustomerAuthRoutes(router *gin.Engine, customerHandler *handlers.CustomerHandler, authMiddleware gin.HandlerFunc) {
	customerGroup := router.Group("/customer")
	{
		customerGroup.POST("/register", customerHandler.RegisterCustomer)
		customerGroup.POST("/login", customerHandler.LoginCustomer)
		customerGroup.POST("/refresh", customerHandler.RefreshCustomer)
		customerGroup.GET("/me", authMiddleware, customerHandler.GetCustomerProfile)
	}
}

// CoachAuthRoutes настраивает маршруты для аутентификации коуча.
func CoachAuthRoutes(router *gin.Engine, coachHandler *handlers.CoachHandler, authMiddleware gin.HandlerFunc) {
	coachGroup := router.Group("/coach")
	{
		coachGroup.POST("/register", coachHandler.RegisterCoach)
		coachGroup.POST("/login", coachHandler.LoginCoach)
		coachGroup.POST("/refresh", coachHandler.RefreshCoach)
		coachGroup.GET("/me", authMiddleware, coachHandler.GetCoachProfile)
	}
}

// ExecutorAuthRoutes настраивает маршруты для аутентификации исполнителя.
func ExecutorAuthRoutes(router *gin.Engine, executorHandler *handlers.ExecutorHandler, authMiddleware gin.HandlerFunc) {
	executorGroup := router.Group("/executor")
	{
		executorGroup.POST("/register", executorHandler.RegisterExecutor)
		executorGroup.POST("/login", executorHandler.LoginExecutor)
		executorGroup.POST("/refresh", executorHandler.RefreshExecutor)
		executorGroup.GET("/me", authMiddleware, executorHandler.GetExecutorProfile)
	}
}
