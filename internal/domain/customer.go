package domain

import "time"

type Customer struct {
	ID         string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ClientType string `gorm:"not null"` //ТОО, ИП, Доверенный представитель

	CompanyName string  `gorm:"not null"`
	IIN         float64 `gorm:"not null"`
	Name        string  `gorm:"not null"`
	JobPosition string  `gorm:"not null"`

	PhoneNumber     float64   `gorm:"not null"`
	Email           string    `gorm:"unique;not null"`
	Address         string    `gorm:"not null"`
	WorkDescription string    `gorm:"not null"`
	PasswordHash    string    `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
