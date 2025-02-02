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

type OrderRepo struct {
	db *pgxpool.Pool
}

func (s OrderRepo) Create(ctx context.Context, req *models.CreateOrder) (string, error) {
	var id = uuid.New().String()
	query := `INSERT INTO orders(order_id, user_id) VALUES ($1, $2)`

	_, err := s.db.Exec(ctx, query, id, req.UserId)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (s OrderRepo) Update(ctx context.Context, req *models.UpdateOrder) (int64, error) {
	var params map[string]interface{}
	query := `
		UPDATE orders 
		SET user_id = :user_id
		WHERE order_id = :order_id`

	params = map[string]interface{}{
		"user_id":  req.UserId,
		"order_id": req.OrderId,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (s OrderRepo) GetById(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error) {
	var (
		userId  sql.NullString
		orderId sql.NullString
	)

	query := `SELECT order_id, user_id FROM orders WHERE order_id = $1 AND is_deleted = 'False'`

	err := s.db.QueryRow(ctx, query, req.OrderId).Scan(
		&orderId,
		&userId,
	)

	if err != nil {
		return nil, err
	}

	return &models.Order{
		OrderId: orderId.String,
		UserId:  userId.String,
	}, nil
}

func (s OrderRepo) GetList(ctx context.Context, req *models.OrderGetListRequest) (*models.OrderGetListResponse, error) {
	var (
		resp   = &models.OrderGetListResponse{}
		where  = " WHERE is_deleted = False "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		//order  = " ORDER BY created_at DESC "
	)
	query := `SELECT COUNT(*) OVER(), order_id, user_id FROM orders`
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
			userId  sql.NullString
			orderId sql.NullString
			count   int
		)
		err := rows.Scan(
			&count,
			&orderId,
			&userId,
		)
		if err != nil {
			return nil, err
		}
		resp.Orders = append(
			resp.Orders,
			&models.Order{
				OrderId: orderId.String,
				UserId:  userId.String,
			})
		resp.Count = count
	}
	return resp, nil
}

func (s OrderRepo) Delete(ctx context.Context, req *models.OrderPrimaryKey) error {
	_, err := s.db.Exec(ctx, "UPDATE orders SET is_deleted = true WHERE order_id = $1", req.OrderId)
	return err
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}
