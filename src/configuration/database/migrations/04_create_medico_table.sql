CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS medico CASCADE;

CREATE TABLE medico (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nomecompleto VARCHAR(100) NOT NULL,
    cpf CHAR(14) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    telefone VARCHAR(20) NOT NULL,
    crm VARCHAR(15) NOT NULL,
    especialidade VARCHAR(50),
    senha VARCHAR(255) NOT NULL,
    unidade_saude_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_medico_id ON medico(id);