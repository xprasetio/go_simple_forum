package internalsql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("error connect to database: %v", err)
		return nil, err
	}
	return db, nil
}
