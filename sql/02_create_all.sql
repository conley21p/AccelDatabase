create table users(
  id text not null primary key default nanoid(),
  username text UNIQUE,
  password text NOT NULL,
  created_at timestamptz not null default now(),
  updated_at timestamptz
);

-- create trigger users_updated_at
--   before update on users
--   for each row
--   execute procedure moddatetime (updated_at);


-- Create the Driver table which inherits from Account
CREATE TABLE drivers (
    id text not null primary key default nanoid(),
    user_id VARCHAR(36) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    policy_number VARCHAR(255) NOT NULL,
    ins_provider VARCHAR(255) NOT NULL,
    policy_start_date TIMESTAMP WITH TIME ZONE NOT NULL,
    policy_end_date TIMESTAMP WITH TIME ZONE NOT NULL,
    license_number VARCHAR(255) NOT NULL,
    license_expire_date TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
)

-- create trigger drivers_updated_at
--   before update on drivers
--   for each row
--   execute procedure moddatetime (updated_at);

-- Create the Buyer table with a foreign key relationship to Account
CREATE TABLE buyers (
    id text not null primary key default nanoid(),
    user_id VARCHAR(36) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    transactionId VARCHAR(255),
    priorDeliveries BIGINT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
)

ALTER TABLE users ADD COLUMN driver_id VARCHAR(36) REFERENCES drivers(id);

ALTER TABLE users ADD COLUMN buyer_id VARCHAR(36) REFERENCES buyers(id);

CREATE TABLE ratings (
    id text not null primary key default nanoid(),
    driver_id VARCHAR(36),
    past_deliveries INTEGER,
    average_rating DECIMAL(2,1),
    created_at timestamptz not null default now(),
    updated_at timestamptz,
    FOREIGN KEY (driver_id) REFERENCES drivers(id)
);


create trigger ratings_updated_at
  before update on ratings
  for each row
  execute procedure moddatetime (updated_at);