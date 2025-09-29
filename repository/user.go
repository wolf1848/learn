package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/wolf1848/taxiportal/model"
	"github.com/wolf1848/taxiportal/service/entity"
)

func (repo *Repository) UserInsert(input *model.User) (*int, error) {

	var id *int

	err := repo.db.QueryRow(
		context.Background(),
		`INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`,
		input.Name, input.Email, input.HashPwd,
	).Scan(&id)

	if err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, model.ErrUniqueEmail
		}

		return nil, err

	}

	return id, nil

}

func (repo *Repository) UserFindByEmail(email string) (*model.User, error) {

	var user model.User

	err := repo.db.QueryRow(
		context.Background(),
		`SELECT id, name, email, password FROM users WHERE email=$1`,
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.HashPwd)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrNoRows
		}
		return nil, err
	}

	return &user, nil

}

func (repo *Repository) UserGetById(id int) (*model.User, error) {

	var user model.User

	err := repo.db.QueryRow(
		context.Background(),
		`SELECT id, name, email, password FROM users WHERE id=$1`,
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.HashPwd)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrNoRows
		}
		return nil, err
	}

	return &user, nil

}
