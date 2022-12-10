package models

import (
	"github.com/Rajendraladkat1919/crm-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Customer struct {
	gorm.Model
	Name      string `gorm:""json:"name"`
	Id        int    `json:"id"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
	Contacted bool   `json:"contacted"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Customer{})
}

func (c *Customer) AddCustomer() *Customer {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func GetAllCustomer() []Customer {
	var Customer []Customer
	db.Find(&Customer)
	return Customer
}
func GetCustomerById(Id int64) (*Customer, *gorm.DB) {
	var getCust Customer
	db := db.Where("ID=?", Id).Find(&getCust)
	return &getCust, db
}
func DeleteCustomer(Id int64) Customer {
	var cust Customer
	db.Where("ID=?", Id).Delete(cust)
	return cust
}
