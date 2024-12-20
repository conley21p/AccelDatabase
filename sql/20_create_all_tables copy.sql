-- User table
CREATE TABLE users (
    id text not null primary key default nanoid(),
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- ContactInfo table
CREATE TABLE contact_info (
    id text not null primary key default nanoid(),
    phone_number VARCHAR(20),
    street_address VARCHAR(255),
    city VARCHAR(100),
    state VARCHAR(50),
    zip_code VARCHAR(20),
    country VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Driver table
CREATE TABLE drivers (
    id text not null primary key default nanoid(),
    user_id VARCHAR(255) REFERENCES users(id),
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Insurance table
CREATE TABLE insurance (
    id text not null primary key default nanoid(),
    policy_number VARCHAR(50),
    ins_provider VARCHAR(100),
    policy_start_date DATE,
    policy_end_date DATE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- License table
CREATE TABLE license (
    id text not null primary key default nanoid(),
    license_number VARCHAR(50),
    license_expire_date DATE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Hauler table
CREATE TABLE haulers (
    id text not null primary key default nanoid(),
    driver_id text not null REFERENCES drivers(id),
    make VARCHAR(50),
    model VARCHAR(50),
    year INTEGER,
    mileage DOUBLE PRECISION,
    towing_capacity DOUBLE PRECISION,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Trailer table
CREATE TABLE trailers (
    id text not null primary key default nanoid(),
    type VARCHAR(50),
    length DOUBLE PRECISION,
    width DOUBLE PRECISION,
    capacity DOUBLE PRECISION,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Hauler-Trailer join table
CREATE TABLE hauler_trailers (
    hauler_id text REFERENCES haulers(id),
    trailer_id text REFERENCES trailers(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (hauler_id, trailer_id)
);

-- Offer table
CREATE TABLE offers (
    id text not null primary key default nanoid(),
    amount DOUBLE PRECISION,
    deadline_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Owner table
CREATE TABLE owners (
    id text not null primary key default nanoid(),
    user_id VARCHAR(255) REFERENCES users(id),
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Transportation table
CREATE TABLE transportation (
    id text not null primary key default nanoid(),
    description TEXT,
    transport_date TIMESTAMP WITH TIME ZONE,
    pickup_address TEXT,
    delivery_address TEXT,
    deliver_by_date TIMESTAMP WITH TIME ZONE,
    pickup_by_date TIMESTAMP WITH TIME ZONE,
    pickup_available_date TIMESTAMP WITH TIME ZONE,
    request_price DOUBLE PRECISION,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Rating table
CREATE TABLE ratings (
    id text not null primary key default nanoid(),
    past_deliveries TEXT,
    average_rating VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Transaction table
CREATE TABLE transactions (
    id text not null primary key default nanoid(),
    payment_method VARCHAR(50),
    amount DOUBLE PRECISION,
    transaction_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Vehicle table
CREATE TABLE vehicles (
    id text not null primary key default nanoid(),
    length INTEGER,
    width INTEGER,
    height INTEGER,
    auto_id text REFERENCES autos(id),
    boat_id text REFERENCES boats(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Auto table
CREATE TABLE autos (
    id text not null primary key default nanoid(),
    make VARCHAR(50),
    model VARCHAR(50),
    year INTEGER,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Boat table
CREATE TABLE boats (
    id text not null primary key default nanoid(),
    make VARCHAR(50),
    model VARCHAR(50),
    year INTEGER,
    with_trailer BOOLEAN,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Conversation table
CREATE TABLE conversations (
    id text not null primary key default nanoid(),
    sender_id VARCHAR(255),
    recipient_id VARCHAR(255),
    content TEXT,
    timestamp TIMESTAMP WITH TIME ZONE
);

-- Message table
CREATE TABLE messages (
    id text not null primary key default nanoid(),
    subject VARCHAR(255)
);