package repositories

import (
	"context"
	db "eattogether/pkg/db"

	"eattogether/internal/models"

	"github.com/jackc/pgx/v5"
	pgconn "github.com/jackc/pgx/v5/pgconn"
)

type PgxInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Row, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

type UserRepo struct {
	DB PgxInterface
}

func (p *UserRepo) GetUser() (models.User, error) {
	var user models.User
	res, err := p.DB.Query(context.Background(), "SELECT * FROM places")
	if err != nil {
		return user, err
	}

	res.Scan(&user)
	return user, nil
}

func CreateUserRepo(db *db.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}
