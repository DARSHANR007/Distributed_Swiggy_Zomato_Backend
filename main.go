package main

import (
	"fmt"
	"food_delivery/DB/rdbms"
)

func main() {
	// Initialize random seed
	// Connect to database
	rdbms.ConnectToDatabase()

	fmt.Println("Customer service is running...")

}
