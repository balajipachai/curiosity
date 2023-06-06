package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

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

func GetCustomerById(customer Customer) Customer {
	db := GetConnection()

	fetchStmnt, err := db.Prepare("SELECT * FROM customers WHERE id = ?")

	if err != nil {
		log.Println("Error in fetch statement...")
		panic(err.Error())
	}

	rows, err := fetchStmnt.Query(customer.CustomerId)

	if err != nil {
		log.Println("Error in fetch statement execute...")
		panic(err.Error())
	}

	fetchedCustomer := Customer{}

	for rows.Next() {
		id := 0
		name := ""
		aadhar := ""

		err := rows.Scan(&id, &name, &aadhar)

		if err != nil {
			log.Println("Error while scanning row...")
			panic(err.Error())
		}

		fetchedCustomer.CustomerId = id
		fetchedCustomer.CustomerName = name
		fetchedCustomer.CustomerAadhar = aadhar
	}

	fmt.Println(fetchedCustomer)

	defer db.Close()

	return fetchedCustomer
}

var template_html = template.Must(template.ParseGlob("templates/*"))

func Home(writer http.ResponseWriter, req *http.Request) {
	customers := GetCustomers()
	log.Println(customers)
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func Create(writer http.ResponseWriter, req *http.Request) {
	template_html.ExecuteTemplate(writer, "Create", nil)
}

func Insert(writer http.ResponseWriter, req *http.Request) {
	customer := Customer{}
	customer.CustomerName = req.FormValue("name")
	customer.CustomerAadhar = req.FormValue("aadhar")
	InsertCustomer(customer)

	customers := GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func Update(writer http.ResponseWriter, req *http.Request) {
	customer := Customer{}
	customerId := req.FormValue("id")
	customerIdInt := 0
	customer.CustomerName = req.FormValue("name")
	customer.CustomerAadhar = req.FormValue("aadhar")

	fmt.Println("aadhar in update", req.FormValue("aadhar"))
	fmt.Println("name in update", req.FormValue("name"))

	fmt.Sscanf(customerId, "%d", &customerIdInt)
	customer.CustomerId = customerIdInt

	fmt.Println("request body in Update", customer)

	UpdateCustomer(customer)

	customers := GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func Delete(writer http.ResponseWriter, req *http.Request) {
	customer := Customer{}
	customerId := req.FormValue("id")
	customerIdInt := 0
	fmt.Sscanf(customerId, "%d", &customerIdInt)
	customer.CustomerId = customerIdInt
	DeleteCustomer(customer)

	customers := GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func Alter(writer http.ResponseWriter, req *http.Request) {
	customer := Customer{}
	customerId := req.FormValue("id")
	customerIdInt := 0
	fmt.Sscanf(customerId, "%d", &customerIdInt)
	customer.CustomerId = customerIdInt
	customer = GetCustomerById(customer)

	template_html.ExecuteTemplate(writer, "Update", customer)
}

func View(writer http.ResponseWriter, req *http.Request) {
	customer := Customer{}
	customerId := req.FormValue("id")
	customerIdInt := 0
	fmt.Sscanf(customerId, "%d", &customerIdInt)
	customer.CustomerId = customerIdInt
	customer = GetCustomerById(customer)

	customers := []Customer{customer}

	template_html.ExecuteTemplate(writer, "View", customers)
}

func main() {
	customer := Customer{
		CustomerId: 41,
	}
	GetCustomerById(customer)

	log.Println("Server started on localhost: http://127.0.0.1:8000")
	http.HandleFunc("/", Home)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/view", View)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8000", nil)

}
