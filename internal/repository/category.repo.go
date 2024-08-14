package repository

import (
	"fmt"

	"backend_coffeeShop.go/internal/models"
	"github.com/jmoiron/sqlx"
)

type CategoryRepositoryInterface interface {
	CreatedData(body *models.Category)(string,error)
	GetData()(*models.Categorys,error)
	UpdateData(body *models.Category , id string) (string , error)
	DeleteData(id string)(string , error)
}

type CategoryRepository struct {
	*sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func(r *CategoryRepository)CreatedData(body *models.Category)(string,error){
	query := `INSERT INTO category (categorie_name) VALUES ($1)`
	categorie_name := body.Categorie_name

	_,err := r.Exec(query , categorie_name)
	if err != nil {
		return "", err
	}

	return "data category created", nil
}

func(r *CategoryRepository)GetData()(*models.Categorys,error){
	query := `SELECT id , categorie_name  FROM category WHERE is_deleted = false`
	body := models.Categorys{}
	if err := r.Select(&body , query); err != nil {
		return nil, err
	}

	return &body, nil
}

func (r *CategoryRepository) UpdateData(body *models.Category , id string) (string , error) {
    query := `UPDATE category SET categorie_name = $1 WHERE id = $2`
	categorie_name := body.Categorie_name
    _, err := r.Exec(query, categorie_name, id)
    if err != nil {
        return "",err
    }
	successMessage := fmt.Sprintf("Category successfully updated to %s", body.Categorie_name)
    return successMessage ,nil
}

func (r *CategoryRepository)DeleteData(id string)(string , error){
	query := `UPDATE category SET is_deleted = true WHERE id = $1`

	_, err := r.Exec(query, id )
	if err != nil {
		return "", err
	}

	return "Delete successful" , nil
}
