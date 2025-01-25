-- Enable the necessary PostgreSQL extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL
);

-- -- Insert sample users with generated UUIDs
-- INSERT INTO users (name) VALUES ('Alice'), ('Bob');