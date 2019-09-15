package data

// Page defines a test paper page structure.
type Page struct {
	ID     int `json:"id"`
	Number int `json:"number"`
	TestID int `json:"testID"`
	UserID int `json:"userID"`
}
