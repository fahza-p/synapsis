package model

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

var OrderFilter = []string{"order_number", "status"}

/* Request */
type OrderReq struct {
	ProductId []int64 `json:"product_id" validate:"gt=0,dive,required,number"`
}

/* Model Data */
type OrderData struct {
	Id            int64   `json:"id"`
	OrderNumber   string  `json:"order_number"`
	UserId        int64   `json:"user_id"`
	TotalProducts int32   `json:"total_product"`
	TotalItems    int32   `json:"total_items"`
	TotalPrice    float64 `json:"total_price"`
	TotalPaid     float64 `json:"total_paid"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at,omitempty"`
	UpdatedAt     string  `json:"updated_at,omitempty"`
	CreatedBy     string  `json:"created_by,omitempty"`
	UpdatedBy     string  `json:"updated_by,omitempty"`
}

/* Response */
type Order struct {
	Id              int64              `json:"id"`
	OrderNumber     string             `json:"order_number"`
	UserId          int64              `json:"user_id"`
	TotalProducts   int32              `json:"total_product"`
	TotalItems      int32              `json:"total_items"`
	TotalPrice      float64            `json:"total_price"`
	TotalPaid       float64            `json:"total_paid"`
	Status          string             `json:"status"`
	StatusChangelog []*StatusChangelog `json:"status_changelog"`
	Items           []*OrderItems      `json:"items"`
	CreatedAt       string             `json:"created_at,omitempty"`
	UpdatedAt       string             `json:"updated_at,omitempty"`
	CreatedBy       string             `json:"created_by,omitempty"`
	UpdatedBy       string             `json:"updated_by,omitempty"`
}

type OrderItems struct {
	Id           int64   `json:"id"`
	OrderId      int64   `json:"order_id"`
	ProductId    int64   `json:"product_id"`
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
	OrderId   int64  `json:"order_id"`
	From      string `json:"from"`
	To        string `json:"to"`
	CreatedAt string `json:"created_at,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
}

/* Functions */
func (u *OrderData) SetOrderCreateData(totalProduct, totalItem int32, totalPrice, totalPaid float64, userId int64, userEmail string, itemData []*CartItems) ([]*OrderItems, *StatusChangelog) {
	resItemsData := []*OrderItems{}
	now := time.Now()
	datetimeFormat := now.Format("2006-01-02 15:04:05")

	u.OrderNumber = fmt.Sprintf("ORD-%d-%v", userId, now.UnixMilli())
	u.UserId = userId
	u.TotalProducts = totalProduct
	u.TotalItems = totalItem
	u.TotalPrice = totalPrice
	u.TotalPaid = totalPaid
	u.Status = "Menunggu Pembayaran"
	u.CreatedAt = datetimeFormat
	u.UpdatedAt = datetimeFormat
	u.CreatedBy = userEmail
	u.UpdatedBy = userEmail

	mapstructure.Decode(itemData, &resItemsData)
	for i := range resItemsData {
		resItemsData[i].Id = 0
		resItemsData[i].ProductName = itemData[i].Name
		resItemsData[i].ProductSku = itemData[i].Sku
		resItemsData[i].ProductPrice = itemData[i].Price
	}

	return resItemsData, &StatusChangelog{
		From:      "",
		To:        "Menunggu Pembayaran",
		CreatedAt: datetimeFormat,
		CreatedBy: userEmail,
	}
}
