package main

import (
	"go_simple_api/controller"
	"go_simple_api/model"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	routing()
}

func routing() {
	db := model.InitDB()
	r := mux.NewRouter()
	userController := controller.NewUserController(db)
	r.HandleFunc("/create", userController.Create).Methods("POST")

	http.ListenAndServe(":8080", r)

}
