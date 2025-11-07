package main

import "fmt"

func main() {

	var productName string = "Laptop"
	var productPrice float64 = 999.996
	var qty int = 5
	var inStock bool = true

	fmt.Printf("Product Name: %s\n", productName)
	fmt.Printf("Price: %.2f\n", productPrice)
	fmt.Printf("Quantity: %d\n", qty)
	fmt.Printf("In Stock: %v\n", inStock)

	if inStock {
		total := (productPrice * float64(qty))
		fmt.Printf("Total Price: %.2f", total)
	} else {
		fmt.Printf("This %v is unavailable\n", productName)
	}

}
