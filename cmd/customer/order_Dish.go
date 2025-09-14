package customer

import (
	"food_delivery/cmd/restaurant"
	"food_delivery/internal"
)

func OrderDish(c *internal.Customer, r *restaurant.Restaurant, dishNames []string, nos int) (bool, float64) {
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
		ID: orderID,
		// CustomerID:   c.ID,
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
