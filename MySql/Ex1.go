package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/crud")
	if err != nil {
		log.Fatal(err)
	}

	var (
		id        int
		firstName string
	)
	rows, err := db.Query("select id, firstName from crud.employees where id = 1;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &firstName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, firstName)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
