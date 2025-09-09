package repository

import (
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"example.com/go-auth-globo/internal/domain"
	"example.com/go-auth-globo/internal/service"
	"example.com/go-auth-globo/internal/utils"
)

type AuthRepository struct {
	db *sql.DB
}

func (r *AuthRepository) DoLogin(input domain.LoginInput) (*domain.LoginOutput, error) {
	var hashed string
	var id uint

	rows, err := r.db.Query("SELECT id, senha FROM users WHERE email=($1)", input.Email)

	if err != nil {
		slog.Error("Failed to get user from database")
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&id, &hashed)

		if err != nil {
			slog.Error("Failed to scan user")
			return nil, err
		}

		err = service.ComparePassword([]byte(hashed), []byte(input.Password))

		if err != nil {
			slog.Error("Password is not matched")
			return nil, err
		}
	} else {
		slog.Error("User not found on database")
		err = errors.New("userNotFound")
		return nil, err
	}

	token, err := service.NewJWTService().CreateToken(id)

	if err != nil {
		return nil, err
	}

	output := domain.LoginOutput{
		Token:   token,
		Expires: utils.FormatDateToISO(time.Now().Add(time.Hour * 1)),
	}

	return &output, nil
}

func NewAuthRepository(db *sql.DB) domain.IAuthRepository {
	return &AuthRepository{db}
}
