package postgres

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type OrderItemRepo struct {
	db *pgxpool.Pool
}

func (s OrderItemRepo) Create(ctx context.Context, req *models.CreateOrderItem) (string, error) {
	var id = uuid.New().String()
	query := `INSERT INTO order_items(item_id, order_id, book_id) VALUES ($1, $2, $3)`

	_, err := s.db.Exec(ctx, query, id, req.OrderId, req.BookId)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (s OrderItemRepo) Update(ctx context.Context, req *models.UpdateOrderItem) (int64, error) {
	var params map[string]interface{}
	query := `
		UPDATE order_items 
		SET order_id = :order_id,
		    book_id = :book_id,
		    updated_at = now()
		WHERE item_id = :item_id`

	params = map[string]interface{}{
		"book_id":  req.BookId,
		"order_id": req.OrderId,
		"item_id":  req.ItemId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (s OrderItemRepo) GetById(ctx context.Context, req *models.OrderItemPrimaryKey) (*models.OrderItem, error) {
	var (
		itemId  sql.NullString
		orderId sql.NullString
		bookId  sql.NullString
	)

	query := `SELECT item_id, order_id, book_id FROM order_items WHERE item_id = $1 AND is_deleted = 'False'`

	err := s.db.QueryRow(ctx, query, req.ItemId).Scan(
		&itemId,
		&orderId,
		&bookId,
	)

	if err != nil {
		return nil, err
	}

	return &models.OrderItem{
		ItemId:  itemId.String,
		OrderId: orderId.String,
		BookId:  bookId.String,
	}, nil
}

func (s OrderItemRepo) GetList(ctx context.Context, req *models.OrderItemGetListRequest) (*models.OrderItemGetListResponse, error) {
	var (
		resp   = &models.OrderItemGetListResponse{}
		where  = " WHERE is_deleted = False "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		//order  = " ORDER BY created_at DESC "
	)
	query := `SELECT COUNT(*) OVER(), item_id, order_id, book_id FROM order_items`
	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	//if req.Search != "" {
	//	where += ` AND title ILIKE '%' || '` + req.Search + `' || '%'`
	//}

	//query += where + order + offset + limit
	query += where + offset + limit

	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			itemId  sql.NullString
			orderId sql.NullString
			bookId  sql.NullString
			count   int
		)
		err := rows.Scan(
			&count,
			&itemId,
			&orderId,
			&bookId,
		)
		if err != nil {
			return nil, err
		}
		resp.OrderItems = append(
			resp.OrderItems,
			&models.OrderItem{
				ItemId:  itemId.String,
				OrderId: orderId.String,
				BookId:  bookId.String,
			})
		resp.Count = count
	}
	return resp, nil
}

func (s OrderItemRepo) Delete(ctx context.Context, req *models.OrderItemPrimaryKey) error {
	_, err := s.db.Exec(ctx, "UPDATE order_items SET is_deleted = true, updated_at = now() WHERE item_id = $1", req.ItemId)
	return err
}

func NewOrderItemRepo(db *pgxpool.Pool) *OrderItemRepo {
	return &OrderItemRepo{
		db: db,
	}
}
