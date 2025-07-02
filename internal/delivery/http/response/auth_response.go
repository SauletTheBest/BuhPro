package responses

// AuthSuccessResponse представляет успешный ответ на регистрацию/логин.
type AuthSuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// LoginResponse представляет ответ на успешный логин.
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// TokenRefreshResponse представляет ответ на успешное обновление токена.
type TokenRefreshResponse struct {
	AccessToken string `json:"access_token"`
}

// ErrorResponse представляет стандартизированный ответ для ошибок.
type ErrorResponse struct {
	Error   string      `json:"error"`
	Details interface{} `json:"details,omitempty"`
}

// UserProfileResponse представляет информацию профиля обычного пользователя.
type UserProfileResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// CustomerProfileResponse представляет информацию профиля клиента.
type CustomerProfileResponse struct {
	ID              string  `json:"id"`
	ClientType      string  `json:"client_type"`
	CompanyName     string  `json:"company_name"`
	IIN             float64 `json:"iin"`
	Name            string  `json:"name"`
	JobPosition     string  `json:"job_position"`
	PhoneNumber     float64 `json:"phone_number"`
	Email           string  `json:"email"`
	Address         string  `json:"address"`
	WorkDescription string  `json:"work_description"`
}

// CoachProfileResponse представляет информацию профиля коуча.
type CoachProfileResponse struct {
	ID                     string  `json:"id"`
	Name                   string  `json:"name"`
	Surname                string  `json:"surname"`
	PhoneNumber            float64 `json:"phone_number"`
	Email                  string  `json:"email"`
	ExpCoach               string  `json:"exp_coach"`
	Specializations        string  `json:"specializations"`
	EducationCertificates  string  `json:"education_certificates"`
	AchievementsExperience string  `json:"achievements_experience"`
	Methodology            string  `json:"methodology"`
	AboutCoach             string  `json:"about_coach"`
}

// ExecutorProfileResponse представляет информацию профиля исполнителя.
type ExecutorProfileResponse struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Surname         string  `json:"surname"`
	Patronymic      string  `json:"patronymic"`
	IIN             float64 `json:"iin"`
	PhoneNumber     float64 `json:"phone_number"`
	Email           string  `json:"email"`
	City            string  `json:"city"`
	ExpWork         string  `json:"exp_work"`
	Specializations string  `json:"specializations"`
	Education       string  `json:"education"`
	WorkFormat      string  `json:"work_format"`
	HourlyRate      float64 `json:"hourly_rate"`
	AboutExecutor   string  `json:"about_executor"`
}
