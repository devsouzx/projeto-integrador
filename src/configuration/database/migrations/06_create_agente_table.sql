CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS agente_comunitario;

CREATE TABLE agente_comunitario (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nomecompleto VARCHAR(100) NOT NULL,
    cpf CHAR(11) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    telefone VARCHAR(20) NOT NULL,
    senha VARCHAR(255) NOT NULL,
    unidade_saude_id INTEGER
);

ALTER TABLE agente_comunitario ADD COLUMN created_at TIMESTAMP DEFAULT now();
ALTER TABLE agente_comunitario ADD COLUMN updated_at TIMESTAMP DEFAULT now();