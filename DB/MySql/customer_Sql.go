
package MySql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type customerConnections struct {
	DB *sql.DB
}


func( c *customerConnections) insertCustomer(customer Customer) error {
	query := "INSERT INTO customers (name, address, phone, email, password) VALUES (?, ?, ?, ?, ?)"
	_, err := c.DB.Exec(query, customer.Name, customer.Address, customer.Phone, customer.Email, customer.Password)
	return err
}

func (c *customerConnections) getCustomerByID(id int) (*Customer, error) {
	query := "SELECT name, address, phone, email, password FROM customers WHERE id = ?"
	row := c.DB.QueryRow(query, id)
	var customer Customer
	err := row.Scan(&customer.Name, &customer.Address, &customer.Phone, &customer.Email, &customer.Password)
	if err != nil {
		return nil, err
	}	
	customer.ID = id
	return &customer, nil
}
func (c *customerConnections) getAllCustomers() ([]Customer, error) {
	query := "SELECT id, name, address, phone, email, password FROM customers"
	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Address, &customer.Phone, &customer.Email, &customer.Password)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}	