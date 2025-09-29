package drivers

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wolf1848/taxiportal/model"
)

func NewPostgres(c *model.Postgres) *pgxpool.Pool {

	var pool *pgxpool.Pool

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Pwd, c.Db, c.Ssl,
	)

	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		panic(err.Error())
	}

	// Настройки пула соединений
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = time.Minute * 30

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		panic(err.Error())
	}

	// Проверка соединения
	if err := pool.Ping(ctx); err != nil {
		panic(err.Error())
	}

	return pool
}

func ShutdownPostgres(ctx context.Context, db *pgxpool.Pool) error {
	if db == nil {
		return nil
	}

	done := make(chan error, 1)
	go func() {
		db.Close()
		done <- nil
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}
