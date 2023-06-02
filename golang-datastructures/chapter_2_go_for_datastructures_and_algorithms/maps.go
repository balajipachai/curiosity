package main

/**
* Maps in GO can be thought of as JSON objects or Hashtables that have a key and an associated value.
 */

import "fmt"

func main() {
	couples := map[string]string{}

	couples["Vaishali"] = "Manjunath"
	couples["Vishakha"] = "Nischith"
	couples["Balaji"] = "Anjali"
	couples["Ketan"] = "Rupali"
	couples["Ajit"] = "Aishwarya"

	fmt.Println("Printing all couples")
	for key, value := range couples {
		fmt.Println(key, "\t", value)
	}

	// Delete
	delete(couples, "Vaishali")
	fmt.Println("****************After delete****************")
	for key, value := range couples {
		fmt.Println(key, "\t", value)
	}
}
