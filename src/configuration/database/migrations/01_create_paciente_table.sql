CREATE TABLE paciente (
    paciente_id UUID PRIMARY KEY,
    nomecompleto VARCHAR(50) NOT NULL,
    cartaosus CHAR(15) NOT NULL,
    nomecompletodamae VARCHAR(50) NOT NULL,
    datadenascimento DATE NOT NULL,
    email VARCHAR(100) NOT NULL,
    telefone VARCHAR(15) NOT NULL,
    senha VARCHAR(100) NOT NULL,
    apelido VARCHAR(20),
    raca VARCHAR(20),
    nacionalidade VARCHAR(20),
    escolaridade VARCHAR(20)
);