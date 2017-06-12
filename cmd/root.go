package cmd

import "fmt"

// Execute adds all child commands to the root command.
func Execute() {
	orders := GetNewOrders()

	fmt.Println("Running...")
	fmt.Printf("ORDERS: %v", orders)
}
