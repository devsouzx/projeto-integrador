CREATE TABLE IF NOT EXISTS endereco (
    id UUID PRIMARY KEY,
    paciente_id UUID NOT NULL,
    cep CHAR(9),
    logradouro VARCHAR(100),
    numero VARCHAR(10),
    complemento VARCHAR(50),
    bairro VARCHAR(50),
    cidade VARCHAR(50),
    uf CHAR(2),
    FOREIGN KEY (paciente_id) REFERENCES paciente(paciente_id)
);