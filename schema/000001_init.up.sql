CREATE TABLE IF NOT EXISTS users
(
    id      serial        not null unique,
    login      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS links
(
    id      serial        not null unique,
    base_url varchar(1024) not null,
    short_url   varchar(10)   not null unique
);