package customer

import "food_delivery/cmd/restaurant"

func SearchRestaurant(restaurantName string) (*restaurant.Restaurant, bool) {
	if hotel, exists := restaurant.RestaurantData[restaurantName]; exists {
		return hotel, true
	}
	return nil, false
}
