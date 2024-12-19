-- Add driver_id to users table
ALTER TABLE users ADD COLUMN driver_id VARCHAR(36) REFERENCES drivers(id);

-- Add owner_id to users table
ALTER TABLE users ADD COLUMN owner_id VARCHAR(36) REFERENCES owners(id);

-- Add contact_info_id to drivers table
ALTER TABLE drivers ADD COLUMN contact_info_id VARCHAR(36) REFERENCES contact_info(id);

-- Add insurance_id to drivers table
ALTER TABLE drivers ADD COLUMN insurance_id VARCHAR(36) REFERENCES insurance(id);

-- Add license_id to drivers table
ALTER TABLE drivers ADD COLUMN license_id VARCHAR(36) REFERENCES license(id);

-- Add rating_id to drivers table
ALTER TABLE drivers ADD COLUMN rating_id VARCHAR(36) REFERENCES ratings(id);

-- Add contact_info_id to owners table
ALTER TABLE owners ADD COLUMN contact_info_id VARCHAR(36) REFERENCES contact_info(id);

-- Add accepted_offer_id to transportation table
ALTER TABLE transportation ADD COLUMN accepted_offer_id VARCHAR(36) REFERENCES offers(id);

-- Add vehicle_id to transportation table
ALTER TABLE transportation ADD COLUMN vehicle_id VARCHAR(36) REFERENCES vehicles(id);

-- Add auto_id to vehicles table (assuming one-to-one with autos if it's not null)
ALTER TABLE vehicles ADD COLUMN auto_id VARCHAR(36) REFERENCES autos(id);

-- Add boat_id to vehicles table (assuming one-to-one with boats if it's not null)
ALTER TABLE vehicles ADD COLUMN boat_id VARCHAR(36) REFERENCES boats(id);