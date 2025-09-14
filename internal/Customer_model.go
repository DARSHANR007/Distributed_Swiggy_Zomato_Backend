package internal

import (
	"food_delivery/cmd/restaurant"
)

type Customer struct {
	Name     string
	Address  string
	Phone    string
	Email    string
	Password string
	Orders   map[string][]OrderHistory
}

type OrderHistory struct {
	ID           int
	CustomerID   int
	RestaurantID int32
	Dishes       map[int]*restaurant.Dish
	TotalAmount  float64
}
