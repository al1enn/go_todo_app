CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);


CREATE TABLE todo_categories
(
    id          serial       not null unique,
    title       varchar(255) not null
);

CREATE TABLE todo_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255) not null,
    is_completed boolean default false,
    is_important boolean default false,
    created_date timestamp default current_timestamp,
    updated_date timestamp default current_timestamp
);



CREATE TABLE user_todo_categories
(
    user_id int not null references users(id) on delete cascade,
    todo_category_id int not null references todo_categories(id) on delete cascade
);



CREATE TABLE todo_item_categories
(
    todo_item_id int not null references todo_items(id) on delete cascade,
    todo_category_id int not null references todo_categories(id) on delete cascade
);

