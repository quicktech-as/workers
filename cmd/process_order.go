package cmd

import (
	"fmt"
	"sync"

	"github.com/quicktech-as/workers/models"

	"github.com/jmoiron/sqlx"
)

// ProcessOrder process a order
func ProcessOrder(order models.Order, db *sqlx.DB, wg *sync.WaitGroup) {
	db.MustExec("UPDATE orders SET status='close' WHERE id=?", order.ID)
	fmt.Printf("ORDER PROCESSED : %s\n", order.UUID)
	fmt.Println("----------------------------------------------")
	wg.Done()
}
