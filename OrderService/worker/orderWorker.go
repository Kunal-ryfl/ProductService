package worker

import (
	"OrderService/model"
	"fmt"
	"strconv"

	"OrderService/pusher"
	//"github.com/gin-gonic/gin"
	//"net/http"
	"OrderService/service"
)

func ProcessOrderRow(orderService service.OrderService, taskChannel chan model.Order, resultChannel chan service.Response) {
	fmt.Println("Worker started")
	client := pusher.PusherInit()

	newOrder := <-taskChannel

	orderID_String := strconv.Itoa(newOrder.ID)

	err := client.Trigger("my-channel", "orderEvent", orderID_String+"order processing started")
	if err != nil {

	}

	var response service.Response
	fmt.Println("in worker")
	response = orderService.CreateOrder(newOrder)

	resultChannel <- response
	fmt.Println("Worker Ended")

}
