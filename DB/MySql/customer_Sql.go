package MySql

import (
	"database/sql"
	"food_delivery/internal"

	_ "github.com/go-sql-driver/mysql"
)

type customerConnections struct {
	DB *sql.DB
}

func (c *customerConnections) insertCustomer(customer internal.Customer) error {
	query := "INSERT INTO customers (name, address, phone, email, password) VALUES (?, ?, ?, ?, ?)"
	_, err := c.DB.Exec(query, customer.Name, customer.Address, customer.Phone, customer.Email, customer.Password)
	return err
}

func (c *customerConnections) getCustomerByID(id int) (*internal.Customer, error) {
	query := "SELECT name, address, phone, email, password FROM customers WHERE id = ?"
	row := c.DB.QueryRow(query, id)
	var customer internal.Customer
	err := row.Scan(&customer.Name, &customer.Address, &customer.Phone, &customer.Email, &customer.Password)
	if err != nil {
		return nil, err
	}
	customer.ID = id
	return &customer, nil
}

func (c *customerConnections) getAllCustomers() ([]internal.Customer, error) {
	query := "SELECT id, name, address, phone, email, password FROM customers"
	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []internal.Customer
	for rows.Next() {
		var customer internal.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Address, &customer.Phone, &customer.Email, &customer.Password)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c *customerConnections) updateCustomer(customer internal.Customer) error {
	query := "UPDATE customers SET name = ?, address = ?, phone = ?, email = ?, password = ? WHERE id = ?"
	_, err := c.DB.Exec(query, customer.Name, customer.Address, customer.Phone, customer.Email, customer.Password, customer.ID)
	return err
}

func (c *customerConnections) deleteCustomer(id int) error {
	query := "DELETE FROM customers WHERE id = ?"
	_, err := c.DB.Exec(query, id)
	return err
}
