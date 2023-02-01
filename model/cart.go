package model

type Cart struct {
	Id         string  `json:"id"`
	UserId     string  `json:"user_id"`
	TotalItems int32   `json:"total_items"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  int64   `json:"created_at,omitempty"`
	UpdatedAt  int64   `json:"updated_at,omitempty"`
	CreatedBy  string  `json:"created_by,omitempty"`
	UpdatedBy  string  `json:"updated_by,omitempty"`
}

type CartItems struct {
	Id        string `json:"id"`
	CartId    string `json:"cart_id"`
	ProductId string `json:"product_id"`
	Qty       int32  `json:"qty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
}
