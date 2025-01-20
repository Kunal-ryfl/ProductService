package controller

import (
	"OrderService/model"
	"OrderService/service"
	"OrderService/worker"
	"fmt"
	"strconv"

	//service2 "ProductService/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	jsonOrders []model.Order
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

type createOrderRqst struct {
	Id         int `json:"id"`
	CustomerId int `json:"customerId"`
	Amount     int `json:"amount"`
	ProductId  int `json:"productId"`
}

//type createUserResp struct {
//	Msg string `json:"msg"`
//}

func (OrderController *OrderController) CreateOrder() gin.HandlerFunc {
	return func(context *gin.Context) {

		var createOrderRqst createOrderRqst
		if err := context.ShouldBindJSON(&createOrderRqst); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error in body": err.Error()})
			return
		}
		//if err != nil {
		//	fmt.Println("Error reading ")
		//}

		newOrder := model.Order{
			ID:         createOrderRqst.Id,
			CustomerId: createOrderRqst.CustomerId,
			Amount:     createOrderRqst.Amount,
			ProductId:  createOrderRqst.ProductId,
		}

		var response service.Response
		response = OrderController.orderService.CreateOrder(newOrder)

		if response.Code == -1 {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": response.Msg,
			})
			return
		}

		context.JSON(200, gin.H{
			"msg": "order Created and validated both",
		})
	}
}
func (OrderController *OrderController) CreateBulkOrder() gin.HandlerFunc {
	return func(context *gin.Context) {

		//jsonOrders := []model.Order{
		//
		//	{
		//		ID:         1,
		//		ProductId:  1,
		//		CustomerId: 100,
		//		Amount:     100,
		//	},
		//	{
		//		ID:         1,
		//		ProductId:  1,
		//		CustomerId: 1,
		//		Amount:     100,
		//	},
		//}

		fmt.Println("order", jsonOrders)

		taskChannel := make(chan model.Order)
		resultChannel := make(chan service.Response)

		createBulkOrderRqst := []createOrderRqst{}

		if err := context.ShouldBindJSON(&createBulkOrderRqst); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error in body": err.Error()})
			return
		}

		fmt.Printf("%+v", createBulkOrderRqst)

		ctr := 0

		for _, order := range createBulkOrderRqst {

			//if err := context.ShouldBindJSON(&createBulkOrderRqst); err != nil {
			//	context.JSON(http.StatusBadRequest, gin.H{"error in body": err.Error()})
			//	return
			//}

			//fmt.Println(u_order)

			newOrder := model.Order{
				ID:         order.Id,
				CustomerId: order.CustomerId,
				Amount:     order.Amount,
				ProductId:  order.ProductId,
			}

			//fmt.Println(newOrder)
			go worker.ProcessOrderRow(OrderController.orderService, taskChannel, resultChannel)

			taskChannel <- newOrder

		}

		for i := 0; i < len(createBulkOrderRqst); i++ {
			//fmt.Println("looping")
			select {
			case res := <-resultChannel:
				{

					if res.Code != -1 {
						ctr++
						fmt.Println(res)
					}

				}
				//default:
				//	break

			}

		}

		//newOrder := model.Order{
		//	ID:         createOrderRqst.Id,
		//	CustomerId: createOrderRqst.CustomerId,
		//	Amount:     createOrderRqst.Amount,
		//	ProductId:  createOrderRqst.ProductId,
		//}

		//var response service.Response
		////response = OrderController.orderService.CreateOrder(newOrder)
		//
		//if response.Code == -1 {
		//	context.JSON(http.StatusBadRequest, gin.H{
		//		"msg": response.Msg,
		//	})
		//	return
		//}

		ctr_string := strconv.Itoa(ctr)
		context.JSON(200, gin.H{
			"msg": ctr_string + " orders Created and validated both",
		})
	}
}

//
//func (ProductController *ProductController) GetProducts() gin.HandlerFunc {
//
//	return func(context *gin.Context) {
//		products := ProductController.productService.GetAllProducts()
//		context.JSON(200, products)
//	}
//
//}
//func (ProductController *ProductController) GetProductById() gin.HandlerFunc {
//
//	return func(context *gin.Context) {
//
//		query := context.Param("id")
//
//		id, _ := strconv.Atoi(query)
//		product := ProductController.productService.GetById(id)
//		context.JSON(200, product)
//	}
//
//}
