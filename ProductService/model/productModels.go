package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	Name     string
	Price    int
	Category string
}
