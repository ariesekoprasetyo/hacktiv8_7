package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Addresses []Address `gorm:"foreignKey:UserID" json:"addresses"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Address struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"-"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
