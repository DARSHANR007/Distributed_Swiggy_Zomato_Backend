package rdbms

import (
	"database/sql"
	"food_delivery/internal"
)

type RestaurantConnections struct {
	DB *sql.DB
}

func NewRestaurantConnections(restuaranthandler *sql.DB) *RestaurantConnections {
	return &RestaurantConnections{DB: restuaranthandler}
}

func (r *RestaurantConnections) InsertRestaurant(restaurant internal.Restaurant) error {
	query := "INSERT INTO restaurants (name, address, phone, email, cuisine) VALUES (?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, restaurant.Name, restaurant.Address, restaurant.Phone, restaurant.Email, restaurant.City)
	return err
}
