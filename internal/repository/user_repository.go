package repository

import (
	"database/sql"

	"example.com/go-auth-globo/internal/domain"
	"example.com/go-auth-globo/internal/service"
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
		service.Logger().Error("Failed to get user from database")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&dto.Name, &dto.Email)

		if err != nil {
			service.Logger().Error("Failed to scan user")
			return nil, err
		}

		u = domain.User{
			Name:  dto.Name.String,
			Email: dto.Email.String,
		}
	}

	return &u, nil
}
