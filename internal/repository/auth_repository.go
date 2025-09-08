package repository

import (
	"database/sql"
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
		service.Logger().Error("Failed to get user from database")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &hashed)

		if err != nil {
			service.Logger().Error("Failed to scan user")
			return nil, err
		}

		err = service.ComparePassword([]byte(hashed), []byte(input.Password))

		if err != nil {
			service.Logger().Error("Password is not matched")
			return nil, err
		}
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
