CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS laudos CASCADE;

CREATE TABLE laudos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    paciente_id UUID NOT NULL REFERENCES paciente(id),
    medico_id UUID NOT NULL REFERENCES medico(id),
    ficha_id UUID REFERENCES ficha(id),
    data_exame TIMESTAMP NOT NULL,
    data_laudo TIMESTAMP NOT NULL,
    resultado VARCHAR(20) NOT NULL,
    cid VARCHAR(10) NOT NULL,
    adequabilidade VARCHAR(50),
    microbiologia VARCHAR(100),
    celulas_endometriais BOOLEAN DEFAULT FALSE,
    comentarios TEXT,
    recomendacoes TEXT,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_laudos_paciente_id ON laudos(paciente_id);
CREATE INDEX IF NOT EXISTS idx_laudos_ficha_id ON laudos(ficha_id);
CREATE INDEX IF NOT EXISTS idx_laudos_medico_id ON laudos(medico_id);

-- Inserção de laudos de teste com ficha_id
INSERT INTO laudos (
    id, paciente_id, medico_id, ficha_id, data_exame, data_laudo, 
    resultado, cid, adequabilidade, microbiologia, 
    celulas_endometriais, comentarios, recomendacoes, status
) VALUES
-- Laudo 1: Resultado Normal (associado à ficha PROT20230001)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '12345678901' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230001' LIMIT 1),
    '2025-01-15 09:00:00',
    '2025-01-20 14:30:00',
    'Normal',
    'Z01.4',
    'Satisfatória',
    'Flora bacteriana mista',
    FALSE,
    'Exame sem alterações significativas.',
    'Retornar para novo exame em 3 anos.',
    'concluido'
),

-- Laudo 2: ASC-US (associado à ficha PROT20230002)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '23456789012' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230002' LIMIT 1),
    '2025-02-10 10:30:00',
    '2025-02-15 16:45:00',
    'ASC-US',
    'R87.61',
    'Satisfatória',
    'Flora bacteriana mista com cocos',
    FALSE,
    'Atipias em células escamosas de significado indeterminado.',
    'Repetir exame em 6 meses.',
    'concluido'
),

-- Laudo 3: ASC-H (associado à ficha PROT20230011)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '11122233344' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230011' LIMIT 1),
    '2025-03-05 11:15:00',
    '2025-03-08 10:00:00',
    'ASC-H',
    'R87.61',
    'Parcialmente satisfatória',
    'Flora bacteriana reduzida',
    TRUE,
    'Atipias em células escamosas, não pode excluir HSIL.',
    'Realizar colposcopia com biópsia.',
    'concluido'
),

-- Laudo 4: LSIL (associado à ficha PROT20230015)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '55566677788' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230015' LIMIT 1),
    '2025-04-20 08:45:00',
    '2025-04-22 15:20:00',
    'LSIL',
    'R87.61',
    'Satisfatória',
    'Flora bacteriana mista',
    FALSE,
    'Lesão intraepitelial de baixo grau.',
    'Repetir exame em 1 ano ou realizar colposcopia.',
    'concluido'
),

-- Laudo 5: HSIL (associado à ficha PROT20230004)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '45678901234' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230004' LIMIT 1),
    '2025-05-12 14:00:00',
    '2025-05-15 11:30:00',
    'HSIL',
    'D06.9',
    'Satisfatória',
    'Flora bacteriana escassa',
    FALSE,
    'Lesão intraepitelial de alto grau.',
    'Encaminhar para colposcopia e biópsia imediatamente.',
    'concluido'
),

-- Laudo 6: Carcinoma (associado à ficha PROT20230013)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '33344455566' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230013' LIMIT 1),
    '2025-06-01 10:00:00',
    '2025-06-03 09:15:00',
    'Carcinoma',
    'C53.9',
    'Insatisfatória',
    'Flora bacteriana ausente',
    TRUE,
    'Células sugestivas de carcinoma escamoso invasivo.',
    'Encaminhamento urgente para oncologia.',
    'concluido'
),

-- Laudo 7: Rascunho (associado à ficha PROT20230017)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '77788899900' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230017' LIMIT 1),
    '2025-07-18 13:30:00',
    '2025-07-18 13:30:00',
    'ASC-US',
    'R87.61',
    'Satisfatória',
    'Flora bacteriana mista',
    FALSE,
    'Possível atipia, necessita revisão.',
    'Aguardando confirmação do patologista.',
    'rascunho'
);