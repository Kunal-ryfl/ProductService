package service

import (
	"OrderService/Repo"
	//"github.com/gin-gonic/gin"
	//"go/types"
	"net/http"
	"strconv"

	//"context"
	//"encoding/json"
	//"fmt"
	"OrderService/model"
)

type OrderService interface {
	//GetAllOrders() []model.Order
	//GetById(id int) model.Order
	CreateOrder(customer model.Order) Response
}

type Response struct {
	Code int
	Msg  string
}

type orderService struct {
	orderRepo Repo.OrderRepo
}

func NewOrderService(Repo Repo.OrderRepo) OrderService {
	return &orderService{orderRepo: Repo}
}

//func (orderService *orderService) GetAllOrders() []model.Order {
//	return orderService.orderRepo.GetOrders()
//}
//
//func (orderService *orderService) GetById(id int) model.Order {
//	return orderService.orderRepo.GetOrderById(id)
//}

func (orderService *orderService) CreateOrder(order model.Order) Response {
	// add validation for productid amd customerId

	custBool := doesCustomerExist(order.CustomerId)
	proBool := doesProductExist(order.ProductId)

	if !custBool {
		return Response{ // error -1
			Code: -1,
			Msg:  "customer not exist",
		}
	}

	if !proBool {
		return Response{ // error -1
			Code: -1,
			Msg:  "product not exist",
		}
	}

	orderService.orderRepo.CreateOrder(order)
	return Response{
		Code: 1,
		Msg:  "all good",
	}
}

func doesProductExist(id int) bool {
	idString := strconv.Itoa(id)
	res, _ := http.Get("http://localhost:8080/api/v1/product/" + idString)

	if res.StatusCode == 400 {
		return false
	}

	return true
}

func doesCustomerExist(id int) bool {
	idString := strconv.Itoa(id)
	res, _ := http.Get("http://localhost:8080/api/v1/customer/" + idString)

	if res.StatusCode == 400 {
		return false
	}

	return true
}
