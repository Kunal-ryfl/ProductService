package model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID      int `gorm:"primaryKey"`
	Name    string
	Phone   string
	Address string
}
