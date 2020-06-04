package entities

// Base Format for Post Data Structure
type Post struct {
	ID    int64  `json:"ID"`
	Title string `json:"Title"`
	Text  string `json:"Text"`
}
