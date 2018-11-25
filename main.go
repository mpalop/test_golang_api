package main

import (
	"fmt"
	"github.com/mpalop/test_golang_api/controllers"
)

func main() {
	fmt.Printf("REST test. Use:\n")
	fmt.Printf("/order <POST> to add orders\n")
	fmt.Printf("/order <GET> to return the list of orders\n")
	fmt.Printf("/order/:orderId <GET> to get back the order <orderId> or 404 if not found, or 400 if <orderId> is not an int\n")
	fmt.Printf("\nstarting server...\n")
	controllers.StartServer()
}
