package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wolf1848/taxiportal/repository/authorize"
	"github.com/wolf1848/taxiportal/repository/register"
)

type Repositories struct {
	register  *register.Repository
	authorize *authorize.Repository
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		register:  register.NewRepository(db),
		authorize: authorize.NewRepository(db),
	}
}

func (r *Repositories) Register() *register.Repository {
	return r.register
}

func (r *Repositories) Authorize() *authorize.Repository {
	return r.authorize
}
