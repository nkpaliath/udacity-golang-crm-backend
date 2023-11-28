package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nkpalaith/udacity-golang-crm-backend/controller"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./static")))
	router.HandleFunc("/customers", controller.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers", controller.AddCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", controller.GetCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", controller.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", controller.DeleteCustomer).Methods("DELETE")

	fmt.Println("Server started on port 8000")
	http.ListenAndServe("localhost:8000", router)
}
