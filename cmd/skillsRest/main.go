package main

import (
	"log"
	"os"
	"os/signal"
	"skillsRest/internal/database"
	"skillsRest/internal/transport"
	"syscall"

	"github.com/gofiber/fiber/v3"
)

func main() {

	if err := database.Connect(); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	app := fiber.New()

	// Подключаем middleware
	// app.Use(logger.New())   // Логирование запросов
	// app.Use(compress.New()) // Сжатие ответов
	// app.Use(recover.New())  // Восстановление после паники
	// app.Use(limiter.New())  // Лимит запросов для предотвращения DDOS атак

	transport.RegisterTaskRoutes(app)

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Завершение сервера...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Ошибка при завершении сервера: %v", err)
	}
	log.Println("Сервер завершен.")
}
