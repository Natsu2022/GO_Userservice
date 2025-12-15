package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"OSCIRA.com/service/user_service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres(cfg *config.Config) (*pgxpool.Pool, error) {

	dsn := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBDriver, cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	conf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	conf.MaxConns = 10
	conf.MinConns = 2
	conf.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL")
	return pool, nil
}
