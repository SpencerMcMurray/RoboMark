package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func init_db() {
	database, _ := sql.Open("sqlite3", "./robomark.db")

	// Create keyword db
	statement, _ := database.Prepare(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY, 
		name TEXT, 
		isTeacher INTEGER)`)
	statement.Exec()

	// Create users db
	statement, _ := database.Prepare(`CREATE TABLE IF NOT EXISTS tests 
	(id INTEGER PRIMARY KEY,
		name TEXT, 
		markNum INTEGER,
		markDenom INTEGER,
	
		userID INTEGER)`) // ID of the teacher
	statement.Exec()

	// Create questions db
	statement, _ := database.Prepare(`CREATE TABLE IF NOT EXISTS questions (
		id INTEGER PRIMARY KEY,
		expectedAnswers TEXT, 
		markNum INTEGER,
		markDenom INTEGER,
		
		pageID INTEGER,
		testID INTEGER,
		userID INTEGER 
	)`)
	statement.Exec()

	// Create pages db
	statement, _ := database.Prepare(`CREATE TABLE IF NOT EXISTS pages (
		id INTEGER PRIMARY KEY,
		number INTEGER, 
		markNum INTEGER,
		markDenom INTEGER,

		testID INTEGER,
		userID INTEGER )`) // ID of the teacher
	statement.Exec()

	// Creates answers/keyword db
	statement, _ := database.Prepare(`CREATE TABLE IF NOT EXISTS keywords (
		id INTEGER PRIMARY KEY,
		questionID INTEGER, 
		string TEXT,

		pageID INTEGER,
		testID INTEGER,
		userID INTEGER)`) // ID of the teacher
	statement.Exec()

	rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}
