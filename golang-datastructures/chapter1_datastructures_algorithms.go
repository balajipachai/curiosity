package main

import (
	"container/heap"
	"container/list"
	"encoding/json"
	"fmt"
	"strconv"
)

/**
* List is a sequential datastructure.
* In Go, lists can be accessed from the container/list package
* Different functions available on the list object are:
* func (*list.List).Back() *list.Element
* func (*list.List).Front() *list.Element
* func (*list.List).Init() *list.List
* func (*list.List).InsertAfter(v any, mark *list.Element) *list.Element
* func (*list.List).InsertBefore(v any, mark *list.Element) *list.Element
* func (*list.List).Len() int
* func (*list.List).MoveAfter(e *list.Element, mark *list.Element)
* func (*list.List).MoveBefore(e *list.Element, mark *list.Element)
* func (*list.List).MoveToBack(e *list.Element)
* func (*list.List).MoveToFront(e *list.Element)
* func (*list.List).PushBack(v any) *list.Element
* func (*list.List).PushBackList(other *list.List)
* func (*list.List).PushFront(v any) *list.Element
* func (*list.List).PushFrontList(other *list.List)
* func (*list.List).Remove(e *list.Element) any
 */

// The function creates a linked list of integers and prints its elements.
func createListAndPrint() {
	intList := list.New().Init()
	intList.PushBack(5)
	intList.PushBack(8)
	intList.PushBack(16)

	fmt.Println("Printing integer lists")
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}

/*
The function "tuples" returns the square and cube of a given number and an error if any.
A tuple is a finite sorted list of elements. It is a datastructure that groups data.
*/
func tuples(num int) (square, cube int, err error) {
	square = num * num
	cube = square * num
	return
}

/**
* An integer heap
* Package heap provides heap operations for any type that implements heap Interface.
**** A heap is a tree with the property that each node is the minimum-valued node in its subtree.
 */
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() any {
	oldSlice := *h
	n := len(oldSlice)
	x := oldSlice[n-1]
	*h = oldSlice[0 : n-1]
	return x
}

/*
COMPOSITE PATTERN: Answer Question 1
*/
// The below code defines an interface named IComposite with a single method named perform().
// @property perform - perform is a method signature that belongs to the IComposite interface. Any
// struct that implements this interface must also implement the perform method. The purpose of the
// perform method is specific to the implementation of the struct that implements the IComposite
// interface.
type IComposite interface {
	perform()
}

// The below code defines a struct type called Leaflet with a single field named "name" of type string.
// @property {string} name - The `name` property is a string type field that represents the name of a
// `Leaflet` object.
type Leaflet struct {
	name string
}

// The `perform()` function is a method of the `Leaflet` struct that implements the `IComposite`
// interface. It simply prints out the name of the leaflet object.
func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet " + leaf.name)
}

// The type "Branch" represents a tree structure with leafs and branches, identified by a name.
// @property {[]Leaflet} leafs - A slice of Leaflet structs that belong to this branch.
// @property {string} name - The name property is a string that represents the name of the branch.
// @property {[]Branch} branches - `branches` is a property of the `Branch` struct which is a slice of
// `Branch` type. It represents the sub-branches of the current branch. Each sub-branch is also a
// `Branch` struct with its own set of properties.
type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

// The `perform()` function is a method of the `Branch` struct that implements the `IComposite`
// interface. It prints out the name of the branch, its leafs, and sub-branches. It then iterates over
// each leaf and sub-branch and calls their respective `perform()` methods recursively. This allows for
// a tree-like structure to be printed out, with each branch and leaf being displayed in a hierarchical
// manner.
func (branch *Branch) perform() {
	fmt.Println("Branch: " + branch.name)
	fmt.Println("leafs: ", branch.leafs)
	fmt.Println("branches: ", branch.branches)

	for _, leaf := range branch.leafs {
		leaf.perform()
	}

	for _, branch := range branch.branches {
		fmt.Println("inside branch.branches")
		fmt.Println("branch: ", branch)
		branch.perform()
	}
}

