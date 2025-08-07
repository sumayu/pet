package bd

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sumayu/pet/internal/logger"
	"go.uber.org/zap"
)
func BD() (*pgxpool.Pool) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,	
	)
	if user == "" || password == "" || dbname == "" || host == "" || port == "" {
		 logger.Error("dsn = nil ,bd")
		 return nil
	}

bd, err := pgxpool.New(context.Background(),dsn)
  if err != nil {
        logger.Error("Failed to create connection pool", zap.Error(err))
        return nil
    }

if err := bd.Ping(context.Background()); err != nil {
	bd.Close()
		logger.Error("Failed to ping database", zap.Error(err))
		return nil
	}
logger.Info("Successfully connected to PostgreSQL")
return bd
}

