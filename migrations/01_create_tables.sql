-- Users table for authentication
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Customers table
CREATE TABLE customers (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    profile_image VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Bank accounts table
CREATE TABLE bank_accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id VARCHAR(50) NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    account_number VARCHAR(50) UNIQUE NOT NULL,
    account_type VARCHAR(20) NOT NULL,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Pockets table (savings goals)
CREATE TABLE pockets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id VARCHAR(50) NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    goal DECIMAL(15, 2),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Term deposits table
CREATE TABLE term_deposits (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_id VARCHAR(50) NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
    amount DECIMAL(15, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    interest_rate DECIMAL(5, 2) NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    maturity_date TIMESTAMP WITH TIME ZONE NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

-- Add indexes
CREATE INDEX idx_customers_name ON customers(name);
CREATE INDEX idx_customers_email ON customers(email);
CREATE INDEX idx_bank_accounts_customer_id ON bank_accounts(customer_id);
CREATE INDEX idx_pockets_customer_id ON pockets(customer_id);
CREATE INDEX idx_term_deposits_customer_id ON term_deposits(customer_id);

-- Insert sample user for testing
INSERT INTO users (username, password, role) VALUES 
('admin', '$2a$10$fRG7j4oGzb4D3KK8HGXt0uT8AUEagh0Dl3ymexo/yfjh7W2/4esG2', 'admin'); -- password: password

-- Insert sample customers
INSERT INTO customers (id, name, email, phone, address, profile_image) VALUES
('C001', 'John Doe', 'john.doe@example.com', '+1-555-123-4567', '123 Main St, Anytown, USA', '/images/avatars/john.jpg'),
('C002', 'Jane Smith', 'jane.smith@example.com', '+1-555-987-6543', '456 Oak Ave, Somewhere, USA', '/images/avatars/jane.jpg'),
('C003', 'Alice Johnson', 'alice.johnson@example.com', '+1-555-555-5555', '789 Pine Rd, Elsewhere, USA', NULL);

-- Insert sample bank accounts
INSERT INTO bank_accounts (customer_id, account_number, account_type, balance, currency) VALUES
('C001', '1234567890', 'Checking', 5280.42, 'USD'),
('C001', '0987654321', 'Savings', 12750.89, 'USD'),
('C002', '1122334455', 'Checking', 3690.21, 'USD'),
('C003', '5566778899', 'Savings', 28456.32, 'USD'),
('C003', '9988776655', 'Checking', 1205.67, 'USD');

-- Insert sample pockets
INSERT INTO pockets (customer_id, name, balance, currency, goal) VALUES
('C001', 'Vacation Fund', 2500.00, 'USD', 5000.00),
('C001', 'Emergency Fund', 7500.00, 'USD', 10000.00),
('C002', 'New Car', 12000.00, 'USD', 25000.00),
('C003', 'Home Renovation', 8750.00, 'USD', 15000.00);

-- Insert sample term deposits
INSERT INTO term_deposits (customer_id, amount, currency, interest_rate, start_date, maturity_date, is_active) VALUES
('C001', 10000.00, 'USD', 3.25, '2024-01-01', '2025-01-01', TRUE),
('C002', 25000.00, 'USD', 3.50, '2024-02-15', '2025-02-15', TRUE),
('C003', 15000.00, 'USD', 3.75, '2023-06-30', '2024-06-30', TRUE),
('C001', 5000.00, 'USD', 2.75, '2023-03-15', '2024-03-15', FALSE);