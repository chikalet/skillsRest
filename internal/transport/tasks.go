package transport

import (
	"skillsRest/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func RegisterTaskRoutes(app *fiber.App) {
	api := app.Group("/")

	api.Get("/tasks", handlers.GetTask)           // Получить все продукты
	api.Post("/tasks", handlers.CreateTask)       // Создать новый продукт
	api.Put("/tasks/:id", handlers.UpdateTask)    // Обновить продукт
	api.Delete("/tasks/:id", handlers.DeleteTask) // Удалить продукт
}
