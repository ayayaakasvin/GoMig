CREATE TABLE IF NOT EXISTS test_users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert some test data
INSERT INTO test_users (name, email) VALUES 
    ('Test User 1', 'test1@example.com'),
    ('Test User 2', 'test2@example.com');