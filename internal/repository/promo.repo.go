package repository

import (
	"fmt"
	"strings"

	"backend_coffeeShop.go/internal/models"
	"github.com/jmoiron/sqlx"
)

type PromoRepositoryInterface interface{
	CreateData(body *models.Promo)(string , error)
	GetData()(*models.Promos , error)
	UpdateData(data *models.Promo, id string) (string, error)
	DeleteData(id string)(string,error)
}

type PromoRepository struct {
	*sqlx.DB
}

func NewPromoRepository(db *sqlx.DB) *PromoRepository {
	return &PromoRepository{db}
}

func (r *PromoRepository)CreateData(body *models.Promo)(string , error){
	query := `INSERT INTO public.promo (discount,value,product_id) VALUES ($1, $2, $3)`
	_, err := r.Exec(query, body.Discount ,body.Value ,body.Product_id)
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("promo %s has been created",body.Discount), nil
}

func (r *PromoRepository)GetData()(*models.Promos , error){
	query := `SELECT discount,value,p2.product_name  FROM public.promo p inner join product p2 on p.product_id = p2.id `
	body := models.Promos{}
	if err := r.Select(&body , query); err != nil {
        return nil, err
    }

    return &body, nil
}

func (r *PromoRepository) UpdateData(data *models.Promo, id string) (string, error) {
    var values []interface{}
    var updates []string
    condition := false

    if data.Discount != "" {
        updates = append(updates, fmt.Sprintf(`discount = $%d`, len(values)+1))
        values = append(values, data.Discount)
        condition = true
    }

    if data.Value > 0 {
        updates = append(updates, fmt.Sprintf(`value = $%d`, len(values)+1))
        values = append(values, data.Value)
        condition = true
    }

    if !condition {
        return "", fmt.Errorf("no fields to update")
    }

    query := `UPDATE promo SET ` + strings.Join(updates, ", ") + fmt.Sprintf(` WHERE id = $%d`, len(values)+1)
    values = append(values, id)

    result, err := r.DB.Exec(query, values...)
    if err != nil {
        return "", fmt.Errorf("query execution error: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return "", fmt.Errorf("failed to get rows affected: %w", err)
    }

    if rowsAffected == 0 {
        return "", fmt.Errorf("promo with id = %s not found", id)
    }

    return "Promo successfully updated", nil
}


func (r *PromoRepository) DeleteData(id string)(string,error){
	query := `UPDATE promo SET is_deleted = true WHERE id = $1`

	_, err := r.Exec(query, id )
	if err != nil {
		return "", err
	}

	return "Delete successful" , nil
}

