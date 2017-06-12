package models

// Order reprentes a order
type Order struct {
	ID        int    `db:"id" json:"id"`
	UUID      string `db:"uuid" json:"uuid"`
	Status    string `db:"status" json:"status"`
	CreatedAt string `db:"created_at" json:"created_at"`
}
