package model

type Order struct {
	Id              string             `json:"id"`
	OrderNumber     string             `json:"order_number"`
	UserId          string             `json:"user_id"`
	TotalProducts   int32              `json:"total_product"`
	TotalItems      int32              `json:"total_items"`
	TotalPrice      float64            `json:"total_price"`
	TotalPaid       float32            `json:"total_paid"`
	Status          string             `json:"status"`
	StatusChangelog []*StatusChangelog `json:"status_changelog"`
	Items           []*OrderItems      `json:"items"`
	CreatedAt       string             `json:"created_at,omitempty"`
	UpdatedAt       string             `json:"updated_at,omitempty"`
	CreatedBy       string             `json:"created_by,omitempty"`
	UpdatedBy       string             `json:"updated_by,omitempty"`
}

type OrderItems struct {
	Id           string  `json:"id"`
	OrderId      string  `json:"order_id"`
	CategoryId   string  `json:"category_id"`
	CategoryName string  `json:"category_name"`
	ProductId    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductSku   string  `json:"product_sku"`
	ProductPrice float64 `json:"product_price"`
	Qty          int32   `json:"qty"`
	TotalPrice   float64 `json:"total_price"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
	CreatedBy    string  `json:"created_by,omitempty"`
	UpdatedBy    string  `json:"updated_by,omitempty"`
}

type StatusChangelog struct {
	OrderId   string `json:"order_id"`
	From      string `json:"from"`
	To        string `json:"to"`
	CreatedAt string `json:"created_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
}
