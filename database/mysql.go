package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() error {
	var err error
	dsn := "root:Diag1234@tcp(localhost:3306)/application"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	fmt.Println("Connected to the database successfully")
	return nil
}
