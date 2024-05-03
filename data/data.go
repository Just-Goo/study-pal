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
	createTableSQL := `CREATE TABLE IF NOT EXISTS studypal (
		"id_note" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"word" TEXT,
		"definition" TEXT,
		"category" TEXT
	);`

	stmt, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("studypal table created")

}

func InsertNote(word, definition, category string) {
	insertNoteSQL := `INSERT INTO studypal (word, definition, category) VALUES (?, ?, ?)`

	stmt, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = stmt.Exec(word, definition, category)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("study note inserted successfully")

}

func DisplayAllNotes() {
	rows, err := db.Query(`SELECT id_note, word, definition, category FROM studypal`)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var idNote int
		var word, definition, category string

		rows.Scan(&idNote, &word, &definition, &category)
		log.Printf("[ID => %d, Word => %s, Definition => %s, Category => %s]", idNote, word, definition, category)
	}

	if rows.Err() != nil {
		log.Fatal(rows.Err())
	}
}
