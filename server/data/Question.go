package data

// Question defines a test paper page structure.
type Question struct {
	ID              string `json:"id"`
	ExpectedAnswers string `json:"expectedAnswers"`
	MarkNum         int8   `json:"markNum"`
	MarkDenom       int8   `json:"markDenom"`

	PageID string `json:"pageID"`
	TestID string `json:"testID"`
	UserID string `json:"userID"`
}
