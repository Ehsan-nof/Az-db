package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(){
	var err error 
	DB, err = sql.Open("sqlite3", "api.db") 
	if err != nil {
		log.Fatalf("could not open db: %v", err) 
    }

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables(){
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			name TEXT NOT NULL,
			PhoneNumber TEXT NOT NULL,
			address TEXT NOT NULL,
			birthYear INETEGER
		)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("could not create user table")
	}

	createPersonelTable := `
		CREATE TABLE IF NOT EXISTS personel(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			code INEGER NOT NULL,
			Description TEXT NOT NULL,
			birthYear INETEGER
		)
	`
	_, err = DB.Exec(createPersonelTable)
	if err != nil {
		panic("could not create user table")
	}


	createItemsTable := `
		CREATE TABLE IF NOT EXISTS events(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			price INTEGER NOT NULL,
			dateTime DATETIME NOT NULL
		)
	`
	_, err = DB.Exec(createItemsTable)

	if err != nil {
		panic("could not create event table")
	}

// 	createRegistrationTable := `
// 	CREATE TABLE IF NOT EXISTS registrations(
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		event_id INTEGER,
// 		user_id INTEGER,
// 		FOREIGN KEY(user_id) REFERENCES users(id),
// 		FOREIGN KEY(event_id) REFERENCES events(id)
// 	)
// `
// _, err = DB.Exec(createRegistrationTable)

// if err != nil {
// 	panic("could not create event table")
// }
}
