package internal

type Rider struct {
	Name        string
	ID          int
	Address     string
	Available   bool
	Phone       string
	Rating      float64
	RideHistory map[string][]Order
}

type RideHistory struct {
	ID           int
	RiderID      int
	CustomerID   int
	RestaurantID int32
	Dishes       map[int]*Dish
	TotalAmount  float64
	OrderAddress string
}
