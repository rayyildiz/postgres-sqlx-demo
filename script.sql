create table if not exists users
(
    id            serial primary key,
    email         text,
    first_name    text,
    last_name     text,
    register_date text,
    insert_order  int
)