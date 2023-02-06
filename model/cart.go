package model

import (
	"time"
)

var CartItemFilter = []string{"sku", "name"}

/* Request */
type CartAddItemReq struct {
	ProductId int64 `json:"product_id" validate:"required,number"`
	Qty       int32 `json:"qty" validate:"required,number"`
}

/* Model Data */
type CartData struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
}

type CartItemData struct {
	Id        int64  `json:"id"`
	CartId    int64  `json:"cart_id"`
	ProductId int64  `json:"product_id"`
	Qty       int32  `json:"qty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
}

/*Response*/
type Cart struct {
	Id           int64   `json:"id"`
	UserId       int64   `json:"user_id"`
	TotalProduct int32   `json:"total_product"`
	TotalItems   int32   `json:"total_items"`
	TotalPrice   float64 `json:"total_price"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
	CreatedBy    string  `json:"created_by,omitempty"`
	UpdatedBy    string  `json:"updated_by,omitempty"`
}

type CartItems struct {
	Id         int64   `json:"id"`
	CartId     int64   `json:"cart_id"`
	ProductId  int64   `json:"product_id"`
	Sku        string  `json:"sku"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Qty        int32   `json:"qty"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"created_at,omitempty"`
	UpdatedAt  string  `json:"updated_at,omitempty"`
	CreatedBy  string  `json:"created_by,omitempty"`
	UpdatedBy  string  `json:"updated_by,omitempty"`
}

/* Functions */
func (u *CartItemData) SetCartItemAddProductData(cartId int64, userEmail string) {
	now := time.Now().Format("2006-01-02 15:04:05")

	u.CartId = cartId
	u.CreatedAt = now
	u.UpdatedAt = now
	u.CreatedBy = userEmail
	u.UpdatedBy = userEmail
}
