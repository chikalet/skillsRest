package database

import (
	//	"database/sql"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	//"os"
)

//var DB *sql.DB

const dsn = "postgres://postgres:root@localhost:5432/pgxdb"

func Connect() error {
	//urlExample := "postgres://username:postgres password@localhost:5432/database_name"
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return fmt.Errorf("ошибка подключения к БД: %v", err)
	}
	defer conn.Close(ctx)
	log.Println("Успешно подключились к базе данных")
	return nil
}
