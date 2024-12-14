
CREATE TABLE ratings (
    id text not null primary key default nanoid(),
    driver_id text not null REFERENCES drivers(id) on delete cascade on update cascade,
    past_deliveries INTEGER,
    average_rating DECIMAL(2,1),
    created_at timestamptz not null default now(),
    updated_at timestamptz
);

create trigger ratings_updated_at
  before update on ratings
  for each row
  execute procedure moddatetime (updated_at);