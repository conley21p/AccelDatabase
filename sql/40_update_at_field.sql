-- Function to update timestamp
CREATE OR REPLACE FUNCTION moddatetime(updated_at_column TIMESTAMP WITH TIME ZONE) RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Users table trigger
CREATE TRIGGER users_updated_at
  BEFORE UPDATE ON users
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- ContactInfo table trigger
CREATE TRIGGER contact_info_updated_at
  BEFORE UPDATE ON contact_info
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Drivers table trigger
CREATE TRIGGER drivers_updated_at
  BEFORE UPDATE ON drivers
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Insurance table trigger
CREATE TRIGGER insurance_updated_at
  BEFORE UPDATE ON insurance
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- License table trigger
CREATE TRIGGER license_updated_at
  BEFORE UPDATE ON license
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Haulers table trigger
CREATE TRIGGER haulers_updated_at
  BEFORE UPDATE ON haulers
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Trailers table trigger
CREATE TRIGGER trailers_updated_at
  BEFORE UPDATE ON trailers
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Offers table trigger
CREATE TRIGGER offers_updated_at
  BEFORE UPDATE ON offers
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Owners table trigger
CREATE TRIGGER owners_updated_at
  BEFORE UPDATE ON owners
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Transportation table trigger
CREATE TRIGGER transportation_updated_at
  BEFORE UPDATE ON transportation
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Ratings table trigger
CREATE TRIGGER ratings_updated_at
  BEFORE UPDATE ON ratings
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Transactions table trigger
CREATE TRIGGER transactions_updated_at
  BEFORE UPDATE ON transactions
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Vehicles table trigger
CREATE TRIGGER vehicles_updated_at
  BEFORE UPDATE ON vehicles
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Autos table trigger
CREATE TRIGGER autos_updated_at
  BEFORE UPDATE ON autos
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Boats table trigger
CREATE TRIGGER boats_updated_at
  BEFORE UPDATE ON boats
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-- Conversations table trigger
-- Note: Conversations table does not have an updated_at field, so no trigger is created for it.