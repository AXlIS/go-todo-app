CREATE SCHEMA IF NOT EXISTS content;

CREATE TABLE IF NOT EXISTS content.users
(
    id            serial primary key,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS content.todo_lists
(
    id          serial primary key,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE IF NOT EXISTS content.users_lists
(
    id      serial primary key,
    user_id int references content.users (id) on delete cascade      not null,
    list_id int references content.todo_lists (id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS content.todo_items
(
    id          serial primary key,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);

CREATE TABLE IF NOT EXISTS content.lists_items
(
    id      serial primary key,
    item_id int references content.todo_items (id) on delete cascade not null,
    list_id int references content.todo_lists (id) on delete cascade not null
);