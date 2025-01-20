package service

import (
	//"context"
	//"encoding/json"
	//"fmt"
	"ProductService/Repo"
	"ProductService/model"
)

type CustomerService interface {
	GetAllCustomers() []model.Customer
	GetById(id int) model.Customer
	CreateCustomer(customer model.Customer)
}

type customerService struct {
	customerRepo Repo.CustomerRepo
}

func NewCustomerService(Repo Repo.CustomerRepo) CustomerService {
	return &customerService{customerRepo: Repo}
}

func (customerService *customerService) GetAllCustomers() []model.Customer {
	return customerService.customerRepo.GetCustomer()
}

func (customerService *customerService) GetById(id int) model.Customer {
	return customerService.customerRepo.GetCustomerById(id)
}

func (customerService *customerService) CreateCustomer(customer model.Customer) {
	customerService.customerRepo.CreateCustomer(customer)
}
