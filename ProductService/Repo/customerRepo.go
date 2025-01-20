package Repo

import (
	"ProductService/model"
	"gorm.io/gorm"
)

type CustomerRepo interface {
	GetCustomer() []model.Customer
	GetCustomerById(id int) model.Customer
	CreateCustomer(customer model.Customer)
}

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *customerRepo {
	return &customerRepo{db: db}
}

func (customerRepo *customerRepo) CreateCustomer(customer model.Customer) {

	//var product model.Product
	if err := customerRepo.db.Save(&customer); err != nil {

	}

}
func (customerRepo *customerRepo) GetCustomer() []model.Customer {
	var customers []model.Customer

	if err := customerRepo.db.Find(&customers); err != nil {

	}

	return customers
}
func (customerRepo *customerRepo) GetCustomerById(id int) model.Customer {
	var customer model.Customer
	if err := customerRepo.db.Find(&customer, id); err != nil {
	}
	return customer
}
