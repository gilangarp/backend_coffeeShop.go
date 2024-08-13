package models

type Profile struct {
	User_id          string `db:"user_id" json:"user_id"`
	Display_name     string `db:"display_name" json:"display_name"`
	First_name       string `db:"first_name" json:"first_name"`
	Last_name        string `db:"last_name" json:"last_name"`
	Birth_date       string `db:"birth_date" json:"birth_date"`
	Image            string `db:"image" json:"image"`
	Delivery_address string `db:"delivery_address" json:"delivery_address"`
}

type Profiles []Profile