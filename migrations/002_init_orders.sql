CREATE TYPE order_status AS ENUM ('active', 'in_progress', 'closed', 'archived');

CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    deadline TIMESTAMP,
    category VARCHAR(100), -- бухгалтерия, налоги, аудит, финансы
    region VARCHAR(100),  -- онлайн / оффлайн
    status order_status DEFAULT 'active',
    client_id UUID REFERENCES users(id),
    budget DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);