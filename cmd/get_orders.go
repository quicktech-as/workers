package cmd

import (
	"workers/connection"
	"workers/models"
)

// GetNewOrders returns new orders
func GetNewOrders() (orders []models.Order) {
	db, err := connection.Get()
	checkerr(err)

	rows, err := db.Queryx("SELECT id, uuid, status, created_at FROM orders WHERE status = 'open'")
	checkerr(err)

	for rows.Next() {
		var order models.Order

		err = rows.StructScan(&order)
		checkerr(err)

		orders = append(orders, order)
	}

	return orders
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
