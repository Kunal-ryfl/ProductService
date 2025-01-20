package main

import (
	"ProductService/config"
	"ProductService/router"
	"net/http"
)

func main() {

	port := ":8080"

	db := config.Db()
	rtr := router.LoadRoutes(db)

	err := http.ListenAndServe(port, rtr)
	if err != nil {
		return
	}

}
