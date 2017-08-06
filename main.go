package main

import (
	"sync"
	"time"

	"github.com/quicktech-as/workers/cmd"
	"github.com/quicktech-as/workers/connection"
	"github.com/quicktech-as/workers/models"
)

func main() {
	db, err := connection.Get()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	newOrdersCh := make(chan models.Order)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			select {
			case order := <-newOrdersCh:
				go cmd.ProcessOrder(order, db, &wg)
			case <-time.After(time.Second * 5):
				go cmd.GetNewOrders(newOrdersCh, db, &wg)
			}
		}
	}()
	wg.Wait()
}
