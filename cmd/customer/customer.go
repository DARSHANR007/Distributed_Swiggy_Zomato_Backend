package customer

import (
	"fmt"
	"food_delivery/cmd/restaurant"
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

func SearchRestaurant(restaurantName string) (*restaurant.Restaurant, bool) {
	if hotel, exists := restaurant.RestaurantData[restaurantName]; exists {
		return hotel, true
	}
	return nil, false
}

func GetCustomerById(id int) *internal.Customer {
	return CustomerData[id]
}

func SearchDish(r *restaurant.Restaurant, dishName string, nos int) (*restaurant.Dish, bool) {
	if _, exists := r.Menu[dishName]; !exists {
		return nil, false
	}

	if r.Menu[dishName].Stock < nos {
		return nil, false
	}

	return r.Menu[dishName], true
}

func GetDish(c *internal.Customer, r *restaurant.Restaurant, dishNames []string, nos int) (bool, float64) {
	total := 0.0
	orderedDishes := make(map[int]*restaurant.Dish)

	for _, dn := range dishNames {
		if dish, ok := SearchDish(r, dn, nos); ok {
			total += dish.Price * float64(nos)
			orderedDishes[dish.ID] = dish
			// reduce stock
			r.Menu[dish.Name].Stock -= nos
		}
	}

	if len(orderedDishes) == 0 {
		return false, 0
	}

	orderID := len(c.Orders[r.Name]) + 1
	newOrder := internal.OrderHistory{
		ID:           orderID,
		CustomerID:   c.ID,
		RestaurantID: r.ID,
		Dishes:       orderedDishes,
		TotalAmount:  total,
	}

	if c.Orders == nil {
		c.Orders = make(map[string][]internal.OrderHistory)
	}
	c.Orders[r.Name] = append(c.Orders[r.Name], newOrder)

	return true, total
}
