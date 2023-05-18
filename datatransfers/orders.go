package datatransfers

import "time"

type Orders struct {
	CustomerName string    `json:"customer_name" binding:"required"`
	OrderedAt    time.Time `json:"ordered_at" binding:"required"`
	Items        []Items   `json:"items" binding:"required"`
}

type Items struct {
	ItemCode    string `json:"item_code" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}

type OrdersUpdate struct {
	CustomerName string  `json:"customer_name" binding:"required"`
	Items        []Items `json:"items" binding:"required"`
}
