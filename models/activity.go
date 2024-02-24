package models

// Activity represents a user activity
type Activity struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Type  string `json:"type"`
    Date  string `json:"date"`
    Point int    `json:"point"`
}
