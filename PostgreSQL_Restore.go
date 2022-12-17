package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strings"

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
	restoreCommand := exec.Command("pg_restore", "-U", "postgres", "-d", "mydatabase", "mydatabase.dump")
	err = restoreCommand.Run()
	if err != nil {
    fmt.Println(err)
		return
	}

	fmt.Println("Database restored successfully")
}
	
