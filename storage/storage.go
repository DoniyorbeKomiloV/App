package storage

import (
	"app/api/models"
	"context"
)

type StorageInterface interface {
	Close()
	Books() BookRepoInterface
	Users() UserRepoInterface
	Order() OrderRepoInterface
}

type BookRepoInterface interface {
	Create(ctx context.Context, req *models.CreateBook) (string, error)
	Update(ctx context.Context, req *models.UpdateBook) (int64, error)
	GetById(ctx context.Context, req *models.BookPrimaryKey) (*models.Book, error)
	GetList(ctx context.Context, req *models.BookGetListRequest) (*models.BookGetListResponse, error)
	Delete(ctx context.Context, req *models.BookPrimaryKey) error
}

type UserRepoInterface interface {
	Create(ctx context.Context, req *models.CreateUser) (string, error)
	Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	GetById(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetList(ctx context.Context, req *models.UserGetListRequest) (*models.UserGetListResponse, error)
	Delete(ctx context.Context, req *models.UserPrimaryKey) error
}

type OrderRepoInterface interface {
	Create(ctx context.Context, req *models.CreateOrder) (string, error)
	Update(ctx context.Context, req *models.UpdateOrder) (int64, error)
	GetById(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error)
	GetList(ctx context.Context, req *models.OrderGetListRequest) (*models.OrderGetListResponse, error)
	Delete(ctx context.Context, req *models.OrderPrimaryKey) error
}
