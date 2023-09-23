CREATE TABLE t_states (
	id VARCHAR(8) PRIMARY KEY,
  "name" VARCHAR(50) NOT NULL,
  uf VARCHAR(2) NOT NULL,
	latitude NUMERIC NOT NULL,
	longitude NUMERIC NOT NULL
);

CREATE TABLE t_cities (
	id VARCHAR(8) PRIMARY KEY,
	state_id VARCHAR(8) NOT NULL REFERENCES t_states (id),
  city VARCHAR(100) NOT NULL,
  slug VARCHAR(100) NOT NULL,
	abbreviation VARCHAR(4) NOT NULL,
	border_towns_id TEXT[],
	latitude NUMERIC NOT NULL,
	longitude NUMERIC NOT NULL
);

CREATE TYPE article_status AS ENUM ('Draft', 'Pending', 'Private', 'Publish', 'Trash');

CREATE TABLE t_article (
	id VARCHAR(8) PRIMARY KEY,
  content TEXT,
	"status" article_status DEFAULT 'Draft',
	city_id VARCHAR(8) NOT NULL REFERENCES t_cities (id),
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ
);

CREATE TABLE t_article_props (
	id VARCHAR(8) PRIMARY KEY,
	article_id VARCHAR(8) NOT NULL REFERENCES t_article (id),
	"name" VARCHAR(30) NOT NULL,
	"value" VARCHAR(150) NOT NULL,
	sortkey INTEGER NOT NULL
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

CREATE TYPE municipal_region_type AS ENUM ('Neighborhood', 'Village', 'District');

CREATE TABLE t_municipal_regions (
	id VARCHAR(8) PRIMARY KEY,
	city_id VARCHAR(8) NOT NULL REFERENCES t_cities (id),
	"name" VARCHAR(150) NOT NULL,
	region municipal_region_type NOT NULL DEFAULT 'Neighborhood',
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ
);

CREATE TABLE t_municipal_regions_props (
	id VARCHAR(8) PRIMARY KEY,
	municipal_regions_id VARCHAR(8) NOT NULL REFERENCES t_municipal_regions (id),
	"name" VARCHAR(30) NOT NULL,
	"value" VARCHAR(150) NOT NULL,
	sortkey INTEGER NOT NULL
);

CREATE TABLE t_address (
	id VARCHAR(8) PRIMARY KEY,
	city_id VARCHAR(8) NOT NULL REFERENCES t_cities (id),
	municipal_regions_id VARCHAR(8) NOT NULL REFERENCES t_municipal_regions (id),
	"name" VARCHAR(150) NOT NULL,
	"number" VARCHAR,
	zip_code INTEGER,
	latitude NUMERIC NOT NULL,
	longitude NUMERIC NOT NULL
);
