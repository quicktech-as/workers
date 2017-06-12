package connection

import (
	"github.com/jmoiron/sqlx"

	// Used mysql drive on sql
	_ "github.com/go-sql-driver/mysql"
)

var (
	// DB connection
	DB  *sqlx.DB
	err error
)

// Get get mysql connection
func Get() (*sqlx.DB, error) {
	if DB == nil {
		DB, err = sqlx.Connect("mysql", "root:root@tcp(192.168.99.100:3306)/workers")
		if err != nil {
			return nil, err
		}
	}

	return DB, nil
}
