package data

// Page defines a test paper page structure.
type keyword struct {
	ID         string `json:"id"`
	QuestionID string `json:"questionID"`
	Word       string `json:"string"`

	PageID string `json:"pageID"`
	TestID string `json:"testID"`
	UserID string `json:"userID"`
}
