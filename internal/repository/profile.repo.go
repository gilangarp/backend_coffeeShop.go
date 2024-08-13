package repository

import (
	"database/sql"
	"fmt"

	"backend_coffeeShop.go/internal/models"
	"github.com/jmoiron/sqlx"
)

type ProfileRepositoryInterface interface {
	CreatedData(body *models.Profile, id string) (string, error)
	GetAllData() (*models.Profiles , error)
	GetDetailData(id string) (*models.Profile, error)
	EditData(data *models.Profile , id string) (*models.Profile, error)
	DeleteData(id string) (string, error)
}

type ProfileRepository struct {
	*sqlx.DB
}

func NewProfileRepository(db *sqlx.DB) *ProfileRepository {
	return &ProfileRepository{db}
}

func (r *ProfileRepository) CreatedData(body *models.Profile, id string) (string, error) {
	query := `
	INSERT INTO profile (
    	user_id,
    	display_name,
    	first_name,
    	last_name,
    	birth_date,
    	image,
    	delivery_address
	) VALUES
	    ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.Exec(query, id, body.Display_name, body.First_name, body.Last_name, body.Birth_date, body.Image, body.Delivery_address)
	if err != nil {
		return "", err
	}

	return "Congratulations on creating your profile.", nil
}

func (r *ProfileRepository) GetAllData() (*models.Profiles , error){
	query := `SELECT * FROM public.profile `
	data := models.Profiles{}

	if err := r.Select(&data , query); err != nil {
		return nil , err
	}

	return &data , nil
}

func (r *ProfileRepository) GetDetailData(id string) (*models.Profile, error) {
	query := `SELECT * FROM public.profile WHERE user_id = $1`
	data := models.Profile{}

	err := r.Get(&data, query , id)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *ProfileRepository) EditData(data *models.Profile , id string) (*models.Profile, error) {
    query := `UPDATE profile SET `
    var values []interface{}
    condition := false

    if data.Display_name != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`display_name = $%d`, len(values)+1)
        values = append(values, data.Display_name)
        condition = true
    }

    if data.First_name != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`first_name = $%d`, len(values)+1)
        values = append(values, data.First_name)
        condition = true
    }

    if data.Last_name != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`last_name = $%d`, len(values)+1)
        values = append(values, data.Last_name)
        condition = true
    }

    if data.Birth_date != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`birth_date = $%d`, len(values)+1)
        values = append(values, data.Birth_date)
        condition = true
    }

    if data.Image != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`image = $%d`, len(values)+1)
        values = append(values, data.Image)
        condition = true
    }

    if data.Delivery_address != "" {
        if condition {
            query += ", "
        }
        query += fmt.Sprintf(`delivery_address = $%d`, len(values)+1)
        values = append(values, data.Delivery_address)
        condition = true
    }

    if !condition {
        return nil, fmt.Errorf("no fields to update")
    }

    query += fmt.Sprintf(` WHERE user_id = $%d RETURNING *`, len(values)+1)
    values = append(values, id)
	fmt.Print("ini dari value:", values)
	fmt.Print("ini dari query:", query)

    row := r.DB.QueryRow(query, values...)
    var profile models.Profile
    err := row.Scan(
        &profile.User_id,
        &profile.Display_name,
        &profile.First_name,
        &profile.Last_name,
        &profile.Birth_date,
        &profile.Image,
        &profile.Delivery_address,
    )
	
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user with id = %s not found", id)
        }
        return nil, fmt.Errorf("query execution error: %w", err)
    }

    return &profile, nil
}

func (r *ProfileRepository) DeleteData(id string) (string, error) {
    query := `DELETE FROM public.profile WHERE user_id = $1 RETURNING user_id`
	_, err := r.Exec(query, id)
	if err != nil {
		return "", err
	}
	return "data deleted", nil
}