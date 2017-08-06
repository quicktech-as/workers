package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

	finish := make(chan os.Signal, 1)
	signal.Notify(finish, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case order := <-newOrdersCh:
				go cmd.ProcessOrder(order, db, &wg)
			case <-time.After(time.Second * 5):
				go cmd.GetNewOrders(newOrdersCh, db, &wg)
			case <-finish:
				fmt.Println("Exiting...")
				wg.Done()
			}
		}
	}()
	wg.Wait()
}
