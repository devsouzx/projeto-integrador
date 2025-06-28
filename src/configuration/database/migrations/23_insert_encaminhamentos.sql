-- Inserção de encaminhamentos de teste, alguns associados a laudos existentes
INSERT INTO encaminhamentos (
    id, paciente_id, medico_id, laudo_id, especialidade, urgencia, 
    justificativa, unidade_referencia, status, 
    data_encaminhamento, data_agendamento, data_conclusao, observacoes
) VALUES
-- Encaminhamento 1: Normal follow-up (sem laudo associado)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '12345678901' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
     (SELECT id FROM laudos WHERE resultado = 'ASC-US' AND status = 'concluido' LIMIT 1),
    'Ginecologia',
    'Rotina',
    'Controle periódico conforme recomendação do laudo anterior.',
    'UBS Central',
    'pendente',
    '2025-01-21 10:00:00',
    '2025-01-21 10:00:00',
    '2025-01-21 10:00:00',
    'Paciente deve agendar retorno em 3 anos conforme recomendação.'
),

-- Encaminhamento 2: ASC-US follow-up (associado ao laudo 2)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '12345678901' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM laudos WHERE resultado = 'ASC-US' AND status = 'concluido' LIMIT 1),
    'Ginecologia',
    'Prioritário',
    'Resultado ASC-US, necessita acompanhamento em 6 meses.',
    'Ambulatório de Ginecologia',
    'agendado',
    '2025-02-16 09:15:00',
    '2025-08-16 14:00:00',
    '2025-01-21 10:00:00',
    'Orientações sobre o resultado foram dadas à paciente.'
),

-- Encaminhamento 3: ASC-H (associado ao laudo 3)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '12345678901' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM laudos WHERE resultado = 'ASC-H' LIMIT 1),
    'Colposcopia',
    'Urgente',
    'Resultado ASC-H, necessita colposcopia com biópsia.',
    'Centro de Referência em Saúde da Mulher',
    'concluido',
    '2025-03-09 11:30:00',
    '2025-03-20 08:30:00',
    '2025-03-20 10:45:00',
    'Paciente compareceu e foi realizada colposcopia com biópsia.'
),

-- Encaminhamento 4: LSIL (associado ao laudo 4)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '12345678901' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM laudos WHERE resultado = 'LSIL' LIMIT 1),
    'Ginecologia',
    'Prioritário',
    'Resultado LSIL, necessita acompanhamento especializado.',
    'Ambulatório de Patologia Cervical',
    'agendado',
    '2025-04-23 16:20:00',
    '2025-05-15 10:00:00',
    '2025-01-21 10:00:00',
    'Paciente orientada sobre a necessidade de acompanhamento.'
),

-- Encaminhamento 5: HSIL (associado ao laudo 5)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '12345678901' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM laudos WHERE resultado = 'HSIL' LIMIT 1),
    'Oncologia Ginecológica',
    'Emergência',
    'Resultado HSIL, suspeita de neoplasia intraepitelial cervical.',
    'Hospital do Câncer',
    'concluido',
    '2025-05-16 08:00:00',
    '2025-05-18 09:00:00',
    '2025-05-25 11:00:00',
    'Paciente encaminhada com urgência, realizou novos exames.'
),

-- Encaminhamento 6: Carcinoma (associado ao laudo 6)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '12345678901' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM laudos WHERE resultado = 'Carcinoma' LIMIT 1),
    'Oncologia',
    'Emergência',
    'Resultado sugestivo de carcinoma escamoso invasivo.',
    'Centro de Tratamento Oncológico',
    'pendente',
    '2025-06-04 07:45:00',
    '2025-01-21 10:00:00',
    '2025-01-21 10:00:00',
    'Encaminhamento urgente, aguardando confirmação de vaga.'
),

-- Encaminhamento 7: Rascunho (associado ao laudo em rascunho)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '12345678901' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM laudos WHERE status = 'rascunho' LIMIT 1),
    'Ginecologia',
    'Rotina',
    'Possível ASC-US, aguardando confirmação do laudo.',
    'UBS Norte',
    'rascunho',
    '2025-07-19 14:30:00',
    '2025-01-21 10:00:00',
    '2025-01-21 10:00:00',
    'Aguardando conclusão do laudo para confirmar encaminhamento.'
),

-- Encaminhamento 8: Sem laudo associado (outro caso)
(
    gen_random_uuid(),
    (SELECT id FROM paciente WHERE cpf = '12345678901' LIMIT 1),
    (SELECT id FROM medico WHERE crm = '12345GO' LIMIT 1),
    (SELECT id FROM laudos WHERE resultado = 'ASC-US' AND status = 'concluido' LIMIT 1),
    'Cardiologia',
    'Prioritário',
    'Sopros cardíacos detectados durante exame ginecológico.',
    'Ambulatório de Cardiologia',
    'agendado',
    '2025-02-28 11:20:00',
    '2025-03-15 15:30:00',
    '2025-01-21 10:00:00',
    'Encaminhamento por achado incidental durante consulta de rotina.'
);