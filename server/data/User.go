package data

// User defines a teacher structure.
type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsTeacher bool   `json:"isTeacher"`
}
