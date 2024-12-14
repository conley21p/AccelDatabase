-- Create the Driver table which inherits from Account
CREATE TABLE drivers (
    id text not null primary key default nanoid(),
    user_id text not null REFERENCES users(id) on delete cascade on update cascade,
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
    updated_at TIMESTAMP WITH TIME ZONE
);

create trigger drivers_updated_at
  before update on drivers
  for each row
  execute procedure moddatetime (updated_at);
  