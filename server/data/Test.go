package data

// Test defines a test paper structure.
type Test struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	MarkNum   int8   `json:"markNum"`
	MarkDenom int8   `json:"markDenom"`

	UserID string `json:"userID"`
}
