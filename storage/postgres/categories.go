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

type CategoryRepo struct {
	db *pgxpool.Pool
}

func (s CategoryRepo) Create(ctx context.Context, req *models.CreateCategory) (string, error) {
	var id = uuid.New().String()
	query := `INSERT INTO categories(id, name, type, picture) VALUES ($1, $2, $3, $4)`

	_, err := s.db.Exec(ctx, query, id, req.Name, req.Type, req.Picture)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (s CategoryRepo) Update(ctx context.Context, req *models.UpdateCategory) (int64, error) {
	var params map[string]interface{}
	query := `
		UPDATE categories 
		SET name = :name, 
		    type = :type,
		    picture = :picture,
		    updated_at = now()
		WHERE id = :id`

	params = map[string]interface{}{
		"id":      req.Id,
		"name":    req.Name,
		"type":    req.Type,
		"picture": req.Picture,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (s CategoryRepo) GetById(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error) {
	var (
		id      sql.NullString
		name    sql.NullString
		Type    sql.NullString
		picture sql.NullString
	)

	query := `SELECT id, name, type, picture FROM categories WHERE id = $1 AND is_deleted = FALSE`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&Type,
		&picture,
	)

	if err != nil {
		return nil, err
	}

	return &models.Category{
		Id:      id.String,
		Name:    name.String,
		Type:    Type.String,
		Picture: picture.String,
	}, nil
}

func (s CategoryRepo) GetList(ctx context.Context, req *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error) {
	var (
		resp   = &models.CategoryGetListResponse{}
		where  = " WHERE is_deleted = FALSE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		//order  = " ORDER BY created_at DESC "
	)
	query := `SELECT COUNT(*) OVER(), id, name, type, picture FROM categories`
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
			count   int
			id      sql.NullString
			name    sql.NullString
			Type    sql.NullString
			picture sql.NullString
		)
		err := rows.Scan(
			&count,
			&id,
			&name,
			&Type,
			&picture,
		)
		if err != nil {
			return nil, err
		}
		resp.Categories = append(
			resp.Categories,
			&models.Category{
				Id:      id.String,
				Name:    name.String,
				Type:    Type.String,
				Picture: picture.String,
			})
		resp.Count = count
	}

	return resp, nil
}

func (s CategoryRepo) Delete(ctx context.Context, req *models.CategoryPrimaryKey) error {
	_, err := s.db.Exec(ctx, "UPDATE categories SET is_deleted = true, updated_at = now() WHERE id = $1", req.Id)
	return err
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}
