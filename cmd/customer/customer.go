package customer

import (
	"fmt"
	"food_delivery/DB/rdbms" // adjust import path
	"food_delivery/internal"
	"log"
)

var CustomerDB *rdbms.CustomerConnections // consistent name

func InitCustomerDB(dbHandler *rdbms.CustomerConnections) {
	if dbHandler == nil {
		log.Fatal("Database handler is nil")
	}
	CustomerDB = dbHandler
}

func Register(name, address, phone, email, password string) bool {
	c := internal.Customer{
		Name:     name,
		Address:  address,
		Phone:    phone,
		Email:    email,
		Password: password,
		Orders:   make(map[string][]internal.OrderHistory),
	}

	if CustomerDB != nil {
		err := CustomerDB.InsertCustomer(c)
		if err != nil {
			fmt.Println("❌ Failed to insert customer into DB:", err)
		} else {
			fmt.Println("✅ Customer saved to DB!")
		}
	}

	return true
}
