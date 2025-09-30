package authorize

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wolf1848/taxiportal/model"
	serviceErrors "github.com/wolf1848/taxiportal/service/errors"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) UserFindByEmail(email string) (*model.User, error) {

	var user model.User

	err := r.db.QueryRow(
		context.Background(),
		`SELECT id, name, email, password FROM users WHERE email=$1`,
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.HashPwd)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, serviceErrors.ErrRepositoryNoRows
		}
		return nil, err
	}

	return &user, nil
}

func (r *Repository) UserGetById(id int) (*model.User, error) {

	var user model.User

	err := r.db.QueryRow(
		context.Background(),
		`SELECT id, name, email, password FROM users WHERE id=$1`,
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.HashPwd)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, serviceErrors.ErrRepositoryNoRows
		}
		return nil, err
	}

	return &user, nil
}
