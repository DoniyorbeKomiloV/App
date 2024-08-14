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

type UserRepo struct {
	db *pgxpool.Pool
}

func (s UserRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {
	var id = uuid.New().String()
	query := `INSERT INTO users(id, first_name, last_name, age, phone, password, card_no) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := s.db.Exec(ctx, query, id, req.FirstName, req.LastName, req.Age, req.Phone, req.Password, req.CardNo)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (s UserRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {
	var params map[string]interface{}
	query := `
		UPDATE users 
		SET first_name = :first_name,
		    last_name = :last_name, 
		    age = :age,
		    phone = :phone, 
		    password = :password,
		    card_no = :card_no
		WHERE id = :id`

	params = map[string]interface{}{
		"id":         req.Id,
		"first_name": req.FirstName,
		"last_name":  req.LastName,
		"age":        req.Age,
		"phone":      req.Phone,
		"password":   req.Password,
		"card_no":    req.CardNo,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (s UserRepo) GetById(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {
	var (
		id        sql.NullString
		firstName sql.NullString
		lastName  sql.NullString
		age       int
		phone     sql.NullString
		password  sql.NullString
		cardNo    sql.NullString
	)

	query := `SELECT id, first_name, last_name, age, phone, password, card_no FROM users WHERE id = $1 AND is_deleted = FALSE`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&firstName,
		&lastName,
		&age,
		&phone,
		&password,
		&cardNo,
	)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:        id.String,
		FirstName: firstName.String,
		LastName:  lastName.String,
		Age:       age,
		Phone:     phone.String,
		Password:  password.String,
		CardNo:    cardNo.String,
	}, nil
}

func (s UserRepo) GetList(ctx context.Context, req *models.UserGetListRequest) (*models.UserGetListResponse, error) {
	var (
		resp   = &models.UserGetListResponse{}
		where  = " WHERE is_deleted = FALSE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
		//order  = " ORDER BY created_at DESC "
	)
	query := `SELECT COUNT(*) OVER(), id, first_name, last_name, age, phone, password, card_no FROM users`
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
			firstName sql.NullString
			lastName  sql.NullString
			age       int
			phone     sql.NullString
			password  sql.NullString
			cardNo    sql.NullString
		)
		err := rows.Scan(
			&count,
			&id,
			&firstName,
			&lastName,
			&age,
			&phone,
			&password,
			&cardNo,
		)
		if err != nil {
			return nil, err
		}
		resp.Users = append(
			resp.Users,
			&models.User{
				Id:        id.String,
				FirstName: firstName.String,
				LastName:  lastName.String,
				Age:       age,
				Phone:     phone.String,
				Password:  password.String,
				CardNo:    cardNo.String,
			})
		resp.Count = count
	}

	return resp, nil
}

func (s UserRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) error {
	_, err := s.db.Exec(ctx, "UPDATE users SET is_deleted = true WHERE id = $1", req.Id)
	return err
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}
