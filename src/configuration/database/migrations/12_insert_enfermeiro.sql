INSERT INTO enfermeiro (
    nomecompleto, 
    cpf, 
    email,
    telefone, 
    coren, 
    senha, 
    unidade_saude_id
)
VALUES (
    'Enfermeira Teste',
    '10987654321',
    'enfermeiro@ufg.br',
    '62999994444',
    'GO123456',
    crypt('@Teste123', gen_salt('bf')),
    (SELECT id FROM unidade_saude WHERE cnes = '1234567')
);
