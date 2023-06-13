package queue

import "fmt"

type Queue []*Order

// The type "Order" contains information about a customer's order, including priority, quantity, product, and customer name.
// @property {int} priority - The priority of the order, which could be used to determine the order in which it should be fulfilled or processed.
// @property {int} quantity - The quantity property in the Order struct represents the number of units of the product that the customer has ordered.
// @property {string} product - The name or identifier of the product being ordered.
// @property {string} customerName - The name of the customer who placed the order.
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

// It is a constructor method for the `Order` struct.
// It takes in four parameters: `priority`, `quantity`, `product`, and `customerName`, and sets the corresponding properties of the
// `Order` struct to these values.
// The `order` parameter is a pointer to the `Order` struct that is being constructed.
func (order *Order) New(priority, quantity int, product, customerName string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customerName = customerName
}

// This method is adding an `Order` to the `Queue`.
// It first checks if the `Queue` is empty, and if it is, it adds the `Order` directly to the `Queue`.
// If the `Queue` is not empty, it traverses the `Queue` and compares the priority of the new `Order` with the
// priorities of the existing `Orders` in the `Queue`.
// It then adds the new `Order` to the `Queue` in the correct position based on its priority.
// If the new `Order` has the lowest priority, it is added to the end of the `Queue`.
func (queue *Queue) Add(order *Order) {
	// When the queue is empty then directly add the element
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		appended := false
		// When the queue is not empty, traverse the queue and depending on the priority
		// add the element to the queue
		for index, queueOrder := range *queue {
			if order.priority > queueOrder.priority {
				*queue = append(((*queue)[:index]), append(Queue{order}, (*queue)[index:]...)...)
				appended = true
				break
			}
		}
		// If appended is false, then add the element to the end of the queue
		if !appended {
			fmt.Println("inside not appended")
			*queue = append(*queue, order)
		}
	}
}

// This is a method for deleting an `Order` from the `Queue`.
// It takes in an `Order` as a parameter and searches for it in the `Queue`.
// If the `Order` is found, it is removed from the `Queue`.
// The method uses a `for` loop to iterate over each `Order` in the `Queue` and compares the `product` and
// `customerName` properties of each `Order` with the corresponding properties of the `Order` passed as
// a parameter.
// If a match is found, the `Order` is removed from the `Queue` using the `append`
// function to create a new slice that excludes the `Order` being deleted.
// If no match is found, the `Queue` remains unchanged.
func (queue *Queue) Delete(order *Order) {
	for index, queueOrder := range *queue {
		if (queueOrder.product == order.product) && (queueOrder.customerName == order.customerName) {
			*queue = append((*queue)[:index], (*queue)[index+1:]...)
			break
		}
	}
}
