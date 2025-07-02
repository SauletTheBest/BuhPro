package models

import "time"

type Order struct {
	ID string 
	Title string
	Description string 
	Deadline time.Time 
	Category string
	Region string 
	Status string // active, in progress , closed
	ClientID string 
	CreatedAt time.Time 
}