package controller

import (
	"ProductService/model"
	"ProductService/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CustomerController struct {
	CustomerService service.CustomerService
}

func NewCustomController(customerService service.CustomerService) *CustomerController {
	return &CustomerController{CustomerService: customerService}
}

type createCustomerRqst struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type createUserResp struct {
	Msg string `json:"msg"`
}

func (CustomerController *CustomerController) CreateCustomer() gin.HandlerFunc {
	return func(context *gin.Context) {

		var createCustomerRqst createCustomerRqst
		if err := context.ShouldBindJSON(&createCustomerRqst); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error in body": err.Error()})
			return
		}
		//if err != nil {
		//	fmt.Println("Error reading ")
		//}

		newCustomer := model.Customer{
			ID:      createCustomerRqst.Id,
			Name:    createCustomerRqst.Name,
			Phone:   createCustomerRqst.Phone,
			Address: createCustomerRqst.Address,
		}

		CustomerController.CustomerService.CreateCustomer(newCustomer)

		context.JSON(200, gin.H{
			"msg": "Customer Created",
		})
	}
}

func (CustomerController *CustomerController) GetCustomers() gin.HandlerFunc {

	return func(context *gin.Context) {
		customers := CustomerController.CustomerService.GetAllCustomers()
		context.JSON(200, customers)
	}

}
func (CustomerController *CustomerController) GetCustomerById() gin.HandlerFunc {

	return func(context *gin.Context) {

		query := context.Param("id")

		id, _ := strconv.Atoi(query)
		customer := CustomerController.CustomerService.GetById(id)

		if customer.ID == 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": "customer not found",
			})
			return
		}

		context.JSON(200, customer)
	}

}
