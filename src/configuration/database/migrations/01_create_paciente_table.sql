CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS paciente (
    paciente_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nomecompleto VARCHAR(50) NOT NULL,
    cpf VARCHAR(11) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL,
    senha VARCHAR(100) NOT NULL
);