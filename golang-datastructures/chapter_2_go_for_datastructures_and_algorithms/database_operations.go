package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// The below code defines a struct type called Customer with fields for customer ID, name, and Aadhar number.
// @property {int} CustomerId - an integer value that represents the unique identifier of a customer.
// @property {string} CustomerName - The name of the customer.
// @property {string} CustomerAadhar - CustomerAadhar is a property of the Customer struct that
// represents the Aadhar number of the customer. Aadhar is a unique identification number issued by the Indian government to its citizens.
type Customer struct {
	CustomerId     int
	CustomerName   string
	CustomerAadhar string
}

// The function returns a connection to a MySQL database.
func GetConnection() (database *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Inevergiveup100%"
	dbName := "go_database"

	db, error := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if error != nil {
		fmt.Println("Error connecting to database...")
		panic(error.Error())
	}
	return db
}

// This function retrieves all customers from a database table and returns them as a slice of Customer structs.
func GetCustomers() []Customer {
	db := GetConnection()

	rows, err := db.Query("SELECT * FROM customers ORDER BY id ASC")

	if err != nil {
		fmt.Println("Error while querying database...")
		panic(err.Error())
	}

	customer := Customer{}
	customers := []Customer{}

	for rows.Next() {
		customerId := 0
		customerName := ""
		customerAadhar := ""

		err := rows.Scan(&customerId, &customerName, &customerAadhar)

		if err != nil {
			fmt.Println("Error while scanning row...")
			panic(err.Error())
		}

		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.CustomerAadhar = customerAadhar

		customers = append(customers, customer)

	}
	// defer statement is used to schedule a function call to be executed later, typically just before the surrounding function returns.
	defer db.Close()

	return customers
}

// The function inserts a new customer into a database table using their name and Aadhar number.
func InsertCustomer(customer Customer) {
	db := GetConnection()

	insertStatement, err := db.Prepare("INSERT INTO customers (name, aadhar) values (?,?)")
	if err != nil {
		fmt.Println("Error in prepare statement...")
		panic(err.Error())
	}

	_, err = insertStatement.Exec(customer.CustomerName, customer.CustomerAadhar)

	if err != nil {
		fmt.Println("Error in executing insert statement...")
		panic(err.Error())
	}

	defer db.Close()
}

// The function updates a customer's name and Aadhar number in a database.
func UpdateCustomer(customer Customer) {
	db := GetConnection()

	updateStatement, err := db.Prepare("UPDATE customers SET name = ?, aadhar = ? WHERE id = ?")
	if err != nil {
		fmt.Println("Error in update statement...")
		panic(err.Error())
	}

	_, err = updateStatement.Exec(customer.CustomerName, customer.CustomerAadhar, customer.CustomerId)
	if err != nil {
		fmt.Println("Error in executing update statement...")
		panic(err.Error())
	}

	defer db.Close()
}

// The function deletes a customer from a database using their ID.
func DeleteCustomer(customer Customer) {
	db := GetConnection()

	deleteStatement, err := db.Prepare("DELETE FROM customers WHERE id = ?")

	if err != nil {
		fmt.Println("Error in delete statement...")
		panic(err.Error())
	}

	_, err = deleteStatement.Exec(customer.CustomerId)

	if err != nil {
		fmt.Println("Error in executing delete execute statement...")
		panic(err.Error())
	}

	defer db.Close()
}

// The function deletes all records from the "customers" table in a database.
func DeleteAll() {
	db := GetConnection()

	deleteAllStatement, err := db.Prepare("DELETE FROM customers")

	if err != nil {
		fmt.Println("Error in delete all statement...")
		panic(err.Error())
	}

	_, err = deleteAllStatement.Exec()

	if err != nil {
		fmt.Println("Error in executing delete all execute statement...")
		panic(err.Error())
	}

	defer db.Close()
}

func main() {
	DeleteAll()

	fmt.Println("Fetching customers from the database...")
	fmt.Println(GetCustomers())

	fmt.Println("Insert a new customer")
	customer := Customer{
		CustomerName:   "Balaji",
		CustomerAadhar: "654852159753",
	}
	InsertCustomer(customer)
	customer = Customer{
		CustomerName:   "Anjali",
		CustomerAadhar: "753159456852",
	}
	InsertCustomer(customer)
	customer = Customer{
		CustomerName:   "Akash",
		CustomerAadhar: "159456852753",
	}
	InsertCustomer(customer)
	customer = Customer{
		CustomerName:   "ToDelete",
		CustomerAadhar: "abcdefghijkl",
	}
	InsertCustomer(customer)
	customer = Customer{
		CustomerName:   "Ajit",
		CustomerAadhar: "951753456852",
	}
	InsertCustomer(customer)
	customer = Customer{
		CustomerName:   "Aishwarya",
		CustomerAadhar: "753654852159",
	}
	InsertCustomer(customer)
	customer = Customer{
		CustomerName:   "Vishakha",
		CustomerAadhar: "654852159753",
	}
	InsertCustomer(customer)
	customer = Customer{
		CustomerName:   "ToUpdate",
		CustomerAadhar: "654852159753",
	}
	InsertCustomer(customer)

	fmt.Println("Fetch customers after inserting a new customer")
	fmt.Println(GetCustomers())

	customers := GetCustomers()
	deleteId := 0
	updateId := 0
	for _, customer := range customers {
		if customer.CustomerName == "ToDelete" {
			deleteId = customer.CustomerId
		}

		if customer.CustomerName == "ToUpdate" {
			updateId = customer.CustomerId
		}
	}

	fmt.Println("Update existing customer")
	customer = Customer{
		CustomerId:     updateId,
		CustomerName:   "Vaishali",
		CustomerAadhar: "852456789123",
	}
	UpdateCustomer(customer)
	fmt.Println(GetCustomers())

	fmt.Println("Deleting existing customer")
	customer = Customer{
		CustomerId: deleteId,
	}
	DeleteCustomer(customer)
	fmt.Println(GetCustomers())
}
