INSERT INTO users ( username, password) 
VALUES 
('driver2',                                --username
 'password'                               --password
),
( 
  'jane_smith2',
  'securepass'
);

-- Insert the new driver and return the id
WITH inserted_driver AS (
    INSERT INTO drivers (
        user_id, 
        first_name, 
        last_name, 
        phone_number, 
        policy_number, 
        ins_provider, 
        policy_start_date, 
        policy_end_date, 
        license_number, 
        license_expire_date
    )
    VALUES (
        '111e4567-e89b-12d3-a456-426614174000',  -- user_id
        'Conley',                               -- first_name
        'Price',                                -- last_name
        '+1234567890',                          -- phone_number
        'POL123456',                            -- policy_number
        'ACME Insurance',                       -- ins_provider
        '2024-01-01 00:00:00+00',               -- policy_start_date (UTC)
        '2025-01-01 00:00:00+00',               -- policy_end_date (UTC)
        'LIC67890',                             -- license_number
        '2026-01-01 00:00:00+00'                -- license_expire_date (UTC)
    )
    RETURNING id
)
-- Update the users table with the returned id
UPDATE users 
SET driver_id = (SELECT id FROM inserted_driver)
WHERE id = (SELECT id FROM users WHERE username = 'driver2');