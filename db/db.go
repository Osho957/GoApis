package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDB() {
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Error connecting with the database: " + err.Error())
	}
	// we will have a maximum of 10 open connections for the database 
	// if there are more than 10 queries, the 11th query will have to wait
	DB.SetMaxOpenConns(10)
	// if there are no queries, the database will keep 5 connections open
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users
	(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Error creating users table: " + err.Error())
	}
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events 
	(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL, 
	description TEXT NOT NULL, 
	location TEXT NOT NULL,
	datetime DATETIME NOT NULL, 
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)`
	// to create a table, we use the Exec method of the DB object
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Error creating events table: " + err.Error())
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations
	(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER,
	user_id INTEGER,
	FOREIGN KEY(event_id) REFERENCES events(id),
	FOREIGN KEY(user_id) REFERENCES users(id)
	)`
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic("Error creating registrations table: " + err.Error())
	}
}
