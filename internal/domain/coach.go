package domain

import "time"

type Coach struct {
	ID      string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name    string `gorm:"not null"`
	Surname string `gorm:"not null"`

	PhoneNumber float64 `gorm:"not null"`
	Email       string  `gorm:"unique;not null"`

	ExpCoach        string `gorm:"not null"` //1-2 года, 3-5 лет, 6-10 лет, Более 10 лет
	Specializations string `gorm:"not null"` // Бизнес-коучинг, Карьерный коучинг, Финансовый коучинг, Лидерство, Личностный рост

	EducationCertificates  string    `gorm:"not null"`
	AchievementsExperience string    `gorm:"not null"`
	Methodology            string    `gorm:"not null"`
	AboutCoach             string    `gorm:"not null"`
	PasswordHash           string    `gorm:"not null"`
	CreatedAt              time.Time `gorm:"autoCreateTime"`
}
