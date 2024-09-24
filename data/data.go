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

func DisplayAllNotes() {
	row, err := db.Query("SELECT * FROM sudybuddy ORDER BY word")

	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()

	for row.Next() {
		var idNote int
		var word string
		var definition string
		var category string

		row.Scan(&idNote, &word, &definition, &category)
		log.Println("[", category, "] ", word, "-", definition)
	}
}
