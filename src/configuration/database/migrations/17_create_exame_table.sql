CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS exame;

CREATE TABLE exame (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    inspecao_colo VARCHAR(50),
    sinais_dst VARCHAR(20),
    observacoes TEXT,
    paciente_id UUID NOT NULL,
    ficha_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_exame_paciente
        FOREIGN KEY (paciente_id)
        REFERENCES paciente(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_exame_ficha
        FOREIGN KEY (ficha_id)
        REFERENCES ficha(id)
        ON DELETE CASCADE
);