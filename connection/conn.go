package connection

import (
	"os"

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
		DB, err = sqlx.Connect("mysql", os.Getenv("MYSQL_URL"))
		if err != nil {
			return nil, err
		}
	}

	return DB, nil
}
