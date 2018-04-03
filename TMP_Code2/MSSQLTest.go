package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	condb, errdb := sql.Open("mssql", "server=SAHIDUL\\MSSQLSERVER17;user id=dev;password=dev321;database=DevTest;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}
	var (
		sqlversion string
	)
	rows, err := condb.Query("select @@servername;")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&sqlversion)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(sqlversion)
	}
	defer condb.Close()
}

//Notes: go get github.com/denisenkom/go-mssqldb
