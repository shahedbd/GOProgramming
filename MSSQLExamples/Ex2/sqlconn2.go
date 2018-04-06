package main

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	//condb, errdb := sql.Open("mssql", "server= MLBD-SAHIDUL-IS\\SQLEXPRESS2014;user id=sa;password=dev321;database=DevTest;")

	query := url.Values{}
	query.Add("app name", "MyAppName")

	u := &url.URL{
		Scheme: "DevTest",
		User:   url.UserPassword("sa", "dev321"),
		Host:   fmt.Sprintf("%s:%d", "MLBD-SAHIDUL-IS\\SQLEXPRESS2014", 1434),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())

	//db.QueryContext(db, "select * from t where ID = @ID and Name = @p2;", sql.Named("ID", 6), "Bob")

}
