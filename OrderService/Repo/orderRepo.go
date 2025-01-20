package Repo

import (
	"OrderService/model"
	"gorm.io/gorm"
)

type OrderRepo interface {
	//GetOrders() []model.Order
	//GetOrderById(id int) model.Order
	CreateOrder(order model.Order)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *orderRepo {
	return &orderRepo{db: db}
}

func (orderRepo *orderRepo) CreateOrder(order model.Order) {

	//var product model.Product
	if err := orderRepo.db.Save(&order); err != nil {

	}
}

//func (orderRepo *orderRepo) GetOrders() []model.Order {
//	var orders []model.Order
//	if err := orderRepo.db.Find(&orders); err != nil {
//
//	}
//
//	return orders
//}
//func (orderRepo *orderRepo) GetOrderById(id int) model.Order {
//	var order model.Order
//	if err := orderRepo.db.Find(&order, id); err != nil {
//
//	}
//	return order
//}
