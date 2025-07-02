package utils

import (
	"fmt"
	"regexp" // Импорт regexp должен быть здесь

	"github.com/go-playground/validator/v10"
)

// IsPasswordComplex проверяет сложность пароля.
// Пароль должен содержать как минимум одну заглавную букву, одну строчную букву, одну цифру и один специальный символ.
func IsPasswordComplex(password string) bool {
	if len(password) < 8 {
		return false
	}
	// Используем raw string literals (обратные кавычки) для регулярных выражений
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)
	return hasUpper && hasLower && hasNumber && hasSpecial
}

// CustomValidationErrors преобразует ошибки валидации в читаемый формат.
func CustomValidationErrors(errs validator.ValidationErrors) []string {
	var messages []string
	for _, err := range errs {
		switch err.Tag() {
		case "required":
			messages = append(messages, fmt.Sprintf("%s is required", err.Field()))
		case "email":
			messages = append(messages, fmt.Sprintf("%s must be a valid email address", err.Field()))
		case "min":
			messages = append(messages, fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param()))
		default:
			messages = append(messages, fmt.Sprintf("%s is not valid", err.Field()))
		}
	}
	return messages
}
