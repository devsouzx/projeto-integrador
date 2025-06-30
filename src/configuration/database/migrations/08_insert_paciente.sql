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
    '12345678902',
    'Paciente Teste',
    '123456789012345',
    'Mae do Paciente',
    '2007-09-07',
    'paciente@ufg.br',
    '62981743824',
    crypt('@Teste123', gen_salt('bf')),
    'PaciTest',
    'Pardo',
    'Brasileiro',
    'Ensino Superior Incompleto',
    true
);

-- Paciente 1 (Jovem)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '12345678901', 'Ana Clara Souza', '111111111111111', 'Maria Souza', '2000-05-15',
    'ana.clara@email.com', '61987654321', crypt('Senha123!', gen_salt('bf')),
    'Aninha', 'Branca', 'Brasileira', 'Ensino Médio Completo', true
);

-- Paciente 2 (Adulto)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '23456789012', 'Carlos Eduardo Lima', '222222222222222', 'Joana Lima', '1985-11-22',
    'carlos.lima@email.com', '61976543210', crypt('Carlos@123', gen_salt('bf')),
    'Cadú', 'Pardo', 'Brasileira', 'Ensino Superior Completo', true
);

-- Paciente 3 (Idoso)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '34567890123', 'Maria das Graças Oliveira', '333333333333333', 'Antônia Oliveira', '1952-03-08',
    'maria.oliveira@email.com', '61965432109', crypt('Maria1952#', gen_salt('bf')),
    'Gracinha', 'Negra', 'Brasileira', 'Ensino Fundamental Completo', true
);

-- Paciente 4 (Gestante)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '45678901234', 'Juliana Santos', '444444444444444', 'Rosa Santos', '1993-07-30',
    'juliana.s@email.com', '61954321098', crypt('Juliana93$', gen_salt('bf')),
    'Juju', 'Branca', 'Brasileira', 'Ensino Superior Completo', true
);

-- Paciente 5 (Criança)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '56789012345', 'Pedro Henrique Alves', '555555555555555', 'Patrícia Alves', '2015-02-14',
    'ph.alves@email.com', '61943210987', crypt('Pedro2015!', gen_salt('bf')),
    'Pedrinho', 'Pardo', 'Brasileira', 'Ensino Fundamental Incompleto', true
);

-- Paciente 6 (Estrangeiro)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '67890123456', 'Juan Carlos Mendez', '666666666666666', 'Isabel Mendez', '1988-09-05',
    'juan.mendez@email.com', '61932109876', crypt('JuanCM88*', gen_salt('bf')),
    'Juanito', 'Branca', 'Venezuelana', 'Ensino Superior Completo', true
);

-- Paciente 7 (Indígena)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '78901234567', 'Aruã Pataxó', '777777777777777', 'Jandira Pataxó', '1997-12-25',
    'arua.pataxo@email.com', '61921098765', crypt('Arua1997@', gen_salt('bf')),
    'Aruã', 'Indígena', 'Brasileira', 'Ensino Médio Completo', true
);

-- Paciente 8 (Hipertenso)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '89012345678', 'Roberto Ferreira', '888888888888888', 'Marta Ferreira', '1972-04-18',
    'roberto.ferreira@email.com', '61910987654', crypt('Roberto72#', gen_salt('bf')),
    'Beto', 'Negra', 'Brasileira', 'Ensino Fundamental Completo', true
);

-- Paciente 9 (Diabético)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '90123456789', 'Sônia Regina Gomes', '999999999999999', 'Adélia Gomes', '1965-08-11',
    'sonia.gomes@email.com', '61909876543', crypt('Sonia1965$', gen_salt('bf')),
    'Soninha', 'Parda', 'Brasileira', 'Ensino Médio Completo', true
);

