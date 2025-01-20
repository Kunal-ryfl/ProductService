package service

import (
	"ProductService/Repo"
	//"context"
	//"encoding/json"
	//"fmt"
	"ProductService/model"
)

type ProductService interface {
	GetAllProducts() []model.Product
	GetById(id int) model.Product
	CreateProduct(customer model.Product)
}

type productService struct {
	productRepo Repo.ProductRepo
}

func NewProductService(Repo Repo.ProductRepo) ProductService {
	return &productService{productRepo: Repo}
}

func (productService *productService) GetAllProducts() []model.Product {
	return productService.productRepo.GetProduct()
}

func (productService *productService) GetById(id int) model.Product {
	return productService.productRepo.GetProductById(id)
}

func (productService *productService) CreateProduct(product model.Product) {
	productService.productRepo.CreateProduct(product)
}
