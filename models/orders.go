package models

import "time"

type Orders struct {
	OrderId      uint      `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `json:"customer_name"`
	Ordered_at   time.Time `gorm:"autoUpdateTime" json:"ordered_at"`
	Created_at   time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Items struct {
	ItemId      uint      `gorm:"primaryKey" json:"item_id"`
	ItemCode    string    `json:"item_code"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Order_id    []Orders  `gorm:"foreignKey:OrderId" json:"order_id"`
	Created_at  time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
