package usecase

import (
	"example.com/go-auth-globo/internal/domain"
)

type UserUsecase struct {
	repo domain.IUserRepository
}

func (uc *UserUsecase) GetUser(email string) (*domain.User, error) {
	u, err := uc.repo.GetUser(email)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func NewUserUsecase(repo domain.IUserRepository) *UserUsecase {
	return &UserUsecase{
		repo,
	}
}
