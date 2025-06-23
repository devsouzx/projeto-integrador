CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS ficha CASCADE;

CREATE TABLE ficha (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    protocolo VARCHAR(20) NOT NULL UNIQUE,
    data_coleta TIMESTAMP NOT NULL,
    responsavel_coleta VARCHAR(100) NOT NULL,
    motivo_exame VARCHAR(50) NOT NULL,
    observacoes TEXT,
    responsavel_id UUID NOT NULL,
    unidade_id INTEGER,
    paciente_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_ficha_paciente
        FOREIGN KEY (paciente_id)
        REFERENCES paciente(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_ficha_id ON ficha(id);
CREATE INDEX idx_ficha_paciente_id ON ficha(paciente_id);