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