// The `add` function is a method of the `Branch` struct that takes a `Leaflet` object as input and
// adds it to the `leafs` slice of the `Branch` object. It does this by using the `append` function to
// add the `Leaflet` object to the end of the `leafs` slice.
func (branch *Branch) add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

// The `addBranch` method is a function of the `Branch` struct that takes a `Branch` object as input
// and adds it to the `branches` slice of the `Branch` object. It does this by using the `append`
// function to add the `Branch` object to the end of the `branches` slice.
func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

// The `getLeaflets()` method is a function of the `Branch` struct that returns a slice of `Leaflet`
// objects. It does this by simply returning the `leafs` property of the `Branch` object.
func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

/*
* ADAPTER STRUCTURAL DESIGN PATTERN
* The adapter pattern provides a wrapper with an interface required by the API client to link
* incompatible types and act as a translator between the two types.
 */
// The below code defines an interface named IProcess with a single method named process().
// @property process - process is a method signature that defines a behavior that any struct or type
// implementing the IProcess interface must implement. It indicates that any type implementing this
// interface must have a method named "process" with no parameters and no return value.
type IProcess interface {
	process()
}

// The type Adapter contains an Adaptee object.
// @property {Adaptee} adaptee - The "adaptee" property is a variable of type "Adaptee", which is
// likely a reference to an object that needs to be adapted to work with a different interface or
// system. The Adapter struct is designed to provide a way to adapt the Adaptee object to work with a
// different
type Adapter struct {
	adaptee Adaptee
}

// The below code defines a struct type called Adaptee with an integer field called adapterType.
// @property {int} adapterType - The `adapterType` property is an integer that represents the type of
// adapter being used. It could be used to differentiate between different types of adapters and
// perform specific actions based on the adapter type.
type Adaptee struct {
	adapterType int
}

// This is a method of the `Adapter` struct that implements the `IProcess` interface. It prints out the
// string "Adapter process" and then calls the `convert()` method of the `adaptee` property of the
// `Adapter` object. The purpose of this method is to provide a way to adapt the `Adaptee` object to
// work with a different interface or system.
func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}

// The `convert()` function is a method of the `Adaptee` struct. It simply prints out the string
// "Adaptee convert method". This function is used in the Adapter pattern to demonstrate how an Adaptee
// object can be adapted to work with a different interface or system.
func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

/*
* BRIDGE STRUCTURAL DESIGN PATTERN
* Bridge decouples the implementation from the abstraction. The abstract base class can be
* subclassed to provide different implementations and allow implementation details to be
* modified easily.
 */

// The below code defines an interface called IDrawShape with a method signature for drawing a shape.
// @property drawShape - drawShape is a method signature of an interface called IDrawShape. It takes
// two arrays of 5 float32 values each, representing the x and y coordinates of a shape, and does not
// return anything. Any struct that implements this interface must provide its own implementation of
// the drawShape method.
type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

// The below code defines an empty struct called DrawShape in the Go programming language.
type DrawShape struct{}

// This is a method implementation of the `IDrawShape` interface for the `DrawShape` struct. It takes
// two arrays of 5 float32 values each, representing the x and y coordinates of a shape, and prints the
// string "Drawing Shape" to the console. This method can be used to draw a shape using the x and y
// coordinates provided as input.
func (drawShape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}

// The below code defines an interface called IContour with two methods, drawContour and
// resizeByFactor.
// @property drawContour - This is a method that takes in two arrays of 5 float values each,
// representing the x and y coordinates of points on a contour, and draws the contour using those
// points.
// @property resizeByFactor - `resizeByFactor` is a method that takes an integer `factor` as input and
// resizes the contour by multiplying the coordinates of the contour by the factor. For example, if the
// factor is 2, the contour will be doubled in size, and if the factor is 0.5 the contour will be half in size.
type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

// The type DrawContour represents a shape with five x and y coordinates, a shape type, and a factor.
// @property x - An array of 5 float32 values representing the x-coordinates of the contour points.
// @property y - `y` is an array of 5 `float32` values that represent the y-coordinates of the vertices
// of a contour shape.
// @property {DrawShape} shape - `shape` is a property of the `DrawContour` struct and it represents
// the shape of the contour. It is of type `DrawShape`, which is likely an enum or a custom type that
// defines different shapes that the contour can take.
// @property {int} factor - The "factor" property is an integer that represents the scaling factor for
// the contour. It is used to adjust the size of the contour when it is drawn.
type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

