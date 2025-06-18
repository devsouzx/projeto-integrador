CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS anamnese;

CREATE TABLE anamnese (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    motivo_exame VARCHAR(50),
    fez_exame VARCHAR(20),
    ultimo_exame_ano VARCHAR(4),
    usa_diu VARCHAR(20),
    gravida VARCHAR(20),
    anticoncepcional VARCHAR(20),
    hormonio_menopausa VARCHAR(20),
    radioterapia VARCHAR(20),
    dum DATE,
    nao_lembra_dum BOOLEAN DEFAULT false,
    sangramento_pos_coito VARCHAR(50),
    sangramento_pos_menopausa VARCHAR(50),
    paciente_id UUID NOT NULL,
    ficha_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_anamnase_paciente
        FOREIGN KEY (paciente_id)
        REFERENCES paciente(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_anamnase_ficha
        FOREIGN KEY (ficha_id)
        REFERENCES ficha(id)
        ON DELETE CASCADE
);