package main

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

func addTest(name string, markNum, markDenom, userID int8) {
	database, _ := sql.Open("sqlite3", "./robomark.db")
	statement, _ := database.Prepare(`
		INSERT INTO TABLE tests VALUES (name, markNum, markDenom, userID)
		(?, ?, ?, ?)
	`)

	statement.Exec(name, markNum, markDenom, userID)
}

func addPage(number, testID, userID string) {
	database, _ := sql.Open("sqlite3", "./robomark.db")
	statement, _ := database.Prepare(`
		INSERT INTO TABLE pages VALUES (number, testID, userID)
		(?, ?, ?, ?, ?, ?)
	`)
	
	statement.Exec(number, testID, userID)
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

func viewQuestion(pageID int8) []Question {
	database, _ := sql.Open("sqlite3", "./robomark.db")
	rows, _ := database.Query("SELECT id, expectedAnswers, markNum, markDenom, pageID, testID, userID FROM questions
	WHERE pageID = " + string(pageID))
	var id string
	var expectedAnswers string
	var markNum int8
	var markDenom int8
	var pageID string
	var testID string
	var userID string

	questions := []Question {}
	var question Question

	for rows.Next() {
		rows.Scan(&id, &expectedAnswers, &markNum, &markDenom, &pageID, &testID, $userID)
		question = Question{
			id, expectedAnswers, markNum, markDenom, pageID, testID, userID
		}
		questions = append(questions, question)
	}

	return questions
}

func viewPage(testID int8) []Page{
	database, _ := sql.Open("sqlite3", "./robomark.db")
	rows, _ := database.Query("SELECT id, number, testID, userID FROM pages
	WHERE testID = " + string(testID))
	var id string
	var number string
	var testID string
	var userID string

	pages := []page {}
	var page Page

	for rows.Next() {
		rows.Scan(&id, &name, &markNum, &markDenom, $userID)
		page = Page{
			id, name, markNum, markDenom, userID
		}
		pages = append(pages, page)
	}

	return pages
}

func viewTests(userID int8) []Test {
	database, _ := sql.Open("sqlite3", "./robomark.db")
	rows, _ := database.Query("SELECT id, name, markNum, markDenom, userID FROM tests
	WHERE userID = " + string(userID))
	var id string
	var name string
	var markNum int8
	var markDenom int8
	var userID string

	tests := []Test {}
	var test Test

	for rows.Next() {
		rows.Scan(&id, &name, &markNum, &markDenom, $userID)
		test = Test{
			id, name, markNum, markDenom, userID
		}
		tests = append(tests, test)
	}

	return tests
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
	statement, _ = database.Prepare(`CREATE TABLE IF NOT EXISTS tests (
		id INTEGER PRIMARY KEY,
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
