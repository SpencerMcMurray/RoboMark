package server

import (
	"database/sql"
	"strconv"
	// this comment is here because the linter said so
	_ "github.com/mattn/go-sqlite3"
)

// returns user id
func addUser(name string, isTeacher bool) {
	database, _ := sql.Open("sqlite3", "./robomark.db")
	statement, _ := database.Prepare(`
		INSERT INTO TABLE users VALUES (name, isTeacher)
		(?, ?)
	`)

	var value := 0
	if isTeacher {
		value = 1
	}
	var newValue := strconv.Itoa(value)

	statement.Exec(name, newValue)
}

func viewUsers() {

}

func addTest(name string, markNum, markDenom, userID int8) {
	database, _ := sql.Open("sqlite3", "./robomark.db")
	statement, _ := database.Prepare(`
		INSERT INTO TABLE tests VALUES (name, markNum, markDenom, userID)
		(?, ?, ?, ?)
	`)

	
	statement.Exec(name, markNum, markDenom, userID)
}

func viewTest() {

}

// expected answers are comma separated
func addQuestion(expectedAnswers string, markNum, markDenom int8, pageID, testID, userID string) {
	database, _ := sql.Open("sqlite3", "./robomark.db")
	statement, _ := database.Prepare(`
		INSERT INTO TABLE questions VALUES (name, markNum, markDenom, pageID, testID, userID)
		(?, ?, ?, ?, ?, ?)
	`)
	
	statement.Exec(expectedAnswers, markNum, markDenom, pageID, testID, userID)
}

func viewQuestion() {

}

func addPage(number, testID, userID string) {
	database, _ := sql.Open("sqlite3", "./robomark.db")
	statement, _ := database.Prepare(`
		INSERT INTO TABLE pages VALUES (number, testID, userID)
		(?, ?, ?, ?, ?, ?)
	`)
	
	statement.Exec(number, testID, userID)
}

func viewPage() {

}

func initDb() {
	database, _ := sql.Open("sqlite3", "./robomark.db")

	// Create keyword db
	statement, _ := database.Prepare(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY, 
		name TEXT, 
		isTeacher INTEGER)`)
	statement.Exec()

	// Create users db
	statement, _ = database.Prepare(`CREATE TABLE IF NOT EXISTS tests 
	(id INTEGER PRIMARY KEY,
		name TEXT, 
		markNum INTEGER,
		markDenom INTEGER,
	
		userID INTEGER)`) // ID of the teacher
	statement.Exec()

	// Create questions db
	statement, _ = database.Prepare(`CREATE TABLE IF NOT EXISTS questions (
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
	statement, _ = database.Prepare(`CREATE TABLE IF NOT EXISTS pages (
		id INTEGER PRIMARY KEY,
		number INTEGER, 
		markNum INTEGER,
		markDenom INTEGER,

		testID INTEGER,
		userID INTEGER )`) // ID of the teacher
	statement.Exec()

	// rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	// var id int
	// var firstname string
	// var lastname string
	// for rows.Next() {
	// 	rows.Scan(&id, &firstname, &lastname)
	// 	fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	// }
}
