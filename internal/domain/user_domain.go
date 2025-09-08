package domain

import "database/sql"

type User struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// TODO: Criar output de GetUser

type UserDTO struct {
	Name     sql.NullString
	Email    sql.NullString
	Password sql.NullString
}

type IUserRepository interface {
	GetUser(email string) (*User, error)
}
