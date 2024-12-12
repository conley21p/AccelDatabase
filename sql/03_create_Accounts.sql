
CREATE TABLE account (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) REFERENCES user(id),
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE,
    phone_number VARCHAR(20)
);

CREATE TABLE driver (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) REFERENCES user(id),
    acct_id VARCHAR(36) REFERENCES account(id),
    years_of_experience INTEGER,
    policy_number VARCHAR(50),
    insProvider VARCHAR(50),
    policy_start_date DATE,
    policy_end_date DATE,
    license_number VARCHAR(50),
    license_expire_date DATE,
    rating_id VARCHAR(36)
);

CREATE TABLE buyer (
    user_id VARCHAR(36) PRIMARY KEY REFERENCES users(id),
    acct_id VARCHAR(36) REFERENCES account(id),
    transaction_id VARCHAR(36),
    prior_deliveries INTEGER
);

CREATE TABLE ratings (
    user_id VARCHAR(36) PRIMARY KEY REFERENCES users(id),
    past_deliveries INTEGER,
    average_rating DECIMAL(2,1)
);


-- Insert initial data (adjust as needed)
INSERT INTO users (id, username, password) VALUES
    ('user1', 'Driver1', 'password123'),
    ('user2', 'Driver2', 'password123');
    ('user1', 'buyer1', 'password123'),
    ('user2', 'buyer2', 'password123');

INSERT INTO driver (id, user_id, years_of_experience, policy_number, insProvider, policy_start_date, policy_end_date, license_number, license_expire_date, rating_id)
VALUES
('driver1', 'user1', 5, 'PL12345', 'InsuranceCo', '2023-01-01', '2024-12-31', 'DL1234567', '2025-12-31', 'rating1'),
('driver2', 'user2', 3, 'PL67890', 'AnotherIns', '2022-04-05', '2024-03-31', 'DL9876543', '2024-09-30', 'rating2');


INSERT INTO buyers (user_id, transaction_id, prior_deliveries)
VALUES
('user1', 'transaction1', 10),
('user2', 'transaction2', 5);
