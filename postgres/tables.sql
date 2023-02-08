CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
	id uuid DEFAULT uuid_generate_v4 (),
	name VARCHAR NOT NULL,
	email VARCHAR UNIQUE NOT NULL,
	password VARCHAR NOT NULL,
	role VARCHAR NOT NULL,
	status bool NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS refresh_tokens (
	id SERIAL PRIMARY KEY,
	user_id uuid NOT NULL REFERENCES users (id) ,
	token VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	expired_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS classrooms (
	id uuid DEFAULT uuid_generate_v4 (),
	name VARCHAR NOT NULL,
	level VARCHAR NOT NULL,
	grade VARCHAR NOT NULL,
	shift VARCHAR NOT NULL,
	description VARCHAR NULL,
	anne VARCHAR NULL,
	year VARCHAR NOT NULL,
	status bool NOT NULL DEFAULT true,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS students (
	id uuid DEFAULT uuid_generate_v4 (),
	name VARCHAR NOT NULL,
	birth_day TIMESTAMP NOT NULL,
	gender VARCHAR NULL,
	anne VARCHAR NULL,
	note VARCHAR NULL,
	ieducar INTEGER UNIQUE NOT NULL,
	educa_df VARCHAR NULL,
	classroom_id NULL REFERENCES classrooms (id),
	status bool NOT NULL DEFAULT true,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);