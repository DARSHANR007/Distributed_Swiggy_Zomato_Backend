package rider

import (
	"food_delivery/internal"
)

type Rider struct {
	Name        string
	ID          int
	Address     string
	Available   bool
	Phone       string
	Rating      float64
	RideHistory map[string][]Order
}

type Order struct {
	CustomerInfo   *internal.Customer
	RestaurantInfo *internal.Restaurant
	OrderAddress   string
	Price          float64
}
