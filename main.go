package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	seconds := 1
	for {
		fmt.Println("Prodding MariaDB server...")
		db, err := sql.Open("mysql", os.Getenv("MARIADB_ADDR"))
		if err == nil {
			err = db.Ping()
			if err == nil {
				_ = db.Close()
				fmt.Println("Server online!")
				os.Exit(0)
			}
			fmt.Printf("MariaDB ping failed: %s. Trying again in %d seconds.\n", err, seconds)
		} else {
			fmt.Printf("MariaDB connection failed: %s. Trying again in %d seconds.\n", err, seconds)
		}
		time.Sleep(time.Duration(seconds)*time.Second)
		seconds++
		if seconds > 3 {
			seconds = 3
		}
	}
}
