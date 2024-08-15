package models

import (
	"time"
)

type Product struct {
	Id                  string     `db:"id" json:"id" form:"id"`
	Product_name        string     `db:"product_name" json:"product_name" form:"product_name"`
	Img_product         string     `db:"img_product" json:"img_product" form:"img_product"`
	Product_price       int        `db:"product_price" json:"product_price" form:"product_price"`
	Product_description string     `db:"product_description" json:"product_description" form:"product_description"`
	Categorie_name      string     `db:"categorie_name" json:"categorie_name" form:"categorie_name"`
	Category_id         int        `db:"category_id" json:"category_id" form:"category_id"`
	Created_at          *time.Time `db:"created_at" json:"created_at" form:"created_at"`
	Updated_at          *time.Time `db:"updated_at" json:"updated_at" form:"updated_at"`
}

type ProductDetail struct {
	Id                  string     `db:"id" json:"id" form:"id"`
	Product_name        string     `db:"product_name" json:"product_name" form:"product_name"`
	Img_product         string     `db:"img_product" json:"img_product" form:"img_product"`
	Product_price       int        `db:"product_price" json:"product_price" form:"product_price"`
	Product_description string     `db:"product_description" json:"product_description" form:"product_description"`
	Categorie_name      string     `db:"categorie_name" json:"categorie_name" form:"categorie_name"`
	Created_at          *time.Time `db:"created_at" json:"created_at" form:"created_at"`
	Updated_at          *time.Time `db:"updated_at" json:"updated_at" form:"updated_at"`
}

type EditProduct struct {
	Product_name        string `db:"product_name" json:"product_name" form:"product_name"`
	Img_product         string `db:"img_product" json:"img_product" form:"img_product"`
	Product_price       int    `db:"product_price" json:"product_price" form:"product_price"`
	Product_description string `db:"product_description" json:"product_description" form:"product_description"`
	Category_id         int    `db:"category_id" json:"category_id" form:"category_id"`
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