CREATE TABLE IF NOT EXISTS agente_comunitario (
    id UUID PRIMARY KEY,
    nomecompleto VARCHAR(100) NOT NULL,
    cpf CHAR(14) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL,
    telefone VARCHAR(20) NOT NULL,
    email_institucional VARCHAR(100) NOT NULL,
    senha VARCHAR(100) NOT NULL,
    unidade_saude_id UUID,
    CONSTRAINT uq_agente_email UNIQUE (email),
    FOREIGN KEY (unidade_saude_id) REFERENCES unidade_saude(id)
);