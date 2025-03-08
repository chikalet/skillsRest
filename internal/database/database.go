package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

var DB *pgxpool.Pool

func Connect() (*sql.DB, string, error) {
	ctx := context.Background()

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, "", fmt.Errorf("ошибка подключения к БД (pgxpool.New): %v", err)
	}

	db := stdlib.OpenDBFromPool(pool)
	if err := db.PingContext(ctx); err != nil {
		pool.Close()
		return nil, "", fmt.Errorf("ошибка проверки соединения с БД (sql.OpenDB): %v", err)
	}

	DB = pool
	log.Println("Успешно подключились к базе данных (pgx)")
	return db, dsn, nil
}

func RunMigrations(dbURL string, migrationsPath string) error {
	migration, err := migrate.New(
		"file://"+migrationsPath,
		dbURL)
	if err != nil {
		return fmt.Errorf("ошибка создания клиента миграций: %w", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("ошибка применения миграций: %w", err)
	}

	fmt.Println("Миграции успешно применены")
	return nil
}
