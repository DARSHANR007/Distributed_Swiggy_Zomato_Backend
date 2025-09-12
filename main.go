package main

import (
	"fmt"
	"food_delivery/MYSQL"
)

func main() {

	db := MYSQL.ConnectToDatabase()
	defer db.Close()
	fmt.Println("Database connection established:", db)
}
