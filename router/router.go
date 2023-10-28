package router

import (
	"crud/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupRouter configures the Echo router and defines routes
func SetupRouter(e *echo.Echo) {
	// Middleware (optional)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// API routes for tasks
	taskGroup := e.Group("/tasks")

	// CRUD routes for tasks
	taskGroup.GET("", handlers.GetTasks)
	taskGroup.GET("/:id", handlers.GetTask)
	taskGroup.POST("", handlers.CreateTask)
	taskGroup.PUT("/:id", handlers.UpdateTask)
	taskGroup.DELETE("/:id", handlers.DeleteTask)
}