// The below code is defining a method called `drawContour` for a struct type `DrawContour`. This
// method takes two arrays of `float32` values as arguments and prints a message "Drawing Contour". It
// also calls the `drawShape` method of the `shape` field of the `DrawContour` struct, passing in the
// `x` and `y` fields of the same struct as arguments.
func (contour DrawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

// The below code is defining a method called `resizeByFactor` for a struct type `DrawContour` in the
// Go programming language. This method takes an integer parameter called `factor` and sets the
// `factor` field of the `DrawContour` struct to the value of `factor`. This method can be used to
// resize the contour drawn by the `DrawContour` struct by a given factor.
func (contour DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

/*
* DECORATOR STRUCTURAL DESIGN PATTERN
* In a scenario where class responsibilities are removed or added, the decorator pattern is
* applied. The decorator pattern helps with subclassing when modifying functionality,
* instead of static inheritance. An object can have multiple decorators and run-time
* decorators.
 */

// The below code defines an empty struct named ProcessClass in the Go programming language.
type ProcessClass struct{}

// The below code is defining a method called `process` for a struct type `ProcessClass` in the Go
// programming language. The method simply prints the string "ProcessClass process" to the console.
func (process *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

// The ProcessDecorator type is a struct that contains a pointer to an instance of the ProcessClass.
// @property processInstance - processInstance is a pointer to an instance of the ProcessClass struct.
// It allows the ProcessDecorator struct to access and modify the properties and methods of the
// ProcessClass struct.
type ProcessDecorator struct {
	processInstance *ProcessClass
}

// The below code is defining a method called `process()` for a struct type `ProcessDecorator`. This
// method checks if the `processInstance` field of the struct is `nil`. If it is `nil`, it prints
// "ProcessDecorator processs". If it is not `nil`, it prints "ProcessDecorator process and " and then
// calls the `process()` method of the `processInstance` field.
func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator processs")
	} else {
		fmt.Printf("ProcessDecorator process and ")
		decorator.processInstance.process()
	}
}

/*
* FACADE STRUCTURAL DESIGN PATTERN
* Facade is used to abstract subsystem interfaces with a helper. The facade design pattern is
* used in scenarios when the number of interfaces increases and the system gets complicated.
* Facade is an entry point to different subsystems, and it simplifies the dependencies
* between the systems. The facade pattern provides an interface that hides the
* implementation details of the hidden code.
 */

// The code defines a Go struct type called "Account" with two fields: "id" and "accountType".
// @property {string} id - The id property is a string that represents the unique identifier of an
// account. It is used to differentiate one account from another.
// @property {string} accountType - The accountType property is a string that represents the type of
// account. It could be a checking account, savings account, credit card account, etc.
type Account struct {
	id          string
	accountType string
}

// The below code is defining a method called `create` for the `Account` struct in Go. This method
// takes a string parameter `accountType` and sets the `accountType` field of the `Account` struct to
// the value of `accountType`. It then returns a pointer to the `Account` struct. The method also
// prints a message to the console indicating that an account is being created with a specific type.
func (account *Account) create(accountType string) *Account {
	fmt.Println("account creation with type")
	account.accountType = accountType
	return account
}

// The below code is defining a method called `deleteById` for a struct type `Account`. This method
// takes a string parameter `id` and prints a message "delete by account id" to the console. However,
// it does not actually delete anything from the account object or any external storage.
func (account *Account) deleteById(id string) {
	fmt.Println("delete by account id")
}

// The below code defines a struct type called Customer with two fields, name and id.
// @property {string} name - The name property is a string that represents the name of a customer.
// @property {int} id - The "id" property is an integer type field that represents a unique identifier
// for a customer. It can be used to differentiate one customer from another and can be used as a
// reference when retrieving or updating customer information.
type Customer struct {
	name string
	id   int
}

// The below code is defining a method called `create` for the `Customer` struct in Go. This method
// takes a string argument `name` and sets the `name` field of the `customer` object to the value of
// `name`. It then returns a pointer to the `customer` object. The method also prints "creating
// customer" to the console.
func (customer *Customer) create(name string) *Customer {
	fmt.Println("creating customer")
	customer.name = name
	return customer
}

// The Transaction type represents a financial transaction with an ID, amount, source account ID, and
// destination account ID.
// @property {string} id - The unique identifier for a transaction.
// @property {float32} amount - The amount property in the Transaction struct represents the amount of
// money involved in the transaction. It is of type float32, which means it can store decimal values up
// to a certain precision.
// @property {string} srcAccountId - The source account ID is a property of the Transaction struct that
// represents the ID of the account from which the transaction is being initiated.
// @property {string} destAccountId - The `destAccountId` property in the `Transaction` struct
// represents the unique identifier of the destination account for the transaction. This property is
// used to identify the account to which the `amount` is being transferred from the `srcAccountId`.
type Transaction struct {
	id            string
	amount        float32
	srcAccountId  string
	destAccountId string
}

// The below code is defining a method called `create` for a `Transaction` struct in Go. This method
// takes in three parameters: `srcAccountId` and `destAccountId` of type string, and `amount` of type
// float32. The method sets the values of these parameters to the corresponding fields in the
// `Transaction` struct and returns a pointer to the updated `Transaction` object. Additionally, the
// method prints the string "creating transaction" to the console.
func (transaction *Transaction) create(srcAccountId, destAccountId string, amount float32) *Transaction {
	fmt.Println("creating transaction")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

// The BranchManagerFacade type contains pointers to an Account, Customer, and Transaction.
// @property account - A pointer to an instance of the Account struct, which likely contains
// information about a bank account such as the account number, balance, and account holder's name.
// @property customer - The customer property is a pointer to a Customer object, which represents a
// customer of a bank branch. It likely contains information such as the customer's name, address,
// contact information, and account details.
// @property transaction - The `transaction` property is a pointer to a `Transaction` object. It is
// likely used to manage and perform transactions for the account and customer associated with the
// `BranchManagerFacade`.
type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

// The function returns a new instance of the BranchManagerFacade struct initialized with pointers to
// an Account, Customer, and Transaction.
func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

// The below code is defining a method called `createCustomerAccount` for the `BranchManagerFacade`
// struct This method takes two string parameters `customerName` and
// `accountType` and returns two pointers to `Customer` and `Account` structs.
func (facade *BranchManagerFacade) createCustomerAccount(customerName, accountType string) (*Customer, *Account) {
	customer := facade.customer.create(customerName)
	account := facade.account.create(accountType)
	return customer, account
}

// The below code is defining a method called `createTransaction` for the `BranchManagerFacade` struct
// This method takes three parameters: `srcAccountId` and `destAccountId`
// of type string, and `amount` of type float32.
func (facade *BranchManagerFacade) createTransaction(srcAccountId, destAccountId string, amount float32) *Transaction {
	tranaction := facade.transaction.create(srcAccountId, destAccountId, amount)
	return tranaction
}

/*
* FLYWEIGHT STRUCTURAL DESIGN PATTERN
* Flyweight is used to manage the state of an object with high variation. The pattern allows
* us to share common parts of the object state among multiple objects, instead of each object
* storing it. Variable object data is referred to as extrinsic state, and the rest of the object state
* is intrinsic.
 */

// The below code defines an interface called DataTransferObject that requires implementing a method
// called getId() that returns a string.
// @property {string} getId - This is a method signature for a function that returns a string
// representing the ID of an object that implements the DataTransferObject interface. This method is
// used to retrieve the unique identifier of an object, which can be useful for various operations such
// as database queries or object comparisons.
type DataTransferObject interface {
	getId() string
}

// The type `DataTransferObjectFactory` contains a pool of `DataTransferObject` objects.
// @property pool - The `pool` property is a map that stores instances of `DataTransferObject`. It is
// likely used to manage and reuse instances of `DataTransferObject` to improve performance and reduce
// memory usage. The keys of the map are strings that can be used to identify and retrieve specific
// instances of `Data
type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

// The below code defines a struct type called Employee with two fields, id and name.
// @property {string} id - The id property is a string that represents the unique identifier of an
// employee. It can be used to distinguish one employee from another in a database or system.
// @property {string} name - The "name" property is a string data type that represents the name of an
// employee.
type Employee struct {
	id   string
	name string
}

// The below code defines a struct type called Manager with three fields: id, name, and dept.
// @property {string} id - The id property is a string that represents the unique identifier of a
// Manager object.
// @property {string} name - The "name" property is a string type field that represents the name of a
// Manager.
// @property {string} dept - "dept" is short for department and is a property of the Manager struct. It
// is likely used to store the department that the manager works in.
type Manager struct {
	id   string
	name string
	dept string
}

// The below code defines a Go struct type called Address with fields for an ID, street lines 1 and 2,
// state, and city.
// @property {string} id - A unique identifier for the address.
// @property {string} streetLine1 - The first line of the street address for the address. For example,
// "123 Main St."
// @property {string} streetLine2 - `streetLine2` is a string property of the `Address` struct. It is
// likely used to store a second line of the street address, such as an apartment or suite number.
// @property {string} state - The "state" property in the "Address" struct represents the state or
// province of the address. For example, "California" or "Ontario".
// @property {string} city - The "city" property in the "Address" struct represents the name of the
// city where the address is located.
type Address struct {
	id          string
	streetLine1 string
	streetLine2 string
	state       string
	city        string
}

// The below code is implementing a factory design pattern in Go. It defines a method
// `getDataTransferObject` which takes a string parameter `dtoType` and returns an object of type
// `DataTransferObject`. The method first checks if an object of the given `dtoType` already exists in
// the `factory.pool` map. If it does, it returns the existing object. If not, it creates a new object
// of the specified type and adds it to the `factory.pool` map for future use. The code uses a switch
// statement to create different types of objects based on the `dtoType`
func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	dto := factory.pool[dtoType]
	if dto == nil {
		fmt.Println("new DTO of dtoType: " + dtoType)

		switch dtoType {
		case "customer":
			factory.pool[dtoType] = Customer{id: 1}
		case "employee":
			factory.pool[dtoType] = Employee{id: "2"}
		case "manager":
			factory.pool[dtoType] = Manager{id: "3"}
		case "address":
			factory.pool[dtoType] = Address{id: "4"}
		}
		dto = factory.pool[dtoType]
	}
	return dto
}

// The below code is defining a method called `getId()` for a struct type `Customer`. This method takes
// no arguments and returns a string. Inside the method, it prints a message "getting customer id" and
// then converts the `id` field of the `Customer` struct to a string using the `strconv.Itoa()`
// function and returns it.
func (customer Customer) getId() string {
	fmt.Println("gettting customer id")
	return strconv.Itoa(customer.id)
}

// The below code is defining a method called `getId()` for the `Employee` struct in the Go programming
// language. This method takes an `Employee` object as its receiver and returns the `id` field of that
// object. Additionally, it prints a message to the console indicating that it is getting the employee
// id.
func (employee Employee) getId() string {
	fmt.Println("getting employee id")
	return employee.id
}

// The below code is defining a method called `getId()` for a struct type `Manager`. This method takes
// no arguments and returns a string. When called, it prints "getting manager id" to the console and
// returns the `id` field of the `Manager` struct.
func (manager Manager) getId() string {
	fmt.Println("getting manager id")
	return manager.id
}

// The below code is defining a method called `getId()` for a struct type `Address`. This method
// returns the `id` field of the `Address` struct and prints a message "getting address id" to the
// console.
func (address Address) getId() string {
	fmt.Println("getting address id")
	return address.id
}

/*
* PRIVATE CLASS DATA STRUCTURAL DESIGN PATTERN
* The private class data pattern secures the data within a class. This pattern encapsulates the
* initialization of the class data. The write privileges of properties within the private class are
* protected, and properties are set during construction. The private class pattern prints the
* exposure of information by securing it in a class that retains the state. The encapsulation of
* class data initialization is a scenario where this pattern is applicable.
 */

// The type AccountDetails contains an id and accountType for an account.
// @property {string} id - The id property is a string that represents the unique identifier of an
// account. It is used to differentiate one account from another.
// @property {string} accountType - The "accountType" property is a string that represents the type of
// account. It could be a checking account, savings account, credit card account, etc.
type AccountDetails struct {
	id          string
	accountType string
}

// The type AccountP contains a pointer to AccountDetails and a string for the customer name.
// @property details - details is a pointer to an instance of the AccountDetails struct. It is used to
// store additional information about the account, such as the account number, balance, and transaction
// history.
// @property {string} CustomerName - CustomerName is a property of the AccountP struct which represents
// the name of the customer associated with the account.
type AccountP struct {
	details      *AccountDetails
	CustomerName string
}

// The below code is defining a method called `setDetails` for a struct type `AccountP`. This method
// takes two string arguments `id` and `accountType` and sets the `details` field of the `AccountP`
// struct to a pointer to a new `AccountDetails` struct with the given `id` and `accountType`.
func (account *AccountP) setDetails(id, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

// The below code is defining a method called `getId()` for a struct type `AccountP`. This method
// returns the `id` field of the `details` field of the `AccountP` struct.
func (account *AccountP) getId() string {
	return account.details.id
}

// The below code is defining a method called `getAccountType()` for a struct type `AccountP`. This
// method returns the `accountType` field of the `details` field of the `AccountP` struct.
func (account *AccountP) getAccountType() string {
	return account.details.accountType
}

/*
* PROXY STRUCTURAL DESIGN PATTERN
* The proxy pattern forwards to a real object and acts as an interface to others. The proxy
* pattern controls access to an object and provides additional functionality. The additional
* functionality can be related to authentication, authorization, and providing rights of access
* to the resource-sensitive object. The real object need not be modified while providing
* additional logic.
 */

// The below code defines an interface in Go called IRealObject that requires implementing a method
// called performAction().
// @property performAction - performAction is a method signature defined within the IRealObject
// interface. Any struct or class that implements this interface must provide an implementation for the
// performAction method.
type IRealObject interface {
	performAction()
}

// The below code defines a struct type named RealObject in the Go programming language.
type RealObject struct{}

// The below code is defining a method called `performAction()` for a struct type `RealObject` in the
// Go programming language. The method simply prints the string "RealObject performAction()" to the
// console.
func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}

// The VirtualProxy type is a struct that contains a pointer to a RealObject.
// @property realObject - The realObject property is a pointer to an instance of the RealObject struct.
// It is used in the VirtualProxy struct to hold a reference to the real object that the proxy is
// representing.
type VirtualProxy struct {
	realObject *RealObject
}

// The below code is implementing a method called `performAction()` for a struct called `VirtualProxy`.
// This method first prints a message "Virtual proxy performAction()", then checks if the `realObject`
// field of the `VirtualProxy` struct is `nil`. If it is `nil`, it creates a new instance of
// `RealObject` and assigns it to the `realObject` field. Finally, it calls the `performAction()`
// method of the `realObject`. This code is implementing the virtual proxy design pattern, where the
// `VirtualProxy` acts as a placeholder for the `RealObject
func (virtualProxy *VirtualProxy) performAction() {
	fmt.Println("Virtual proxy performAction()")
	if virtualProxy.realObject == nil {
		virtualProxy.realObject = &RealObject{}
	}
	virtualProxy.realObject.performAction()
}

// The main function demonstrates the usage of various Go language features such as tuples, integer
// heap, and composite pattern.
func main() {
	createListAndPrint()
	square, cube, err := tuples(5)
	if err != nil {
		fmt.Errorf("%w error occured", err)
	}
	fmt.Printf("Square and Cube of 5 is %d and %d\n", square, cube)
	fmt.Println("*************************Integer Heap******************************")
	h := &intHeap{10, 34, 5, 87}
	// Initialize heap
	heap.Init(h)
	heap.Push(h, 100)
	fmt.Printf("minimum of heap is %d\n", (*h)[0])
	fmt.Println("heap Len: ", h.Len())
	fmt.Println(h)
	fmt.Println("heap Less(0, 3) => : ", h.Less(0, 3))
	h.Swap(0, 3)
	fmt.Println(h)
	for h.Len() > 0 {
		fmt.Printf("%d\t", h.Pop())
	}
	fmt.Println("*************************Integer Heap******************************")
	fmt.Println()

	fmt.Println("*************************Composite Pattern******************************")
	branch := &Branch{name: "branch 1"}
	leaf1 := Leaflet{name: "leaf 1"}
	leaf2 := Leaflet{name: "leaf 2"}

	branch2 := Branch{name: "branch 2"}

	branch.add(leaf1)
	branch.add(leaf2)

	leaf3 := Leaflet{name: "leaf 3"}
	leaf4 := Leaflet{name: "leaf 4"}
	branch2.add(leaf3)
	branch2.add(leaf4)

	branch.addBranch(branch2)
	branch.perform()

	fmt.Println("*************************Adapter Pattern******************************")
	var processor IProcess = Adapter{}
	processor.process()

	fmt.Println("*************************Bridge Pattern******************************")
	x := [5]float32{1, 2, 3, 4, 5}
	y := [5]float32{1, 2, 3, 4, 5}
	var contour IContour = DrawContour{x, y, DrawShape{}, 2}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)

	fmt.Println("*************************Decorator Pattern******************************")
	var process = &ProcessClass{}
	var decorator = &ProcessDecorator{}
	decorator.process()
	decorator.processInstance = process
	decorator.process()

	fmt.Println("*************************Facade Pattern******************************")
	facade := NewBranchManagerFacade()
	customer, account := facade.createCustomerAccount("Balaji Shetty Pachai", "Current")
	fmt.Println(customer.name)
	fmt.Println(account.accountType)
	transaction := facade.createTransaction("1234567", "789456", 159753)
	fmt.Println(transaction.amount)

	fmt.Println("*************************Flyweight Pattern******************************")
	factory := DataTransferObjectFactory{make(map[string]DataTransferObject)}
	customerF := factory.getDataTransferObject("customer")
	fmt.Println("Customer: ", customerF.getId())
	employee := factory.getDataTransferObject("employee")
	fmt.Println("Employee: ", employee.getId())
	manager := factory.getDataTransferObject("manager")
	fmt.Println("Manager: ", manager.getId())
	address := factory.getDataTransferObject("address")
	fmt.Println("Address: ", address.getId())

	fmt.Println("*************************Private class data Pattern******************************")
	accountP := &AccountP{CustomerName: "Balaji Shetty Pachai"}
	accountP.setDetails("7777", "current")
	jsonAccount, _ := json.Marshal(accountP)
	fmt.Println("Private class hidden", string(jsonAccount))
	fmt.Println("Account Id: ", accountP.getId())
	fmt.Println("Account Type: ", accountP.getAccountType())

	fmt.Println("*************************Proxy Pattern******************************")
	object := VirtualProxy{}
	object.realObject = &RealObject{}
	object.performAction()

}

