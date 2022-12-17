package main

import (
	"database/sql"
	"fmt"
	"os/exec"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to the database
	db, err := sql.Open("postgres", "user=postgres password=password host=localhost port=5432 dbname=mydatabase sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Create a backup of the database
	backupCommand := exec.Command("pg_dump", "-U", "postgres", "-Fc", "mydatabase", ">", "mydatabase.dump")
	err = backupCommand.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Database backed up successfully")
}
