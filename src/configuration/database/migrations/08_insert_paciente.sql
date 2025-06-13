INSERT INTO paciente (
    cpf,
    nomecompleto,
    cartaosus,
    nomecompletodamae,
    datadenascimento,
    email,
    telefone,
    senha,
    apelido,
    raca,
    nacionalidade,
    escolaridade
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
    'Ensino Superior Incompleto'
);
