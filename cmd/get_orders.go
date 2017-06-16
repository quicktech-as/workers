package cmd

import (
	"fmt"
	"workers/models"

	"github.com/jmoiron/sqlx"
)

// GetNewOrders returns new orders
func GetNewOrders(newOrdersCh chan models.Order, db *sqlx.DB) {
	rows, err := db.Queryx("SELECT id, uuid, status, created_at FROM orders WHERE status = 'open'")
	checkerr(err)

	for rows.Next() {
		var order models.Order

		err = rows.StructScan(&order)
		checkerr(err)

		fmt.Printf("NEW ORDER: %s\n", order.UUID)

		db.MustExec("UPDATE orders SET status='processing' WHERE id=?", order.ID)
		fmt.Printf("PROCESSING ORDER: %s\n", order.UUID)

		newOrdersCh <- order
	}
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
