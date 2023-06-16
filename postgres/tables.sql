CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
	id uuid DEFAULT uuid_generate_v4 (),
	name VARCHAR NOT NULL,
	email VARCHAR UNIQUE NOT NULL,
	password VARCHAR NOT NULL,
	role VARCHAR NOT NULL,
	status bool NOT NULL,
	avatar_url VARCHAR NULL,
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
	classroom_id uuid NULL REFERENCES classrooms (id),
	status bool NOT NULL DEFAULT true,
	address VARCHAR NULL,
	city VARCHAR NULL,
	cep VARCHAR NULL,
	fones VARCHAR NULL,
	cpf VARCHAR UNIQUE NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS parents (
	id uuid DEFAULT uuid_generate_v4 (),
	name VARCHAR NOT NULL,
	cpf VARCHAR UNIQUE NULL,
	email VARCHAR NULL,
	fones VARCHAR NULL,
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS parents_students (
	parent_id uuid REFERENCES parents (id) ON UPDATE CASCADE ON DELETE SET NULL,
	student_id uuid REFERENCES students (id) ON UPDATE CASCADE ON DELETE SET NULL,
	relationship VARCHAR NOT NULL,
	responsible bool NOT NULL DEFAULT false,
	CONSTRAINT parents_students_pk PRIMARY KEY (parent_id, student_id) -- explicit pk
);

CREATE TABLE IF NOT EXISTS teachers (
	id uuid DEFAULT uuid_generate_v4 (),
	name VARCHAR NOT NULL,
	nick VARCHAR NOT NULL,
	birth_day TIMESTAMP NULL,
	gender VARCHAR NULL,
	cpf VARCHAR UNIQUE NULL,
	fones VARCHAR NULL,
	email VARCHAR NULL,
	license VARCHAR NOT NULL,
	note VARCHAR NULL,
	status bool NOT NULL DEFAULT true,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);