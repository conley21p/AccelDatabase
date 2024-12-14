create table users(
  id text not null primary key default nanoid(),
  username text UNIQUE,
  password text NOT NULL,
  created_at timestamptz not null default now(),
  updated_at timestamptz
);

create trigger users_updated_at
  before update on users
  for each row
  execute procedure moddatetime (updated_at);
