package models

type Promo struct {
	Discount   string
	Value      float64
	Product_id string
}

type GetPromo struct {
	Discount     string
	Value        float64
	Product_name string
}

type Promos []GetPromo