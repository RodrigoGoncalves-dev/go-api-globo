package service

import (
	"fmt"
	"time"

	"example.com/go-auth-globo/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secretKey []byte
	issure    string
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.RegisteredClaims
}

func (j *jwtService) CreateToken(id uint) (string, error) {
	claim := &Claim{
		Sum: id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: j.issure,
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 1),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString(j.secretKey)

	if err != nil {
		Logger().Error("Unsuccessfully to signed token")
		return "", err
	}

	return t, nil
}

func (j *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(j.secretKey), nil
	})

	return err == nil
}

func NewJWTService() *jwtService {
	secretKey := config.AppInfo.SECRET_KEY
	issuer := config.AppInfo.ISSUER

	return &jwtService{
		secretKey: []byte(secretKey),
		issure:    issuer,
	}
}
