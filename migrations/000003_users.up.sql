create table Users (
    id          serial
        primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    name varchar(100) unique,
    password varchar(200)
);

alter table dishes add column user_id int references Users on delete cascade;
alter table goods add column user_id int references Users on delete cascade;
alter table lists add column user_id int references Users on delete cascade;