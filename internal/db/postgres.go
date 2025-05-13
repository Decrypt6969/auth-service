package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/decrypt6969/auth-service/internal/config"
	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectPostgres() {
	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	user := config.GetEnv("DB_USER", "")
	pass := config.GetEnv("DB_PASSWORD", "")
	name := config.GetEnv("DB_NAME", "")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	DB, err = pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatal("failed to connect to postgres:", err)
	}

	log.Println("connected to postgres")
}
