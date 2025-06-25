INSERT INTO encaminhamentos (
    especialidade,
    urgencia,
    justificativa,
    unidade_referencia,
    data_encaminhamento,
    data_agendamento,
    data_conclusao,
    observacoes,
    status,
    paciente_id,
    medico_id,
    laudo_id
) VALUES (
    'ghfcvxvxv',
    'xvxvxvxv',
    'kkhmnbnb',
    'hknbmnbmnb',
    '2025-08-09',
    '2025-08-09',
    '2025-08-09',
    'sdfsfsfsffsdf',
    'sdfsfsdfs',
    (SELECT id FROM paciente WHERE cpf = '12345678901'),
    (SELECT id FROM medico WHERE crm = '12345GO'),
    (SELECT id FROM laudos WHERE paciente_id = (SELECT id FROM paciente WHERE cpf = '12345678901'))
)