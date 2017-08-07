package cmd

import (
	"fmt"
	"sync"

	"github.com/quicktech-as/workers/models"

	"github.com/jmoiron/sqlx"
)

// GetNewOrders returns new orders
func GetNewOrders(newOrdersCh chan models.Order, db *sqlx.DB, wg *sync.WaitGroup) {
	rows, err := db.Queryx("SELECT id, uuid, status, created_at FROM orders WHERE status = 'open'")
	checkerr(err)

	for rows.Next() {
		var order models.Order

		err = rows.StructScan(&order)
		checkerr(err)

		fmt.Printf("NEW ORDER: %s\n", order.UUID)

		db.MustExec("UPDATE orders SET status='processing' WHERE id=?", order.ID)
		fmt.Printf("PROCESSING ORDER: %s\n", order.UUID)

		wg.Add(1)
		newOrdersCh <- order
	}

	fmt.Println("EMPTY ORDERS")
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
