package routes

import (
	"github.com/gorilla/mux"
	"githun.com/RajendraLadkat1919/pkg/controllers"
)

var AddCustomerRouters = func(router *mux.Router) {
	router.HandleFunc("/customer/", controllers.AddCustomer).Methods("POST")
	router.HandleFunc("/customer/", controllers.GetCustomer).Methods("GET")
	router.HandleFunc("/customer/{custId}", controllers.GetCustomerById).Methods("GET")
	router.HandleFunc("/customer/{custId}", controllers.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customer/{custId}", controllers.DeleteCustomerById).Methods("DELETE")
}
