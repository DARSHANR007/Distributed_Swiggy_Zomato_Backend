package internal

type Restaurant struct {
	Name     string
	ID       int32
	Address  []string
	Phone    string
	Email    string
	Password string
	Menu     map[string]*Dish
}

type Dish struct {
	Name         string
	ID           int
	Price        float64
	Description  string
	RestaurantID int
	Available    bool
	Category     string
	Stock        int
	Cuisine      string
	Rating       float64
	Reviews      []string
	Tags         []string
}
