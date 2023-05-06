CREATE DATABASE IF NOT EXISTS "nossobr";
USE "nossobr";

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE t_cities (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  city VARCHAR(100) NOT NULL,
  "state" VARCHAR(2) NOT NULL
);