-- Paciente 10 (Prioritário)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '01234567890', 'Francisco da Silva', '000000000000000', 'Francisca Silva', '1948-01-03',
    'chico.silva@email.com', '61998765432', crypt('Chico1948!', gen_salt('bf')),
    'Chico', 'Pardo', 'Brasileira', 'Ensino Fundamental Incompleto', true
);

-- Paciente 11 (Jovem Adulto)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '11122233344', 'Fernanda Oliveira', '123456789012346', 'Márcia Oliveira', '1998-08-25',
    'fernanda.oliveira@email.com', '61911223344', crypt('Fer123@', gen_salt('bf')),
    'Fê', 'Branca', 'Brasileira', 'Ensino Superior Completo', true
);

-- Paciente 12 (Meia-idade)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '22233344455', 'Ricardo Nunes', '223456789012345', 'Alice Nunes', '1975-04-12',
    'ricardo.nunes@email.com', '61922334455', crypt('Ric@1975', gen_salt('bf')),
    'Ric', 'Pardo', 'Brasileira', 'Ensino Médio Completo', true
);

-- Paciente 13 (Idoso)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '33344455566', 'Antônia Bezerra', '323456789012345', 'Carmem Bezerra', '1949-11-30',
    'antonia.b@email.com', '61933445566', crypt('Tonha49#', gen_salt('bf')),
    'Tonha', 'Negra', 'Brasileira', 'Ensino Fundamental Completo', true
);

-- Paciente 14 (Gestante)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '44455566677', 'Camila Rocha', '423456789012345', 'Silvia Rocha', '1995-02-14',
    'camila.rocha@email.com', '61944556677', crypt('Cami95$', gen_salt('bf')),
    'Cami', 'Branca', 'Brasileira', 'Ensino Superior Incompleto', true
);

-- Paciente 15 (Adolescente)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '55566677788', 'Lucas Mendes', '523456789012345', 'Patrícia Mendes', '2007-07-07',
    'lucas.mendes@email.com', '61955667788', crypt('Lu2007!', gen_salt('bf')),
    'Luca', 'Pardo', 'Brasileira', 'Ensino Fundamental Incompleto', true
);

-- Paciente 16 (Estrangeiro)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '66677788899', 'Sophie Martin', '623456789012345', 'Claire Martin', '1990-09-21',
    'sophie.martin@email.com', '61966778899', crypt('SophieFR90', gen_salt('bf')),
    'Soph', 'Branca', 'Francesa', 'Ensino Superior Completo', true
);

-- Paciente 17 (Indígena)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '77788899900', 'Jaci Tupinambá', '723456789012345', 'Araci Tupinambá', '1988-03-19',
    'jaci.t@email.com', '61977889900', crypt('Jaci1988', gen_salt('bf')),
    'Jaci', 'Indígena', 'Brasileira', 'Ensino Médio Completo', true
);

-- Paciente 18 (Hipertenso)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '88899900011', 'Mauro Andrade', '823456789012345', 'Rita Andrade', '1968-12-05',
    'mauro.a@email.com', '61988990011', crypt('Mauro68#', gen_salt('bf')),
    'Mau', 'Negra', 'Brasileira', 'Ensino Fundamental Completo', true
);

-- Paciente 19 (Diabético)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '99900011122', 'Lúcia Pereira', '923456789012345', 'Maria Pereira', '1972-06-28',
    'lucia.pereira@email.com', '61999001122', crypt('Lu1972$', gen_salt('bf')),
    'Lú', 'Parda', 'Brasileira', 'Ensino Médio Completo', true
);

-- Paciente 20 (Prioritário)
INSERT INTO paciente (
    cpf, nome, cns, nome_mae, data_nascimento, email, telefone, 
    senha, apelido, raca, nacionalidade, escolaridade, is_verified
) VALUES (
    '00011122233', 'Dona Maria', '023456789012345', 'Joana Silva', '1939-05-01',
    'dona.maria@email.com', '61900112233', crypt('Maria1939', gen_salt('bf')),
    'Dona', 'Branca', 'Brasileira', 'Analfabeto', true
);