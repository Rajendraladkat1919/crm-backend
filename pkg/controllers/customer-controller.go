package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Rajendraladkat1919/crm-backend/pkg/models"
	"github.com/gorilla/mux"
)

var NewCustomer models.Customer

func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	newCustomer := models.GetAllCustomer()
	res, _ := json.Marshal(newCustomer)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	custId := vars["getCust"]
	Id, err := strconv.ParseInt(custId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}
	custDetails, _ := models.GetCustomerById(Id)
	res, _ := json.Marshal(custDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
