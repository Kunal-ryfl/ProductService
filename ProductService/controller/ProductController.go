package controller

import (
	"ProductService/model"
	"ProductService/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

type createProductRqst struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

//type createUserResp struct {
//	Msg string `json:"msg"`
//}

func (ProductController *ProductController) CreateProduct() gin.HandlerFunc {
	return func(context *gin.Context) {

		var createProductRqst createProductRqst
		if err := context.ShouldBindJSON(&createProductRqst); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error in body": err.Error()})
			return
		}
		//if err != nil {
		//	fmt.Println("Error reading ")
		//}

		newProduct := model.Product{
			ID:       createProductRqst.Id,
			Name:     createProductRqst.Name,
			Price:    createProductRqst.Price,
			Category: createProductRqst.Category,
		}

		ProductController.productService.CreateProduct(newProduct)

		context.JSON(200, gin.H{
			"msg": "Product Created",
		})
	}
}

func (ProductController *ProductController) GetProducts() gin.HandlerFunc {

	return func(context *gin.Context) {
		products := ProductController.productService.GetAllProducts()
		context.JSON(200, products)
	}

}
func (ProductController *ProductController) GetProductById() gin.HandlerFunc {

	return func(context *gin.Context) {

		query := context.Param("id")

		id, _ := strconv.Atoi(query)
		product := ProductController.productService.GetById(id)

		if product.ID == 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": "product not found",
			})
			return
		}

		context.JSON(200, product)
	}

}
