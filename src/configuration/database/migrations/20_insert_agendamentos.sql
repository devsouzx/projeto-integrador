INSERT INTO enfermeiro (
    nomecompleto, 
    cpf, 
    email, 
    telefone, 
    coren, 
    senha, 
    unidade_saude_id
) VALUES 
(
    'Enfermeira Ana Silva', 
    '11122233344', 
    'ana.silva@saude.gov.br', 
    '61999887766', 
    'COREN-DF 123456', 
    crypt('Ana@1234', gen_salt('bf')), 
    1
),
(
    'Enfermeiro Carlos Oliveira', 
    '22233344455', 
    'carlos.oliveira@saude.gov.br', 
    '61988776655', 
    'COREN-DF 654321', 
    crypt('Carlos@1234', gen_salt('bf')), 
    2
),
(
    'Enfermeira Mariana Santos', 
    '33344455566', 
    'mariana.santos@saude.gov.br', 
    '61977665544', 
    'COREN-DF 987654', 
    crypt('Mariana@1234', gen_salt('bf')), 
    3
);

INSERT INTO agendamento (
    paciente_id,
    profissional_id,
    profissional_tipo,
    data_hora,
    tipo_consulta,
    status,
    observacoes
) VALUES
-- Agendamentos confirmados
(
    (SELECT id FROM paciente WHERE cpf = '12345678901'),
    (SELECT id FROM enfermeiro WHERE cpf = '10987654321'),
    'enfermeiro',
    '2025-12-01 09:00:00',
    'Vacinação',
    'confirmado',
    'Paciente agendada para vacina de reforço contra hepatite B.'
),
(
    (SELECT id FROM paciente WHERE cpf = '23456789012'),
    (SELECT id FROM enfermeiro WHERE cpf = '10987654321'),
    'enfermeiro',
    '2025-12-01 10:30:00',
    'Curativo',
    'confirmado',
    'Revisão de curativo em ferida no pé esquerdo.'
),

-- Agendamentos pendentes
(
    (SELECT id FROM paciente WHERE cpf = '34567890123'),
    (SELECT id FROM enfermeiro WHERE cpf = '10987654321'),
    'enfermeiro',
    '2025-12-02 14:00:00',
    'Aferição de pressão',
    'pendente',
    'Paciente solicitou verificação da pressão arterial após sintomas de tontura.'
),

-- Agendamentos concluídos
(
    (SELECT id FROM paciente WHERE cpf = '45678901234'),
    (SELECT id FROM enfermeiro WHERE cpf = '10987654321'),
    'enfermeiro',
    '2025-11-25 08:30:00',
    'Pré-natal',
    'concluido',
    'Consulta de rotina do pré-natal – 24 semanas de gestação.'
),

-- Agendamentos para o Enfermeiro Carlos Oliveira
(
    (SELECT id FROM paciente WHERE cpf = '56789012345'),
    (SELECT id FROM enfermeiro WHERE cpf = '10987654321'),
    'enfermeiro',
    '2025-12-03 09:00:00',
    'Vacinação infantil',
    'confirmado',
    'Vacinação contra sarampo – 2ª dose.'
),
(
    (SELECT id FROM paciente WHERE cpf = '67890123456'),
    (SELECT id FROM enfermeiro WHERE cpf = '10987654321'),
    'enfermeiro',
    '2025-12-03 10:30:00',
    'Teste de glicemia',
    'pendente',
    'Paciente relatou sintomas de hipoglicemia recentemente.'
),

-- Agendamento cancelado
(
    (SELECT id FROM paciente WHERE cpf = '78901234567'),
    (SELECT id FROM enfermeiro WHERE cpf = '10987654321'),
    'enfermeiro',
    '2025-11-28 11:00:00',
    'Curativo',
    'cancelado',
    'Consulta cancelada por ausência do paciente.'
),

-- Agendamentos para a Enfermeira Mariana Santos
(
    (SELECT id FROM paciente WHERE cpf = '89012345678'),
    (SELECT id FROM enfermeiro WHERE cpf = '10987654321'),
    'enfermeiro',
    '2025-12-04 08:00:00',
    'Aferição de pressão',
    'confirmado',
    'Paciente hipertenso realizando acompanhamento semanal.'
),
(
    (SELECT id FROM paciente WHERE cpf = '90123456789'),
    (SELECT id FROM enfermeiro WHERE cpf = '10987654321'),
    'enfermeiro',
    '2025-12-04 09:30:00',
    'Aplicação de insulina',
    'confirmado',
    'Insulina aplicada conforme prescrição médica.'
),
(
    (SELECT id FROM paciente WHERE cpf = '01234567890'),
    (SELECT id FROM enfermeiro WHERE cpf = '10987654321'),
    'enfermeiro',
    '2025-11-30 14:00:00',
    'Curativo',
    'concluido',
    'Curativo trocado com sucesso – evolução positiva.'
);
