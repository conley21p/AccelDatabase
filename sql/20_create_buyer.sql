-- Create the Buyer table with a foreign key relationship to Account
CREATE TABLE buyers (
    id text not null primary key default nanoid(),
    user_id text not null REFERENCES users(id) on delete cascade on update cascade,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    -- transactionId VARCHAR(255),
    priorDeliveries BIGINT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE
);

create trigger buyers_updated_at
  before update on buyers
  for each row
  execute procedure moddatetime (updated_at);
