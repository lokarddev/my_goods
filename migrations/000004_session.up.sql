create table refresh_session (
  id serial primary key ,
  user_id int references Users on delete cascade,
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  refresh_token uuid,
  user_agent varchar(200),
  fingerprint varchar(200),
  ip varchar(15),
  expires_in bigint
);
