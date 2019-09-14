package data

// Page defines a test paper page structure.
type Page struct {
	ID     string `json:"id"`
	Number string `json:"number"`
	TestID string `json:"testID"`
	UserID string `json:"userID"`
}