/**
Questions:
1. k-way merge algorithms?
2. Borrey-Moore algorithms?
3. Ukkonne algorithms?

Questions and exercises
1. Give an example where you can use a composite pattern.
Ans: Explained via code

2. For an array of 10 elements with a random set of integers, identify the maximum and minimum. Calculate the complexity of the algorithm.
Ans: O(n)

3. To manage the state of an object, which structural pattern is relevant?
Ans: Flyweight is used to manage the state of an object with high variation.

4. A window is sub-classed to add a scroll bar to make it a scrollable window. Which pattern is applied in this scenario?
Ans: Decorator Pattern.

5. Find the complexity of a binary tree search algorithm.
Ans: O(log n)

6. Identify the submatrices of 2x2 in a 3x3 matrix. What is the complexity of the algorithm that you have used?
Ans: O(n^3)

7. Explain with a scenario the difference between brute force and backtracking algorithms.
Ans: Brute force algorithm is the one wherein we try out each option to come to a solution whereas in backtracking algorithm the intermediate solutions are stored
	and those intermediate solutions are used instead of evaluating each solution everytime.
	For example: To search an element in an array we approach by using Brute Force however in case of fibonacci series we know what is fib(n-1) and fib(n-2) and thus use
	thier solutions directly, instead of evaluating them every now and then.
*/
