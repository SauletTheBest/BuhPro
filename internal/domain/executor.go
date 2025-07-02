package domain

import "time"

type Executor struct {
	ID          string  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string  `gorm:"not null"`
	Surname     string  `gorm:"not null"`
	Patronymic  string  `gorm:"not null"`
	IIN         float64 `gorm:"not null"`
	PhoneNumber float64 `gorm:"not null"`
	Email       string  `gorm:"unique;not null"`
	City        string  `gorm:"not null"`
	ExpWork     string  `gorm:"not null"`

	Specializations string  `gorm:"not null"` // Бухгалтерский учет, Налоговое консультирование, Аудиторские услуги, Финансовый анализ, Подготовка отчетности, Восстановление учета, Управленческий учет, Международные стандарты (МСФО), Налоговое планирование, Кадровое делопроизводство
	Education       string  `gorm:"not null"`
	WorkFormat      string  `gorm:"not null"` //Удаленно, В офисе клиента, Смешанный формат, Гибкий график
	HourlyRate      float64 `gorm:"not null"`
	AboutExecutor   string  `gorm:"not null"`

	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
