-- Supondo que já exista uma unidade com CNES '1234567'
INSERT INTO gestor (
    nomecompleto, 
    cpf,
    email,
    telefone, 
    senha, 
    unidade_saude_id
)
VALUES (
    'Gestor Teste',
    '12345678901',
    'gestor@ufg.br',
    '62988887777',
    crypt('@Teste123', gen_salt('bf')),
    5915120
);
