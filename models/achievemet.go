package models

// Achievement represents a user achievement
type Achievement struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Date  string `json:"date"`
}