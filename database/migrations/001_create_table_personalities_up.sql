create table IF NOT EXISTS personalities(
    id serial primary key not null,
    name varchar(200) not null,
    history varchar(400)
);