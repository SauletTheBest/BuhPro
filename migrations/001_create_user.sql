-- Включение расширения для генерации UUID
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Таблица для обычных пользователей
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Таблица для refresh токенов
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    token TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Таблица для клиентов (Customer)
CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_type TEXT NOT NULL,
    company_name TEXT NOT NULL,
    iin DOUBLE PRECISION NOT NULL,
    name TEXT NOT NULL,
    job_position TEXT NOT NULL,
    phone_number DOUBLE PRECISION NOT NULL,
    email TEXT UNIQUE NOT NULL,
    address TEXT NOT NULL,
    work_description TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Таблица для коучей (Coach)
CREATE TABLE IF NOT EXISTS coaches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    phone_number DOUBLE PRECISION NOT NULL,
    email TEXT UNIQUE NOT NULL,
    exp_coach TEXT NOT NULL,
    specializations TEXT NOT NULL,
    education_certificates TEXT NOT NULL,
    achievements_experience TEXT NOT NULL,
    methodology TEXT NOT NULL,
    about_coach TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Таблица для исполнителей (Executor)
CREATE TABLE IF NOT EXISTS executors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    patronymic TEXT NOT NULL,
    iin DOUBLE PRECISION NOT NULL,
    phone_number DOUBLE PRECISION NOT NULL,
    email TEXT UNIQUE NOT NULL,
    city TEXT NOT NULL,
    exp_work TEXT NOT NULL,
    specializations TEXT NOT NULL,
    education TEXT NOT NULL,
    work_format TEXT NOT NULL,
    hourly_rate DOUBLE PRECISION NOT NULL,
    about_executor TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Таблица для заказов (Order) - пустая для будущего
CREAT   E TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- TODO: Добавить поля для заказов
    created_at TIMESTAMP DEFAULT NOW()
);

-- Таблица для ответов/предложений (Response) - пустая для будущего
CREATE TABLE IF NOT EXISTS responses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- TODO: Добавить поля для ответов
    created_at TIMESTAMP DEFAULT NOW()
);

-- Таблица для рейтингов (Rating) - пустая для будущего
CREATE TABLE IF NOT EXISTS ratings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- TODO: Добавить поля для рейтингов
    created_at TIMESTAMP DEFAULT NOW()
);

-- Таблица для курсов (Course) - пустая для будущего
CREATE TABLE IF NOT EXISTS courses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- TODO: Добавить поля для курсов
    created_at TIMESTAMP DEFAULT NOW()
);

-- Таблица для платежей (Payment) - пустая для будущего
CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- TODO: Добавить поля для платежей
    created_at TIMESTAMP DEFAULT NOW()
);