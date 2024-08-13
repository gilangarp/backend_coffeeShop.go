package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type claims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewJWT(id, email string) *claims {
	return &claims{
		Id:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "coffee shops",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		},
	}
}

func (c *claims) GenerateToken() (string, error) {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(secret))
}

func VerifyToken(token string) (*claims, error) {
	secret := os.Getenv("JWT_SECRET")
	data, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claimData := data.Claims.(*claims)
	return claimData, nil
}