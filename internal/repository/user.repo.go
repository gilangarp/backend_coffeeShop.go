package repository

import (
	"database/sql"
	"fmt"

	"backend_coffeeShop.go/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryInterface interface {
	CreateData(body *models.User)(string , error)
	GetAllData()(*models.UserDetails , error)
	GetDetailData(id string)(*models.UserDetail , error)
	UpdateData(data *models.User, id string) (*models.User, error)
	DeleteData(id string)(string , error)
}

type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository{
	return &UserRepository{db}
}

func (r *UserRepository) CreateData(body *models.User)(string , error){
	query := `INSERT INTO users (email, phone , password ) VALUES ( $1 , $2 , $3 )`

	_, err := r.Exec(query,body.Email , body.Phone , body.Password )
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

func (r *UserRepository) GetDetailData(id string)(*models.UserDetail , error){
	query := `select id , email , phone  FROM users WHERE id = $1 AND is_deleted = FALSE `
	data := models.UserDetail{}

	err := r.Get(&data, query , id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *UserRepository) UpdateData(data *models.User, id string) (*models.User, error) {
    query := `UPDATE users SET `
    var values []interface{}
    condition := false

    if data.Email != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`email = $%d`, len(values)+1)
        values = append(values, data.Email)
        condition = true
    }

    if data.Phone != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`phone = $%d`, len(values)+1)
        values = append(values, data.Phone)
        condition = true
    }

    if data.Password != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`password = $%d`, len(values)+1)
        values = append(values, data.Password)
        condition = true
    }

    if !condition {
        return nil, fmt.Errorf("no fields to update")
    }

    query += fmt.Sprintf(` WHERE id = $%d RETURNING email, phone, password`, len(values)+1)
    values = append(values, id)

    row := r.DB.QueryRow(query, values...)
    var user models.User
    err := row.Scan(
        &user.Email,
        &user.Phone,
        &user.Password,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user with id = %s not found", id)
        }
        return nil, fmt.Errorf("query execution error: %w", err)
    }

    return &user, nil
}

func (r *UserRepository) DeleteData(id string)(string , error) {
	query := `UPDATE users SET is_deleted = true WHERE id = $1`
	_, err := r.Exec(query,id )
	if err != nil {
		return "", err
	}

	return "Delete successful" , nil
}