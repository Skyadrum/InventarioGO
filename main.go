package main

import (
	"InventarioGO/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	databaseConnection := database.DbConnection()

	//Logica

	defer databaseConnection.Close()
	fmt.Println(databaseConnection)

}
