package models

type Category struct {
	ID             string `db:"id" json:"id" valid:"-"`
	Categorie_name string `db:"categorie_name" json:"categorie_name" valid:"-"`
}

type Categorys []Category