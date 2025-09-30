package register

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wolf1848/taxiportal/model"
	"github.com/wolf1848/taxiportal/service/register/entity"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertUser(user *model.User) error {
	err := r.db.QueryRow(
		context.Background(),
		`INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`,
		user.Name, user.Email, user.HashPwd,
	).Scan(&user.ID)

	if err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return entity.ErrUniqueEmail
		}

		return err
	}

	return nil
}
