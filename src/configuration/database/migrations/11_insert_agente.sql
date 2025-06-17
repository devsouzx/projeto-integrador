INSERT INTO agente_comunitario (
    nomecompleto, 
    cpf,
    telefone, 
    email, 
    senha,
    unidade_saude_id
)
VALUES (
    'Agente Teste',
    '12345678901',
    '62989990000',
    'agente@ufg.br',
    crypt('@Teste123', gen_salt('bf')),
    5915120
);
