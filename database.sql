CREATE DATABASE IF NOT EXISTS "nossobr";
USE "nossobr";

CREATE TABLE t_states (
	id VARCHAR(8) PRIMARY KEY,
  "name" VARCHAR(50) NOT NULL,
  uf VARCHAR(2) NOT NULL
);

CREATE TABLE t_cities (
	id VARCHAR(8) PRIMARY KEY,
	state_id VARCHAR(8) NOT NULL REFERENCES t_states (id),
  city VARCHAR(100) NOT NULL,
  slug VARCHAR(100) NOT NULL,
	border_towns TEXT[]
);

CREATE TABLE t_article (
	id VARCHAR(8) PRIMARY KEY,
  content TEXT,
	city_id VARCHAR(8) NOT NULL REFERENCES t_cities (id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ
);

CREATE TABLE t_children (
	id VARCHAR(8) PRIMARY KEY,
	city_id VARCHAR(8) NOT NULL REFERENCES t_cities (id),
	"name" VARCHAR(100) NOT NULL,
	"url" VARCHAR,
	short_desc VARCHAR(100) NOT NULL,
	biography TEXT,
	date_birth DATE,
	date_death DATE,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ
);