package entities

import "time"

type OrderEntity struct {
    ID          string    `gorm:"type:uuid;primary_key" json:"-"`
    Title       string    `gorm:"size:255" json:"-"`
    Description string    `json:"-"`
    Deadline    time.Time `json:"-"`
    Category    string    `json:"-"`
    Region      string    `json:"-"`
    Status      string    `json:"-"`
    ClientID    string    `json:"-"`
    CreatedAt   time.Time `json:"-"`
}