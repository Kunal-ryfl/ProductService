package Repo

import (
	"ProductService/model"
	"gorm.io/gorm"
)

type ProductRepo interface {
	GetProduct() []model.Product
	GetProductById(id int) model.Product
	CreateProduct(product model.Product)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *productRepo {
	return &productRepo{db: db}
}

func (productRepo *productRepo) CreateProduct(product model.Product) {

	//var product model.Product
	if err := productRepo.db.Save(&product); err != nil {

	}

}
func (productRepo *productRepo) GetProduct() []model.Product {
	var products []model.Product
	if err := productRepo.db.Find(&products); err != nil {

	}

	return products
}
func (productRepo *productRepo) GetProductById(id int) model.Product {
	var product model.Product
	if err := productRepo.db.Find(&product, id); err != nil {

	}
	return product
}
