package router

import (
	"ProductService/Repo"
	"ProductService/controller"
	"ProductService/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoadRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	rtr := router.Group("api/v1")

	customerRepo := Repo.NewCustomerRepo(db)
	customerService := service.NewCustomerService(customerRepo)
	customerController := controller.NewCustomController(customerService)

	productRepo := Repo.NewProductRepo(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	rtr.GET("/product", productController.GetProducts())        // get all products
	rtr.GET("/product/:id", productController.GetProductById()) //get product by id
	rtr.GET("/customer", customerController.GetCustomers())
	rtr.GET("/customer/:id", customerController.GetCustomerById())

	rtr.POST("/product", productController.CreateProduct())
	rtr.POST("/customer", customerController.CreateCustomer())

	return router
}
