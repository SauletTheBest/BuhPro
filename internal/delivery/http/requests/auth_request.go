package requests

// AuthRequest представляет общую структуру для запросов на аутентификацию (регистрация/логин).
type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// RefreshRequest представляет структуру для запроса на обновление токена.
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// CustomerRegisterRequest представляет структуру для регистрации клиента.
// Расширяет AuthRequest для включения специфичных полей Customer.
// Если вы хотите, чтобы все поля передавались с фронтенда, раскомментируйте их.
type CustomerRegisterRequest struct {
	AuthRequest
	ClientType      string  `json:"client_type" validate:"required"`
	CompanyName     string  `json:"company_name" validate:"required"`
	IIN             float64 `json:"iin" validate:"required"`
	Name            string  `json:"name" validate:"required"`
	JobPosition     string  `json:"job_position" validate:"required"`
	PhoneNumber     float64 `json:"phone_number" validate:"required"`
	Address         string  `json:"address" validate:"required"`
	WorkDescription string  `json:"work_description" validate:"required"`
}

// CoachRegisterRequest представляет структуру для регистрации коуча.
// Расширяет AuthRequest для включения специфичных полей Coach.
type CoachRegisterRequest struct {
	AuthRequest
	Name                   string  `json:"name" validate:"required"`
	Surname                string  `json:"surname" validate:"required"`
	PhoneNumber            float64 `json:"phone_number" validate:"required"`
	ExpCoach               string  `json:"exp_coach" validate:"required"`
	Specializations        string  `json:"specializations" validate:"required"`
	EducationCertificates  string  `json:"education_certificates" validate:"required"`
	AchievementsExperience string  `json:"achievements_experience" validate:"required"`
	Methodology            string  `json:"methodology" validate:"required"`
	AboutCoach             string  `json:"about_coach" validate:"required"`
}

// ExecutorRegisterRequest представляет структуру для регистрации исполнителя.
// Расширяет AuthRequest для включения специфичных полей Executor.
type ExecutorRegisterRequest struct {
	AuthRequest
	Name            string  `json:"name" validate:"required"`
	Surname         string  `json:"surname" validate:"required"`
	Patronymic      string  `json:"patronymic" validate:"required"`
	IIN             float64 `json:"iin" validate:"required"`
	PhoneNumber     float64 `json:"phone_number" validate:"required"`
	City            string  `json:"city" validate:"required"`
	ExpWork         string  `json:"exp_work" validate:"required"`
	Specializations string  `json:"specializations" validate:"required"`
	Education       string  `json:"education" validate:"required"`
	WorkFormat      string  `json:"work_format" validate:"required"`
	HourlyRate      float64 `json:"hourly_rate" validate:"required"`
	AboutExecutor   string  `json:"about_executor" validate:"required"`
}
