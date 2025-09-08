package repository

import (
	"context"
	"fmt"
	"ims/config"
	"ims/internal/domain/handlers"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

func (db *Database) Close() {
	db.pool.Close()
}

func NewDatabase(ctx context.Context, cfg *config.Config) (*Database, error) {

	pool, err := pgxpool.Connect(ctx, fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User,
		cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &Database{pool: pool}, nil
}

type SQLCategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *Database) handlers.CategoryRepository {
	return &SQLCategoryRepository{db: db.pool}
}

type SQLProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *Database) handlers.ProductRepository {
	return &SQLProductRepository{db: db.pool}
}
