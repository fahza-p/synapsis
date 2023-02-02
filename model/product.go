package model

type Product struct {
	Id         string  `json:"id"`
	CategoryId string  `json:"category_id"`
	Sku        string  `json:"sku"`
	Name       string  `json:"name"`
	Image      string  `json:"image"`
	Price      float64 `json:"price"`
	Stock      int32   `json:"stock"`
	CreatedAt  string  `json:"created_at,omitempty"`
	UpdatedAt  string  `json:"updated_at,omitempty"`
	CreatedBy  string  `json:"created_by,omitempty"`
	UpdatedBy  string  `json:"updated_by,omitempty"`
}
