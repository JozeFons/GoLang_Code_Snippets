package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

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

	// Prompt the user for confirmation
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Are you sure you want to restore the database from the backup file? (y/n): ")
	confirmation, _ := reader.ReadString('\n')
	confirmation = strings.TrimSpace(confirmation)
	if strings.ToLower(confirmation) != "y" {
		fmt.Println("Database restore cancelled")
		return
	}

	// Restore the database from the backup file
	restoreQuery := `mysql mydatabase < mydatabase.sql`
	_, err = db.Exec(restoreQuery)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Database restored successfully")
}
