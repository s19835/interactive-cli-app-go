package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTable() {
	createTableSql := `CREATE TABLE IF NOT EXISTS studybuddy(
		"idNote" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"word" TEXT,
		"definition" TEXT,
		"category" TEXT
	);`

	statement, err := db.Prepare(createTableSql)
	if err != nil {
		log.Fatal(err)
	}

	statement.Exec()
	log.Println("studybuddy table created")
}

func InsertNote(word string, definition string, category string) {
	insertNoteSQL := `INSERT INTO studybuddy (word, definition, category)
	VALUES (?, ?, ?)`

	statment, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = statment.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Insert study note successfully")
}
