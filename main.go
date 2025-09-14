package main

import (
	"food_delivery/DB/rdbms"
	"food_delivery/cmd/customer"
)

func main() {
	db := rdbms.ConnectToDatabase()
	customerDB := rdbms.NewCustomerConnections(db)
	customer.InitCustomerDB(customerDB)
	customer.Register("John Doe", "123 Main St", "555-1234", "john@gmail.com", "securepassword")

}
