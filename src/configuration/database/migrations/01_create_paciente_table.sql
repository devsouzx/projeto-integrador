CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS paciente CASCADE;

CREATE TABLE paciente (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cpf CHAR(11) NOT NULL UNIQUE,
    nome VARCHAR(100) NOT NULL,
    cns VARCHAR(15) NOT NULL UNIQUE,
    nome_mae VARCHAR(100),
    data_nascimento DATE NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    telefone VARCHAR(20) NOT NULL,
    senha VARCHAR(255) NOT NULL,
    apelido VARCHAR(50),
    raca VARCHAR(30),
    nacionalidade VARCHAR(50) DEFAULT 'Brasileira',
    escolaridade VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_verified BOOLEAN DEFAULT FALSE,
    verification_token VARCHAR(255),
    token_expires_at TIMESTAMP
);
