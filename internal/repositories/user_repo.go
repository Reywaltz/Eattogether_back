package repositories

import (
	"context"
	db "eattogether/pkg/db"

	"eattogether/internal/models"
)

type UserRepo struct {
	DB db.PgxInterface
}

func (p *UserRepo) GetUser(username string, password string) (models.User, error) {
	var user models.User
	res := p.DB.QueryRow(
		context.Background(),
		"SELECT id, username, password, role FROM users WHERE username=$1 and password=$2",
		username,
		password,
	)

	err := res.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUserRepo(db *db.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}
