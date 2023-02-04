package model

import (
	"html"
	"time"
)

/* Request */
type ProductCreateReq struct {
	CategoryId int64   `json:"category_id" validate:"required,number"`
	Sku        string  `json:"sku" validate:"required,alphanum,max=20"`
	Name       string  `json:"name" validate:"required,ascii,max=200"`
	Image      string  `json:"image" validate:"omitempty,url"`
	Price      float64 `json:"price" validate:"required,number,min=0"`
	Stock      int32   `json:"stock" validate:"required,number,min=0"`
}

type Product struct {
	Id         int64   `json:"id"`
	CategoryId int64   `json:"category_id"`
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

/* Functions */
func (u *Product) SetProductCreateData(req *ProductCreateReq, userEmail string) {
	now := time.Now().Format("2006-01-02 15:04:05")

	u.CategoryId = req.CategoryId
	u.Name = html.EscapeString(req.Name)
	u.Sku = html.EscapeString(req.Sku)
	u.Image = html.EscapeString(u.Image)
	u.Price = req.Price
	u.Stock = req.Stock
	u.CreatedAt = now
	u.UpdatedAt = now
	u.CreatedBy = userEmail
	u.UpdatedBy = userEmail
}
