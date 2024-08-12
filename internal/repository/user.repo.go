package repository

import (
	"fmt"

	"backend_coffeeShop.go/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryInterface interface {
	CreateData(body *models.User)(string , error)
	GetAllData()(*models.UserDetails , error)
}

type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository{
	return &UserRepository{db}
}


func (r *UserRepository) CreateData(body *models.User)(string , error){
	query := `INSERT INTO users (email, phone, password) VALUES ( $1 , $2 , $3)`

	_, err := r.Exec(query,body.Email , body.Phone , body.Password)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("congratulations account %s has been registered",body.Email) , nil
}

func (r *UserRepository) GetAllData()(*models.UserDetails , error){
	query := `select id , email , phone  FROM users `
	data := models.UserDetails{}

	err := r.Select(&data, query)
	if err != nil {
		return nil, err
	}
	return &data, nil
}