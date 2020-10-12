package database

import (
	"database/sql"
)

func DbConnection() *sql.DB {
	connectionString := "root:admin@tcp(localhost:3306)/northwind"

	connectionDB, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error())
	}

	return connectionDB
}
