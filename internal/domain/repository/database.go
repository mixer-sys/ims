package repository

import (
	"context"
	"fmt"
	"ims/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func (db *Database) Close() {
	db.Pool.Close()
}

func NewDatabase(ctx context.Context, cfg *config.Config) (*Database, error) {

	pool, err := pgxpool.Connect(ctx, fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User,
		cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return &Database{Pool: pool}, nil
}
