package data

// Test defines a test paper structure.
type Test struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	MarkNum   int    `json:"markNum"`
	MarkDenom int    `json:"markDenom"`

	UserID int `json:"userID"`
}
