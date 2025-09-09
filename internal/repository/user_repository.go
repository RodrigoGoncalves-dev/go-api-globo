package repository

import (
	"database/sql"
	"errors"
	"log/slog"

	"example.com/go-auth-globo/internal/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) domain.IUserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUser(email string) (*domain.User, error) {
	var dto domain.UserDTO
	var u domain.User
	rows, err := r.DB.Query("SELECT nome, email FROM users WHERE email=($1)", email)

	if err != nil {
		slog.Error("Failed to get user from database")
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {

		err = rows.Scan(&dto.Name, &dto.Email)

		if err != nil {
			slog.Error("Failed to scan user")
			return nil, err
		}

		u = domain.User{
			Name:  dto.Name.String,
			Email: dto.Email.String,
		}
	} else {
		slog.Error("User not found on database")
		err = errors.New("userNotFound")
		return nil, err
	}
	slog.Info("Successfully found user on database")

	return &u, nil
}
