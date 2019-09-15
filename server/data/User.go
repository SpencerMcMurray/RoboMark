package data

// User defines a teacher structure.
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsTeacher bool   `json:"isTeacher"`
}
