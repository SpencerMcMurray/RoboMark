package data

import (
	"database/sql"
	"fmt"
	"strconv"

	// May need this.
	_ "github.com/go-sql-driver/mysql"
)

// AddUser .
func AddUser(name string, isTeacher bool) {
	db, err := sql.Open("mysql", "robomark:Gw3CVo~1Z4d~@tcp(den1.mysql6.gear.host)/robomark")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	statement, _ := db.Prepare(`
		INSERT INTO TABLE users VALUES (name, isTeacher)
		(?, ?)
	`)

	value := 0
	if isTeacher {
		value = 1
	}
	newValue := strconv.Itoa(value)

	statement.Exec(name, newValue)
}

// AddTest .
func AddTest(name string, markNum, markDenom, userID int) {
	db, err := sql.Open("mysql", "robomark:Gw3CVo~1Z4d~@tcp(den1.mysql6.gear.host)/robomark")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	statement, _ := db.Prepare(`
		INSERT INTO TABLE tests VALUES (name, markNum, markDenom, userID)
		(?, ?, ?, ?)
	`)

	statement.Exec(name, markNum, markDenom, userID)
}

// AddPage .
func AddPage(number, testID, userID string) {
	db, err := sql.Open("mysql", "robomark:Gw3CVo~1Z4d~@tcp(den1.mysql6.gear.host)/robomark")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	statement, _ := db.Prepare(`
		INSERT INTO TABLE pages VALUES (number, testID, userID)
		(?, ?, ?)
	`)

	statement.Exec(number, testID, userID)
}

// AddQuestion .
func AddQuestion(expectedAnswers string, markNum, markDenom int, pageID, testID, userID string) {
	db, err := sql.Open("mysql", "robomark:Gw3CVo~1Z4d~@tcp(den1.mysql6.gear.host)/robomark")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	statement, _ := db.Prepare(`
		INSERT INTO TABLE questions VALUES (name, markNum, markDenom, pageID, testID, userID)
		(?, ?, ?, ?, ?, ?)
	`)

	statement.Exec(expectedAnswers, markNum, markDenom, pageID, testID, userID)
}

// ViewQuestion .
func ViewQuestion(pageID int) []Question {
	db, err := sql.Open("mysql", "robomark:Gw3CVo~1Z4d~@tcp(den1.mysql6.gear.host)/robomark")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, _ := db.Query(`SELECT * FROM questions WHERE page = ?`, pageID)

	questions := []Question{}
	for rows.Next() {
		var id int
		var expectedAnswers string
		var markNum int
		var markDenom int
		var testID int
		var userID int

		rows.Scan(&id, &expectedAnswers, &markNum, &markDenom, &pageID, &testID, &userID)
		questions = append(questions, Question{id, expectedAnswers, markNum, markDenom, pageID, testID, userID})
	}

	return questions
}

// ViewPage .
func ViewPage(testID int) []Page {
	db, err := sql.Open("mysql", "robomark:Gw3CVo~1Z4d~@tcp(den1.mysql6.gear.host)/robomark")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, _ := db.Query(`SELECT * FROM pages WHERE test = ?`, testID)

	pages := []Page{}
	for rows.Next() {
		var id int
		var userID int
		var name string
		var markNum int
		var markDenom int

		rows.Scan(&id, &name, &markNum, &markDenom, &userID)
		pages = append(pages, Page{id, len(pages) + 1, testID, userID})
	}

	return pages
}

// ViewTests .
func ViewTests(userID int) []Test {
	db, err := sql.Open("mysql", "robomark:Gw3CVo~1Z4d~@tcp(den1.mysql6.gear.host)/robomark")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, _ := db.Query(`SELECT * FROM tests WHERE user = ?`, userID)
	tests := []Test{}

	for rows.Next() {
		var IDStr string
		var name string
		var markNum int
		var markDenom int
		rows.Scan(&IDStr, &name, &markNum, &markDenom, &userID)

		id, err := strconv.Atoi(IDStr)
		if err != nil {
			panic(err)
		}

		tests = append(tests, Test{id, name, markNum, markDenom, userID})
	}

	fmt.Println(tests)

	return tests
}
