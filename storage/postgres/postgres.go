package postgres

import (
	"app/config"
	"app/storage"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type store struct {
	db        *pgxpool.Pool
	user      *UserRepo
	category  *CategoryRepo
	book      *BookRepo
	order     *OrderRepo
	orderItem *OrderItemRepo
}

func (s *store) Users() storage.UserRepoInterface {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}
	return s.user
}

func (s *store) Category() storage.CategoryRepoInterface {
	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}
	return s.category
}

func (s *store) Books() storage.BookRepoInterface {
	if s.book == nil {
		s.book = NewBookRepo(s.db)
	}
	return s.book
}

func (s *store) Order() storage.OrderRepoInterface {
	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}
	return s.order
}

func (s *store) OrderItem() storage.OrderItemRepoInterface {
	if s.orderItem == nil {
		s.orderItem = NewOrderItemRepo(s.db)
	}
	return s.orderItem
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageInterface, error) {

	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = cfg.PostgresMaxConnection

	pool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pool,
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}
