package db

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var sqlFolder = "db/sql/"

func InitDB() *sql.DB {
	log.Println("Initializing DB ...")
	var err error
	db, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	// create tables
	createTables()
	log.Println("DB initialization complete.")
	return db
}

func GetDB() *sql.DB {
	return db
}

func createTables() bool {
	// reads the sql file from the sql folder
	// and executes the sql statement
	log.Println("Creating tables ...")
	currentDir, _ := os.Getwd()
	log.Println("Current Directory: ", currentDir)
	sqlString, err := os.ReadFile(
		sqlFolder + "tables.sql",
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(sqlString))
	if err != nil {
		log.Fatal(err)
	}

	return true
}
