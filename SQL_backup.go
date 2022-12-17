package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydatabase")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Create a backup of the database
	backupQuery := `mysqldump mydatabase > mydatabase.sql`
	_, err = db.Exec(backupQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Database backed up successfully")
}
