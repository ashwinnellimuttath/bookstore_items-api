package app

import (
	"itemsModule/controllers"
	"net/http"
)

func mapURLs() {
	router.HandleFunc("/ping", controllers.PingController.Ping)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
