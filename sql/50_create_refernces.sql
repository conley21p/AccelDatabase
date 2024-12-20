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

-- Add auto_id to vehicles table
ALTER TABLE vehicles ADD COLUMN auto_id text REFERENCES autos(id);

-- Add boat_id to vehicles table
ALTER TABLE vehicles ADD COLUMN boat_id text REFERENCES boats(id);

-- Add driver_id to insurance table
ALTER TABLE insurance ADD COLUMN driver_id text REFERENCES drivers(id);

-- Add driver_id to license table
ALTER TABLE license ADD COLUMN driver_id text REFERENCES drivers(id);

-- Add driver_id to haulers table
ALTER TABLE haulers ADD COLUMN driver_id text REFERENCES drivers(id);

-- Add hauler_id to trailers table
ALTER TABLE trailers ADD COLUMN hauler_id text REFERENCES haulers(id);

-- Add driver_id to transportation table
ALTER TABLE transportation ADD COLUMN driver_id text REFERENCES drivers(id);

-- Add driver_id to offers table
ALTER TABLE offers ADD COLUMN driver_id text REFERENCES drivers(id);

-- Add transportation_id to vehicles table
ALTER TABLE vehicles ADD COLUMN transportation_id text REFERENCES transportation(id);
