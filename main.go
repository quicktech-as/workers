package main

import (
	"workers/cmd"
	"workers/connection"
	"workers/models"
)

func main() {
	newOrdersCh := make(chan models.Order)

	db, err := connection.Get()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	for {
		cmd.GetNewOrders(newOrdersCh, db)

		go func() {
			select {
			case order := <-newOrdersCh:
				go cmd.ProcessOrder(order, db)
			}
		}()
	}
}
