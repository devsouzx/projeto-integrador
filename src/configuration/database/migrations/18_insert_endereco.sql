-- Endereço para Paciente Teste
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74000000', 'Rua 10', '100', 'Apto 101', 'Setor Central', 'Goiânia', 'GO', 'Próximo ao Mercado Central', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '12345678902')
);

-- Endereço para Paciente 1 (Ana Clara Souza)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74210010', 'Rua 57', '200', 'Casa 2', 'Setor Marista', 'Goiânia', 'GO', 'Próximo ao Colégio Marista', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '12345678901')
);

-- Endereço para Paciente 2 (Carlos Eduardo Lima)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74823050', 'Avenida T-10', '500', 'Bloco B', 'Setor Bueno', 'Goiânia', 'GO', 'Próximo ao Flamboyant', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '23456789012')
);

-- Endereço para Paciente 3 (Maria das Graças Oliveira)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74690010', 'Rua 9', '300', 'Fundos', 'Setor Leste Universitário', 'Goiânia', 'GO', 'Próximo à UFG', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '34567890123')
);

-- Endereço para Paciente 4 (Juliana Santos)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74333010', 'Rua C-146', '150', 'Apto 302', 'Setor Nova Suíça', 'Goiânia', 'GO', 'Próximo ao Parque Vaca Brava', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '45678901234')
);

-- Endereço para Paciente 5 (Pedro Henrique Alves)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74825070', 'Rua T-63', '80', 'Casa', 'Setor Bueno', 'Goiânia', 'GO', 'Próximo à Praça do Sol', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '56789012345')
);

-- Endereço para Paciente 6 (Juan Carlos Mendez)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74423030', 'Rua 68', '45', 'Apto 101', 'Setor Sul', 'Goiânia', 'GO', 'Próximo ao Shopping Sul', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '67890123456')
);

-- Endereço para Paciente 7 (Aruã Pataxó)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74660020', 'Rua 136', '22', 'Chácara', 'Setor Marista', 'Goiânia', 'GO', 'Próximo ao Bosque dos Buritis', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '78901234567')
);

-- Endereço para Paciente 8 (Roberto Ferreira)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74535070', 'Rua 104', '350', 'Casa', 'Setor Sul', 'Goiânia', 'GO', 'Próximo ao Hospital das Clínicas', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '89012345678')
);

-- Endereço para Paciente 9 (Sônia Regina Gomes)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74305050', 'Rua 16', '120', 'Apto 201', 'Setor Oeste', 'Goiânia', 'GO', 'Próximo ao Estádio Serra Dourada', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '90123456789')
);

-- Endereço para Paciente 10 (Francisco da Silva)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74610040', 'Rua 3', '50', 'Casa', 'Setor Criméia Leste', 'Goiânia', 'GO', 'Próximo à Praça da Criméia', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '01234567890')
);

-- Endereço para Paciente 11 (Fernanda Oliveira)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74255020', 'Rua 91', '600', 'Apto 501', 'Setor Sul', 'Goiânia', 'GO', 'Próximo ao Parque Mutirama', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '11122233344')
);

-- Endereço para Paciente 12 (Ricardo Nunes)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74315010', 'Rua 20', '75', 'Casa', 'Setor Pedro Ludovico', 'Goiânia', 'GO', 'Próximo ao Tribunal de Justiça', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '22233344455')
);

-- Endereço para Paciente 13 (Antônia Bezerra)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74680060', 'Rua 132', '30', 'Apto 102', 'Setor Leste Vila Nova', 'Goiânia', 'GO', 'Próximo ao Hospital Memorial', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '33344455566')
);

-- Endereço para Paciente 14 (Camila Rocha)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74303030', 'Rua 15', '200', 'Apto 303', 'Setor Oeste', 'Goiânia', 'GO', 'Próximo ao Zoológico', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '44455566677')
);

-- Endereço para Paciente 15 (Lucas Mendes)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74555070', 'Rua 107', '90', 'Casa', 'Setor Sul', 'Goiânia', 'GO', 'Próximo ao Colégio Objetivo', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '55566677788')
);

-- Endereço para Paciente 16 (Sophie Martin)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74815040', 'Rua T-27', '400', 'Apto 201', 'Setor Bueno', 'Goiânia', 'GO', 'Próximo à Praça Tamandaré', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '66677788899')
);

-- Endereço para Paciente 17 (Jaci Tupinambá)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74670050', 'Rua 122', '25', 'Chácara', 'Setor Marista', 'Goiânia', 'GO', 'Próximo ao Parque Areião', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '77788899900')
);

-- Endereço para Paciente 18 (Mauro Andrade)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74325010', 'Rua 25', '180', 'Casa', 'Setor Pedro Ludovico', 'Goiânia', 'GO', 'Próximo ao Fórum', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '88899900011')
);

-- Endereço para Paciente 19 (Lúcia Pereira)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74455030', 'Rua 75', '220', 'Apto 401', 'Setor Sul', 'Goiânia', 'GO', 'Próximo ao Shopping Bougainville', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '99900011122')
);

-- Endereço para Paciente 20 (Dona Maria)
INSERT INTO endereco (
    cep, logradouro, numero, complemento, bairro, cidade, uf, referencia, codigo_municipio, paciente_id
) VALUES (
    '74650010', 'Rua 108', '10', 'Casa', 'Setor Leste Universitário', 'Goiânia', 'GO', 'Próximo ao Campus II da PUC', '5208707', 
    (SELECT id FROM paciente WHERE cpf = '00011122233')
);