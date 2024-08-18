package models

import (
	"time"
)

type Product struct {
	Id                  string     `db:"id" json:"id" form:"id"`
	Product_name        string     `db:"product_name" json:"product_name" form:"product_name"`
	Image_url         	string     `db:"image_url" json:"image_url" form:"image_url"`
	Price       		int        `db:"price" json:"price" form:"price"`
	Description  		string     `db:"description" json:"description" form:"description"`
	Categorie_name      string     `db:"categorie_name" json:"categorie_name" form:"categorie_name"`
	Category_id         int        `db:"category_id" json:"category_id" form:"category_id"`
	Created_at          *time.Time `db:"created_at" json:"created_at" form:"created_at"`
	Updated_at          *time.Time `db:"updated_at" json:"updated_at" form:"updated_at"`
}

type ProductDetail struct {
	Id                  string     `db:"id" json:"id" form:"id"`
	Product_name        string     `db:"product_name" json:"product_name" form:"product_name"`
	Image_url         	string     `db:"image_url" json:"image_url" form:"image_url"`
	Price       		int        `db:"price" json:"price" form:"price"`
	Description  		string     `db:"description" json:"description" form:"description"`
	Categorie_name      string     `db:"categorie_name" json:"categorie_name" form:"categorie_name"`
	Created_at          *time.Time `db:"created_at" json:"created_at" form:"created_at"`
	Updated_at          *time.Time `db:"updated_at" json:"updated_at" form:"updated_at"`
}

type EditProduct struct {
	Product_name        string 	   `db:"product_name" json:"product_name" form:"product_name"`
	Image_url         	string     `db:"image_url" json:"image_url" form:"image_url"`
	Price       		int        `db:"price" json:"price" form:"price"`
	Description  		string     `db:"description" json:"description" form:"description"`
	Category_id         int        `db:"category_id" json:"category_id" form:"category_id"`
}


type Filter struct {
    Category    string `form:"category"`
    Favorite    string `form:"favorite"`
    SearchText  string `form:"SearchText"`
    Limit       int    `form:"limit"`
    Page        int    `form:"page"`
    Promo       bool
    SortBy      string `form:"sortBy"`  
}

type Products []Product