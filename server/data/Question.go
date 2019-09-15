package data

// Question defines a test paper page structure.
type Question struct {
	ID        int    `json:"id"`
	Question  string `json:"question"`
	Answers   string `json:"answers"`
	MarkNum   int    `json:"markNum"`
	MarkDenom int    `json:"markDenom"`

	TestID int `json:"test"`
	UserID int `json:"user"`
}
