package models

// User represents the user's information
type User struct {
	ID			 string		   `json:"id"`
    Name         string        `json:"name"`
    Point        int           `json:"point"`
    Activities   []Activity    `json:"activities"`
    Achievements []Achievement `json:"achievements"`
}


