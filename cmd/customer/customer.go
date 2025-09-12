package customer

import (
	"fmt"
	"food_delivery/internal"
)

var CustomerData = make(map[int]*internal.Customer)

func Register(name, address, phone, email, password string, id int) bool {
	if _, exists := CustomerData[id]; exists {
		fmt.Println("Customer already exists with ID:", id)
		return false
	}

	c := internal.Customer{
		Name:     name,
		Address:  address,
		Phone:    phone,
		Email:    email,
		ID:       id,
		Password: password,
		Orders:   make(map[string][]internal.OrderHistory),
	}

	CustomerData[c.ID] = &c
	return true
}

func GetCustomerById(id int) *internal.Customer {
	return CustomerData[id]
}
