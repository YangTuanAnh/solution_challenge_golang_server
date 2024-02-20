package models

import (
	"time"
)

// User represents the user's information
type User struct {
    Name         string        `json:"name"`
    Point        int           `json:"point"`
    Activities   []Activity    `json:"activities"`
    Achievements []Achievement `json:"achievements"`
}

// Activity represents a user activity
type Activity struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Type  string `json:"type"`
    Date  string `json:"date"`
    Point int    `json:"point"`
}

// Achievement represents a user achievement
type Achievement struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Date  string `json:"date"`
}