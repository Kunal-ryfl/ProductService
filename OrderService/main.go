package main

import (
	"OrderService/config"
	"OrderService/router"
	"net/http"
)

func main() {

	port := ":8081"

	db := config.Db()
	rtr := router.LoadRoutes(db)

	err := http.ListenAndServe(port, rtr)
	if err != nil {
		return
	}

}
