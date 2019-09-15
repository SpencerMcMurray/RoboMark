package data

// Question defines a test paper page structure.
type Question struct {
	ID              int    `json:"id"`
	ExpectedAnswers string `json:"expectedAnswers"`
	MarkNum         int    `json:"markNum"`
	MarkDenom       int    `json:"markDenom"`

	PageID int `json:"pageID"`
	TestID int `json:"testID"`
	UserID int `json:"userID"`
}
