INSERT INTO paciente (
    cpf,
    nome,
    cns,
    nome_mae,
    data_nascimento,
    email,
    telefone,
    senha,
    apelido,
    raca,
    nacionalidade,
    escolaridade, 
    is_verified
) VALUES (
    '12345678901',
    'Paciente Teste',
    '123456789012345',
    'Mae do Paciente',
    '2007-09-07',
    'paciente@ufg.br',
    '62981345678',
    crypt('@Teste123', gen_salt('bf')),
    'PaciTest',
    'Pardo',
    'Brasileiro',
    'Ensino Superior Incompleto',
    true
);