CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS encaminhamentos;

CREATE TABLE encaminhamentos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    paciente_id UUID NOT NULL REFERENCES paciente(id),
    medico_id UUID NOT NULL REFERENCES medico(id),
    laudo_id UUID REFERENCES laudos(id),
    especialidade VARCHAR(100) NOT NULL,
    urgencia VARCHAR(20) NOT NULL,
    justificativa TEXT NOT NULL,
    unidade_referencia VARCHAR(100) NOT NULL,
    status VARCHAR(20) NOT NULL,
    data_encaminhamento TIMESTAMP NOT NULL,
    data_agendamento TIMESTAMP,
    data_conclusao TIMESTAMP,
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_encaminhamento_paciente_id ON encaminhamentos(paciente_id);
CREATE INDEX IF NOT EXISTS idx_encaminhamento_medico_id ON encaminhamentos(medico_id);
CREATE INDEX IF NOT EXISTS idx_encaminhamento_status ON encaminhamentos(status);