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

type BookRepo struct {
	db *pgxpool.Pool
}

func (s BookRepo) Create(ctx context.Context, req *models.CreateBook) (string, error) {
	var id = uuid.New().String()
	query := `INSERT INTO books(id, title, author, publisher, category, num_pages, lang) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := s.db.Exec(ctx, query, id, req.Title, req.Author, req.Publisher, req.Category, req.NumPages, req.Lang)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (s BookRepo) Update(ctx context.Context, req *models.UpdateBook) (int64, error) {
	var params map[string]interface{}
	query := `
		UPDATE books 
		SET title = :title,
		    author = :author, 
		    publisher = :publisher,
		    category = :category, 
		    num_pages = :num_pages,
			lang = :lang
		WHERE id = :id`

	params = map[string]interface{}{
		"id":        req.Id,
		"title":     req.Title,
		"author":    req.Author,
		"publisher": req.Publisher,
		"category":  req.Category,
		"num_pages": req.NumPages,
		"lang":      req.Lang,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (s BookRepo) GetById(ctx context.Context, req *models.BookPrimaryKey) (*models.Book, error) {
	var (
		id        sql.NullString
		title     sql.NullString
		author    sql.NullString
		publisher sql.NullString
		category  sql.NullString
		numPages  int
		lang      sql.NullString
	)

	query := `SELECT id, title, author, publisher, category, num_pages, lang FROM books WHERE id = $1 AND is_deleted = 'False'`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&title,
		&author,
		&publisher,
		&category,
		&numPages,
		&lang,
	)

	if err != nil {
		return nil, err
	}

	return &models.Book{
		Id:        id.String,
		Title:     title.String,
		Author:    author.String,
		Publisher: publisher.String,
		Category:  category.String,
		NumPages:  numPages,
		Lang:      lang.String,
	}, nil
}

func (s BookRepo) GetList(ctx context.Context, req *models.BookGetListRequest) (*models.BookGetListResponse, error) {
	var (
		resp   = &models.BookGetListResponse{}
		where  = " WHERE is_deleted = False "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		//order  = " ORDER BY created_at DESC "
	)
	query := `SELECT COUNT(*) OVER(), id, title, author, publisher, category, num_pages, lang FROM books`
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
			count     int
			id        sql.NullString
			title     sql.NullString
			author    sql.NullString
			publisher sql.NullString
			category  sql.NullString
			numPages  int
			lang      sql.NullString
		)
		err := rows.Scan(
			&count,
			&id,
			&title,
			&author,
			&publisher,
			&category,
			&numPages,
			&lang,
		)
		if err != nil {
			return nil, err
		}
		resp.Books = append(
			resp.Books,
			&models.Book{
				Id:        id.String,
				Title:     title.String,
				Author:    author.String,
				Publisher: publisher.String,
				Category:  category.String,
				NumPages:  numPages,
				Lang:      lang.String,
			})
		resp.Count = count
	}
	return resp, nil
}

func (s BookRepo) Delete(ctx context.Context, req *models.BookPrimaryKey) error {
	_, err := s.db.Exec(ctx, "UPDATE books SET is_deleted = true WHERE id = $1", req.Id)
	return err
}

func NewBookRepo(db *pgxpool.Pool) *BookRepo {
	return &BookRepo{
		db: db,
	}
}
