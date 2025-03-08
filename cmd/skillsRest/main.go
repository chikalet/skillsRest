package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"skillsRest/internal/database"
	"skillsRest/internal/transport"
	"syscall"

	"github.com/gofiber/fiber/v3"
)

func main() {

	db, dbURL, err := database.Connect()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	migrationsPath := "./migrations"
	if err := database.RunMigrations(dbURL, migrationsPath); err != nil {
		log.Fatalf("Ошибка при выполнении миграций: %v", fmt.Errorf("ошибка миграций: %w", err))
	}

	app := fiber.New()

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
