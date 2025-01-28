package storage

import (
	"app/api/models"
	"context"
)

type StorageInterface interface {
	Close()
	Users() UserRepoInterface
	Category() CategoryRepoInterface
	Books() BookRepoInterface
	Order() OrderRepoInterface
	OrderItem() OrderItemRepoInterface
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

type CategoryRepoInterface interface {
	Create(ctx context.Context, req *models.CreateCategory) (string, error)
	Update(ctx context.Context, req *models.UpdateCategory) (int64, error)
	GetById(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(ctx context.Context, req *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error)
	Delete(ctx context.Context, req *models.CategoryPrimaryKey) error
}

type OrderRepoInterface interface {
	Create(ctx context.Context, req *models.CreateOrder) (string, error)
	Update(ctx context.Context, req *models.UpdateOrder) (int64, error)
	GetById(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error)
	GetList(ctx context.Context, req *models.OrderGetListRequest) (*models.OrderGetListResponse, error)
	Delete(ctx context.Context, req *models.OrderPrimaryKey) error
}

type OrderItemRepoInterface interface {
	Create(ctx context.Context, req *models.CreateOrderItem) (string, error)
	Update(ctx context.Context, req *models.UpdateOrderItem) (int64, error)
	GetById(ctx context.Context, req *models.OrderItemPrimaryKey) (*models.OrderItem, error)
	GetList(ctx context.Context, req *models.OrderItemGetListRequest) (*models.OrderItemGetListResponse, error)
	Delete(ctx context.Context, req *models.OrderItemPrimaryKey) error
}
