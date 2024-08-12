package repository

import (
	"fmt"

	"backend_coffeeShop.go/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepositoryInterface interface {
	GetByEmail(email string)(*models.UserLogin , error)
}

type AuthRepository struct {
	*sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func(r *AuthRepository) GetByEmail(email string)(*models.UserLogin , error){
	result := models.UserLogin{}
	query := `SELECT id , email , phone , password from users WHERE email = $1`
	err := r.Get(&result , query , email)
	if err != nil {
		return nil , err
	}
	fmt.Printf("dari res %s ",&result )

	return &result , nil
}