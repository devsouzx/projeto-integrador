CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS paciente CASCADE;

CREATE TABLE paciente (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    cpf CHAR(11) NOT NULL UNIQUE,
    nomecompleto VARCHAR(100) NOT NULL,
    cartaosus CHAR(15) NOT NULL,
    nomecompletodamae VARCHAR(100) NOT NULL,
    datadenascimento DATE NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    telefone VARCHAR(20) NOT NULL,
    senha VARCHAR(255) NOT NULL,
    apelido VARCHAR(30),
    raca VARCHAR(30),
    nacionalidade VARCHAR(50),
    escolaridade VARCHAR(50)
);
