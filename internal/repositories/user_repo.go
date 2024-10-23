package repositories

import (
	"context"
	db "eattogether/pkg/db"

	"eattogether/internal/models"
)

type UserRepo struct {
	DB db.PgxInterface
}

func (p *UserRepo) GetUser() (models.User, error) {
	var user models.User
	res := p.DB.QueryRow(context.Background(), "SELECT * FROM places")

	res.Scan(&user)
	return user, nil
}

func CreateUserRepo(db *db.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}
