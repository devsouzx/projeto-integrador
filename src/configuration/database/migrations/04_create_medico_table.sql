CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS medico;

CREATE TABLE medico (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nomecompleto VARCHAR(100) NOT NULL,
    cpf CHAR(14) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    telefone VARCHAR(20) NOT NULL,
    crm VARCHAR(15) NOT NULL,
    especialidade VARCHAR(50),
    senha VARCHAR(255) NOT NULL,
    unidade_saude_id UUID,

    CONSTRAINT fk_medico_unidade
        FOREIGN KEY (unidade_saude_id)
        REFERENCES unidade_saude(id)
        ON DELETE SET NULL
);
