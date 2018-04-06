package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	db, err := sql.Open("mssql", "sqlserver://sa:dev321@MLBD-SAHIDUL-IS//SQLEXPRESS2014?timeout=1m30s")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("select * from Employees")
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	// panic at 30 seconds...
	// panic: read tcp {my ip}->{server ip}: i/o timeout

}
