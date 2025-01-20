package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID         int `gorm:"primaryKey"`
	CustomerId int
	Amount     int
	ProductId  int
}
