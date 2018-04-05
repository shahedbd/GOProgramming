package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	//condb, errdb := sql.Open("mssql", "server=MLBD-SAHIDUL-IS\\MSSQLSERVER17;user id=dev;password=dev321;database=DevTest;")
	condb, errdb := sql.Open("mssql", "server= MLBD-SAHIDUL-IS\\SQLEXPRESS2014;user id=sa;password=dev321;database=DevTest;")
	if errdb != nil {
		fmt.Println(" Error open db:", errdb.Error())
	}

	var (
		id           int
		FirstName    string
		LastName     string
		Department   string
		EmployeeType string
	)

	abc := condb.Stats()
	fmt.Println(abc)

	abc2 := condb.Ping()
	fmt.Println(abc2)

	rows, err := condb.Query("select * from Employees")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("select * from Employees")
	for rows.Next() {
		err := rows.Scan(&id, &FirstName, &LastName, &Department, &EmployeeType)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(id, FirstName, LastName, Department, EmployeeType)
	}
	defer condb.Close()

}
