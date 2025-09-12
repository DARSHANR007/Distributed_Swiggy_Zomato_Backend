package internal

import (
	"food_delivery/cmd/restaurant"
)

type Order struct {
	CustomerInfo   *Customer
	RestaurantInfo *restaurant.Restaurant
	OrderAddress   string
	Price          float64
}
