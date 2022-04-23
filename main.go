package main

import (
	"go_simple_api/controller"
	"go_simple_api/model"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	model.Init()
	routing()
}

func routing() {
	r := mux.NewRouter()
	userController := controller.NewUserController()
	r.HandleFunc("/create", userController.Create).Methods("POST")

	http.ListenAndServe(":8080", r)

}
