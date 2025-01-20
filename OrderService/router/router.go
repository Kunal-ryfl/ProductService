package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	//"OrderService/router"
	"OrderService/Repo"
	"OrderService/controller"
	"OrderService/service"
)

func LoadRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	rtr := router.Group("api/v1")

	orderRepo := Repo.NewOrderRepo(db)
	orderService := service.NewOrderService(orderRepo)
	orderController := controller.NewOrderController(orderService)

	//rtr.GET("/order")     // get all orders
	//rtr.GET("/order/:id") //get order by id

	rtr.POST("/order", orderController.CreateOrder())
	rtr.POST("/order/bulk", orderController.CreateBulkOrder())

	return router
}
