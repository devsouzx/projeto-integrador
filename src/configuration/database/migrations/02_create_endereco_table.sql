CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS endereco;

CREATE TABLE endereco (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cep CHAR(9) NOT NULL,
    logradouro VARCHAR(100) NOT NULL,
    numero VARCHAR(10) NOT NULL,
    complemento VARCHAR(50),
    bairro VARCHAR(50) NOT NULL,
    cidade VARCHAR(50) NOT NULL,
    uf CHAR(2) NOT NULL,
    referencia VARCHAR(100),
    codigo_municipio VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    paciente_id UUID NOT NULL,
    CONSTRAINT fk_endereco_paciente
        FOREIGN KEY (paciente_id)
        REFERENCES paciente(id)
        ON DELETE CASCADE
);