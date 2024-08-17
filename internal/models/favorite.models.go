package models

type Favorite struct {
	Id         string `db:"id" json:"id"`
	User_id    string `db:"user_id" json:"user_id"`
	Product_id string `db:"product_id" json:"product_id"`
}

type FavoriteGet struct {
	Display_name  string `db:"display_name" json:"display_name"`
	Product_name  string `db:"product_name" json:"product_name"`
	Product_price int    `db:"product_price" json:"product_price"`
	Img_product   string `db:"img_product" json:"img_product"`
}

type Favorites []Favorite
type FavoriteGets []FavoriteGet