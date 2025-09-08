package usecase

import (
	"example.com/go-auth-globo/internal/domain"
)

type AuthUsecase struct {
	repo domain.IAuthRepository
}

func (l *AuthUsecase) DoLogin(input domain.LoginInput) (*domain.LoginOutput, error) {
	output, err := l.repo.DoLogin(input)

	if err != nil {
		return nil, err
	}

	return output, nil
}

func NewAuthUsecase(repo domain.IAuthRepository) *AuthUsecase {
	return &AuthUsecase{
		repo,
	}
}
