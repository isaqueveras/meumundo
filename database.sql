CREATE DATABASE IF NOT EXISTS "nossobr";
USE "nossobr";

CREATE TABLE t_estados (
	id VARCHAR(8) PRIMARY KEY,
  nome VARCHAR(50) NOT NULL,
  sigla VARCHAR(2) NOT NULL
);

CREATE TABLE t_cidades (
	id VARCHAR(8) PRIMARY KEY,
	estado_id VARCHAR(8) NOT NULL REFERENCES t_estados (id),
  city VARCHAR(100) NOT NULL,
  slug VARCHAR(100) NOT NULL
);

CREATE TABLE t_artigos (
	id VARCHAR(8) PRIMARY KEY,
  conteudo TEXT,
  info JSON NOT NULL,
	cidade_id VARCHAR(8) NOT NULL REFERENCES t_cidades (id),
	criacao TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	atualizacao TIMESTAMPTZ
);
