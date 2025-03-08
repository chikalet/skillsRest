package handlers

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"log"
	"skillsRest/internal/database"
	"skillsRest/internal/models"
	"time"
)

func GetTask(c fiber.Ctx) error {
	ctx := context.Background()

	err := database.DB.Ping(ctx)
	if err != nil {
		log.Printf("Ошибка пинга базы данных: %v", err)
		return c.Status(500).SendString("Ошибка подключения к базе данных")
	}
	rows, err := database.DB.Query(ctx, "SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		log.Printf("Ошибка выполнения запроса: %v", err)
		return c.Status(500).SendString("Ошибка запроса к базе данных")
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Printf("Ошибка сканирования данных: %v", err)
			return c.Status(500).SendString("Ошибка обработки данных")
		}
		tasks = append(tasks, task)
	}
	return c.JSON(tasks)
}

func CreateTask(c fiber.Ctx) error {
	ctx := context.Background()
	task := new(models.Task)
	if err := c.Bind().Body(task); err != nil {
		log.Printf("Неверный формат запроса: %v", err)
		return c.Status(400).SendString("Неверный формат запроса")
	}

	_, err := database.DB.Exec(ctx, "INSERT INTO tasks (title, description) VALUES ($1, $2)",
		task.Title, task.Description)
	if err != nil {
		log.Printf("Ошибка вставки данных в базу: %v", err)
		return c.Status(500).SendString("Ошибка вставки данных в базу")
	}

	return c.Status(201).SendString("Задача успешно создана")
}

func UpdateTask(c fiber.Ctx) error {
	ctx := context.Background()
	currentTime := time.Now()
	id := c.Params("id")
	task := new(models.Task)

	if err := c.Bind().Body(task); err != nil {
		log.Printf("Неверный формат запроса: %v", err)
		return c.Status(400).SendString("Неверный формат запроса")
	}

	_, err := database.DB.Exec(ctx, "UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5",
		task.Title, task.Description, task.Status, currentTime, id)
	if err != nil {
		log.Printf("Ошибка при выполнении UPDATE запроса: %v", err)
		return c.Status(500).SendString("Ошибка обновления данных")
	}

	return c.SendString("Задача успешно обновлена")
}

func DeleteTask(c fiber.Ctx) error {
	ctx := context.Background()
	id := c.Params("id")
	_, err := database.DB.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		log.Println("Ошибка удаления задачи:", id)
		return c.Status(500).SendString("Ошибка удаления задачи")
	}

	return c.SendString("Задача успешно удалена")
}
