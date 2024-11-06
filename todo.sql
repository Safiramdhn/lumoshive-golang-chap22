CREATE TYPE user_status_enum AS ENUM (
	'active',
	'inactive'
)

CREATE TYPE todo_status_enum AS ENUM (
	'not_started',
	'done',
	'deleted'
)

CREATE TABLE users (
	id serial primary key not null,
	name varchar not null,
	email varchar not null unique,
	password varchar not null,
	token varchar,
	user_status user_status_enum default 'active'
)

SELECT * FROM users

CREATE TABLE todos (
	id serial primary key not null,
	description varchar not null,
	user_id int not null,
	todo_status todo_status_enum default 'not_started'
)

SELECT * FROM todos