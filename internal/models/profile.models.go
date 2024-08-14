package models

type Profile struct {
	User_id          string `db:"user_id" json:"user_id" form:"user_id"`
	Display_name     string `db:"display_name" json:"display_name" form:"display_name" `
	First_name       string `db:"first_name" json:"first_name" form:"first_name"`
	Last_name        string `db:"last_name" json:"last_name" form:"last_name" `
	Image            string `db:"image" json:"image" `
	Delivery_address string `db:"delivery_address" json:"delivery_address" form:"delivery_address" `
	Birth_date       string `db:"birth_date" json:"birth_date" form:"birth_date" `
}

type Profiles []Profile