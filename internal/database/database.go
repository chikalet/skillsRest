// package database

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// var DB *pgxpool.Pool

// func Connect() error {
// 	ctx := context.Background()

// 	dbHost := "localhost"
// 	dbPort := "5432"
// 	dbUser := "postgres"
// 	dbPassword := "root"
// 	dbName := "test"

// 	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

// 	config, err := pgxpool.ParseConfig(connStr)
// 	if err != nil {
// 		return fmt.Errorf("ошибка парсинга DSN: %v", err)
// 	}

// 	pool, err := pgxpool.NewWithConfig(ctx, config)
// 	if err != nil {
// 		return fmt.Errorf("ошибка подключения к БД: %v", err)
// 	}

// 	if err := pool.Ping(ctx); err != nil {
// 		return fmt.Errorf("ошибка ping БД: %v", err)
// 	}

//		DB = pool
//		log.Println("Успешно подключились к базе данных")
//		return nil
//	}
package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {
	ctx := context.Background()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return fmt.Errorf("ошибка подключения к БД: %v", err)
	}

	DB = pool // Сохраняем пул соединений
	log.Println("Успешно подключились к базе данных")
	return nil
}
