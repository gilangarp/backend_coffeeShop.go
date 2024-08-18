package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"backend_coffeeShop.go/internal/models"
	"github.com/jmoiron/sqlx"
)

type ProductRepositoryInterface interface {
    CreatedProduct(body *models.Product) (string, error)
    GetAllProduct(params *models.Filter) (*models.Products, error)
    GetDetailProduct(id string) (*models.ProductDetail, error)
    EditProduct(body *models.EditProduct, id string) (*models.EditProduct, error)
    DeleteProduct(id string) (string, error)
}

type ProductRepository struct {
	*sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) CreatedProduct(body *models.Product) (string, error) {
    query := `
    INSERT INTO public.product(
        product_name,
        image_url,
        price,
        description,
        category_id
        )
    VALUES
        ($1, $2, $3, $4, $5)
    `

    _, err := r.Exec(query, body.Product_name, body.Image_url, body.Price, body.Description, body.Category_id)
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("product %s has been created",body.Product_name), nil
}

func (r *ProductRepository) GetAllProduct(params *models.Filter) (*models.Products, error) {
    query := `
        SELECT p.id, p.product_name, p.image_url, p.price, p.description, c.categorie_name, p.created_at 
        FROM public.product p
        JOIN public.category c ON p.category_id = c.id
    `
    
    values := []interface{}{}
    whereClauses := []string{"p.is_deleted = FALSE"}

    if params.Promo {
        query += ` INNER JOIN public.promo prm ON p.id = prm.product_id `
    }

    if params.SearchText != "" {
        whereClauses = append(whereClauses, fmt.Sprintf("p.product_name ILIKE $%d", len(values)+1))
        values = append(values, fmt.Sprintf("%%%s%%", params.SearchText))
    }

    if params.Category != "" {
        whereClauses = append(whereClauses, fmt.Sprintf("c.categorie_name = $%d", len(values)+1))
        values = append(values, params.Category)
    }

    if len(whereClauses) > 0 {
        query += " WHERE " + strings.Join(whereClauses, " AND ")
    }

    if params.SortBy != "" {
        sortOrder := "DESC"
        switch params.SortBy {
        case "cheapest":
            sortOrder = "ASC"
        case "most_expensive":
            sortOrder = "DESC"
        default:
            return nil, fmt.Errorf("invalid sort parameter: %s", params.SortBy)
        }
        query += fmt.Sprintf(" ORDER BY p.price %s", sortOrder)
    }

    if params.Limit > 0 && params.Page > 0 {
        limit := params.Limit
        offset := (params.Page - 1) * limit
        query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(values)+1, len(values)+2)
        values = append(values, limit, offset)
    }

    var data models.Products
    if err := r.Select(&data, query, values...); err != nil {
        return nil, err
    }

    return &data, nil
}

func (r *ProductRepository) GetDetailProduct(id string) (*models.ProductDetail, error) {
	query := `SELECT p.id , p.product_name ,p.image_url ,p.price ,p.description ,c.categorie_name , p.created_at  FROM public.product p 
    join category c on p.category_id = c.id 
    WHERE p.id = $1`
	data := models.ProductDetail{}

	err := r.Get(&data, query , id)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

/* Edit Product */
func (r *ProductRepository) EditProduct(body *models.EditProduct, id string) (*models.EditProduct, error) {
    query := `UPDATE product SET `
    var values []interface{}
    condition := false

    if body.Product_name != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`product_name = $%d`, len(values)+1)
        values = append(values, body.Product_name)
        condition = true
    }

    if body.Image_url != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`image_url = $%d`, len(values)+1)
        values = append(values, body.Image_url)
        condition = true
    }

    if body.Price > 0 {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`price = $%d`, len(values)+1)
        values = append(values, body.Price)
        condition = true
    }

    if body.Description != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`description = $%d`, len(values)+1)
        values = append(values, body.Description)
        condition = true
    }

    if body.Category_id > 0 {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`category_id = $%d`, len(values)+1)
        values = append(values, body.Category_id)
        condition = true
    }

    if !condition {
        return nil, fmt.Errorf("no fields to update")
    }

    query += fmt.Sprintf(` WHERE id = $%d RETURNING product_name, image_url, price, description, category_id`, len(values)+1)
    values = append(values, id)

    row := r.DB.QueryRow(query, values...)
    var product models.EditProduct
    err := row.Scan(
        &product.Product_name,
        &product.Image_url,
        &product.Price,
        &product.Description,
        &product.Category_id,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("product with id = %s not found", id)
        }
        return nil, fmt.Errorf("query execution error: %w", err)
    }

    return &product, nil
}

/* Delete Product */
func (r *ProductRepository) DeleteProduct(id string) (string, error){
    query := `UPDATE product SET is_deleted = true WHERE id = $1`
    result, err := r.Exec(query, id)
    if err != nil {
        return "", fmt.Errorf("error while delete product: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return "", fmt.Errorf("error while fetching affected rows: %w", err)
    }

    if rowsAffected == 0 {
        return "", fmt.Errorf("product with ID %s not found", id)
    }

    return "Delete successful", nil
}
