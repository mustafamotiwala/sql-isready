package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	var (
		userid   = flag.String("U", "", "login_id")
		password = flag.String("P", "", "password")
		server   = flag.String("S", "localhost", "server_name[\\instance_name]")
		database = flag.String("d", "", "db_name")
	)
	flag.Parse()

	dsn := "server=" + *server + ";user id=" + *userid + ";password=" + *password + ";database=" + *database
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		os.Exit(1)
	}
	err = db.Close()
	if err != nil {
		fmt.Println("Error disconnecting from the database. 🧐")
		os.Exit(1)
	}
	os.Exit(0)
}