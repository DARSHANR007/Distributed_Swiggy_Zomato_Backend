package customer

import "food_delivery/cmd/restaurant"

func SearchDish(r *restaurant.Restaurant, dishName string, nos int) (*restaurant.Dish, bool) {
	if _, exists := r.Menu[dishName]; !exists {
		return nil, false
	}

	if r.Menu[dishName].Stock < nos {
		return nil, false
	}

	return r.Menu[dishName], true
}
