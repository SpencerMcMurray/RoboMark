package data

// Keyword defines a keyword a question might be looking for.
type Keyword struct {
	ID   string `json:"id"`
	Word string `json:"string"`

	QuestionID string `json:"questionID"`
	PageID     string `json:"pageID"`
	TestID     string `json:"testID"`
	UserID     string `json:"userID"`
}
