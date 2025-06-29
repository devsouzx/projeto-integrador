CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS agendamento CASCADE;

CREATE TABLE agendamento (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    paciente_id UUID NOT NULL REFERENCES paciente(id),
    profissional_id UUID,
    profissional_tipo VARCHAR(20),
    data_hora TIMESTAMP NOT NULL,
    tipo_consulta VARCHAR(50) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pendente',
    observacoes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_agendamento_profissional ON agendamento(profissional_id, profissional_tipo);
CREATE INDEX idx_agendamento_paciente ON agendamento(paciente_id);
CREATE INDEX idx_agendamento_data ON agendamento(data_hora);