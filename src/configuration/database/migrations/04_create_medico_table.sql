CREATE TABLE IF NOT EXISTS medico (
    id UUID PRIMARY KEY,
    nomecompleto VARCHAR(100) NOT NULL,
    cpf CHAR(14) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL,
    telefone VARCHAR(20) NOT NULL,
    crm VARCHAR(15) NOT NULL,
    especialidade VARCHAR(50),
    senha VARCHAR(100) NOT NULL,
    unidade_saude_id UUID,
    CONSTRAINT uq_medico_email UNIQUE (email),
    FOREIGN KEY (unidade_saude_id) REFERENCES unidade_saude(id)
);