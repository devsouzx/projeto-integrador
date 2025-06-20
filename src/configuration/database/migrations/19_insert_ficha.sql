-- Ficha para Paciente 1 (Ana Clara Souza)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230001', '2023-01-10 08:30:00', 'Enfermeira Maria Silva', 'Rotina', 'Paciente assintomática', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '12345678901')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Sim', '2021', 'Não', 'Não', 'Sim', 'Não', 'Não', '2023-01-01', false, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '12345678901'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230001')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Colo sem alterações',
    (SELECT id FROM paciente WHERE cpf = '12345678901'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230001')
);

-- Ficha para Paciente 2 (Carlos Eduardo Lima)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id, unidade_id
) VALUES (
    'PROT20230002', '2023-01-15 09:45:00', 'Enfermeiro João Santos', 'Sangramento', 'Paciente com queixa de sangramento', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '23456789012'),
    2506327
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Sangramento', 'Não', '2019', 'Não', 'Não', 'Não', 'Não', 'Não', NULL, true, 
    'Sim', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '23456789012'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230002')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Alterado', 'Não', 'Presença de lesão no colo',
    (SELECT id FROM paciente WHERE cpf = '23456789012'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230002')
);

-- Ficha para Paciente 3 (Maria das Graças Oliveira)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230003', '2023-01-20 10:15:00', 'Enfermeira Ana Paula', 'Rotina', 'Paciente idosa - acompanhamento anual', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '34567890123')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Sim', '2020', 'Não', 'Não', 'Não', 'Sim', 'Não', NULL, true, 
    'Não', 'Não',
    (SELECT id FROM paciente WHERE cpf = '34567890123'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230003')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Colo sem alterações',
    (SELECT id FROM paciente WHERE cpf = '34567890123'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230003')
);

-- Ficha para Paciente 4 (Juliana Santos - Gestante)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230004', '2023-02-05 11:30:00', 'Enfermeira Carla Mendes', 'Pré-natal', 'Paciente gestante - 12 semanas', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '45678901234')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Pré-natal', 'Sim', '2022', 'Não', 'Sim', 'Não', 'Não', 'Não', '2023-01-10', false, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '45678901234'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230004')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Colo uterino sem alterações',
    (SELECT id FROM paciente WHERE cpf = '45678901234'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230004')
);

-- Ficha para Paciente 5 (Pedro Henrique Alves - Criança)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230005', '2023-02-10 14:00:00', 'Enfermeiro Pedro Alves', 'Queixa', 'Paciente pediátrico com queixa dos pais', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '56789012345')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Queixa', 'Não', NULL, 'Não', 'Não', 'Não', 'Não', 'Não', NULL, true, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '56789012345'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230005')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Exame sem alterações',
    (SELECT id FROM paciente WHERE cpf = '56789012345'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230005')
);

-- Ficha para Paciente 6 (Juan Carlos Mendez - Estrangeiro)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230006', '2023-02-15 15:20:00', 'Enfermeira Luana Costa', 'Rotina', 'Paciente estrangeiro - primeiro exame no Brasil', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '67890123456')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Sim', '2021', 'Não', 'Não', 'Não', 'Não', 'Não', '2023-02-01', false, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '67890123456'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230006')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Colo sem alterações visíveis',
    (SELECT id FROM paciente WHERE cpf = '67890123456'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230006')
);

-- Ficha para Paciente 7 (Aruã Pataxó - Indígena)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230007', '2023-03-01 10:45:00', 'Enfermeira Sofia Mendonça', 'Rotina', 'Paciente indígena - acompanhamento especial', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '78901234567')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Não', NULL, 'Não', 'Não', 'Não', 'Não', 'Não', NULL, true, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '78901234567'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230007')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Exame sem alterações',
    (SELECT id FROM paciente WHERE cpf = '78901234567'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230007')
);

-- Ficha para Paciente 8 (Roberto Ferreira - Hipertenso)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230008', '2023-03-10 13:15:00', 'Enfermeira Raquel Gomes', 'Rotina', 'Paciente hipertenso - acompanhamento', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '89012345678')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Sim', '2020', 'Não', 'Não', 'Não', 'Não', 'Não', '2023-03-01', false, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '89012345678'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230008')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Colo sem alterações',
    (SELECT id FROM paciente WHERE cpf = '89012345678'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230008')
);

-- Ficha para Paciente 9 (Sônia Regina Gomes - Diabética)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230009', '2023-03-15 09:00:00', 'Enfermeiro Marcos Oliveira', 'Rotina', 'Paciente diabética - acompanhamento anual', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '90123456789')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Sim', '2020', 'Não', 'Não', 'Não', 'Sim', 'Não', NULL, true, 
    'Não', 'Sim',
    (SELECT id FROM paciente WHERE cpf = '90123456789'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230009')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Alterado', 'Não', 'Colo com aspecto atrófico',
    (SELECT id FROM paciente WHERE cpf = '90123456789'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230009')
);

-- Ficha para Paciente 10 (Francisco da Silva - Prioritário)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230010', '2023-04-01 11:45:00', 'Enfermeira Fernanda Lima', 'Sangramento', 'Paciente idoso com sangramento', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '01234567890')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Sangramento', 'Não', '2018', 'Não', 'Não', 'Não', 'Não', 'Não', NULL, true, 
    'Não', 'Sim',
    (SELECT id FROM paciente WHERE cpf = '01234567890'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230010')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Alterado', 'Não', 'Presença de lesão suspeita',
    (SELECT id FROM paciente WHERE cpf = '01234567890'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230010')
);

-- Ficha para Paciente 11 (Fernanda Oliveira)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230011', '2023-04-10 14:30:00', 'Enfermeiro Gustavo Henrique', 'Rotina', 'Paciente jovem adulta - checkup', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '11122233344')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Sim', '2022', 'Não', 'Não', 'Sim', 'Não', 'Não', '2023-04-01', false, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '11122233344'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230011')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Colo sem alterações',
    (SELECT id FROM paciente WHERE cpf = '11122233344'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230011')
);

-- Ficha para Paciente 12 (Ricardo Nunes)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230012', '2023-04-15 10:00:00', 'Enfermeira Patrícia Almeida', 'Rotina', 'Paciente meia-idade - checkup', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '22233344455')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Sim', '2021', 'Não', 'Não', 'Não', 'Não', 'Não', '2023-04-05', false, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '22233344455'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230012')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Exame sem alterações',
    (SELECT id FROM paciente WHERE cpf = '22233344455'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230012')
);

-- Ficha para Paciente 13 (Antônia Bezerra)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230013', '2023-05-02 08:45:00', 'Enfermeira Larissa Costa', 'Rotina', 'Paciente idosa - acompanhamento', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '33344455566')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Sim', '2020', 'Não', 'Não', 'Não', 'Sim', 'Não', NULL, true, 
    'Não', 'Não',
    (SELECT id FROM paciente WHERE cpf = '33344455566'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230013')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Colo sem alterações',
    (SELECT id FROM paciente WHERE cpf = '33344455566'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230013')
);

-- Ficha para Paciente 14 (Camila Rocha - Gestante)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230014', '2023-05-10 13:20:00', 'Enfermeira Juliana Pereira', 'Pré-natal', 'Paciente gestante - 20 semanas', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '44455566677')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Pré-natal', 'Sim', '2021', 'Não', 'Sim', 'Não', 'Não', 'Não', '2023-01-15', false, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '44455566677'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230014')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Colo sem alterações visíveis',
    (SELECT id FROM paciente WHERE cpf = '44455566677'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230014')
);

-- Ficha para Paciente 15 (Lucas Mendes - Adolescente)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230015', '2023-05-15 09:30:00', 'Enfermeiro Rodrigo Santos', 'Queixa', 'Paciente adolescente com queixa', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '55566677788')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Queixa', 'Não', NULL, 'Não', 'Não', 'Não', 'Não', 'Não', NULL, true, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '55566677788'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230015')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Exame sem alterações',
    (SELECT id FROM paciente WHERE cpf = '55566677788'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230015')
);

-- Ficha para Paciente 16 (Sophie Martin - Estrangeira)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230016', '2023-06-01 11:15:00', 'Enfermeira Carolina Mendes', 'Rotina', 'Paciente estrangeira - checkup', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '66677788899')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Sim', '2022', 'Não', 'Não', 'Não', 'Não', 'Não', '2023-05-20', false, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '66677788899'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230016')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Colo sem alterações',
    (SELECT id FROM paciente WHERE cpf = '66677788899'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230016')
);

-- Ficha para Paciente 17 (Jaci Tupinambá - Indígena)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230017', '2023-06-10 14:45:00', 'Enfermeiro Diego Oliveira', 'Rotina', 'Paciente indígena - acompanhamento', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '77788899900')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Sim', '2021', 'Não', 'Não', 'Não', 'Não', 'Não', '2023-06-01', false, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '77788899900'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230017')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Exame sem alterações',
    (SELECT id FROM paciente WHERE cpf = '77788899900'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230017')
);

-- Ficha para Paciente 18 (Mauro Andrade - Hipertenso)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230018', '2023-06-15 10:30:00', 'Enfermeira Vanessa Souza', 'Rotina', 'Paciente hipertenso - acompanhamento', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '88899900011')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Sim', '2021', 'Não', 'Não', 'Não', 'Não', 'Não', '2023-06-05', false, 
    'Não', 'Não aplicável',
    (SELECT id FROM paciente WHERE cpf = '88899900011'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230018')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Colo sem alterações',
    (SELECT id FROM paciente WHERE cpf = '88899900011'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230018')
);

-- Ficha para Paciente 19 (Lúcia Pereira - Diabética)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230019', '2023-07-01 08:15:00', 'Enfermeira Tatiana Rocha', 'Rotina', 'Paciente diabética - acompanhamento', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '99900011122')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Rotina', 'Sim', '2021', 'Não', 'Não', 'Não', 'Sim', 'Não', NULL, true, 
    'Não', 'Não',
    (SELECT id FROM paciente WHERE cpf = '99900011122'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230019')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Normal', 'Não', 'Colo sem alterações',
    (SELECT id FROM paciente WHERE cpf = '99900011122'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230019')
);

-- Ficha para Paciente 20 (Dona Maria - Prioritária)
INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id
) VALUES (
    'PROT20230020', '2023-07-10 09:00:00', 'Enfermeira Gabriela Lima', 'Sangramento', 'Paciente idosa com sangramento pós-menopausa', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '00011122233')
);

INSERT INTO anamnese (
    motivo_exame, fez_exame, ultimo_exame_ano, usa_diu, gravida, anticoncepcional, 
    hormonio_menopausa, radioterapia, dum, nao_lembra_dum, sangramento_pos_coito, 
    sangramento_pos_menopausa, paciente_id, ficha_id
) VALUES (
    'Sangramento', 'Não', '2018', 'Não', 'Não', 'Não', 'Não', 'Não', NULL, true, 
    'Não', 'Sim',
    (SELECT id FROM paciente WHERE cpf = '00011122233'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230020')
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Alterado', 'Não', 'Presença de lesão suspeita',
    (SELECT id FROM paciente WHERE cpf = '00011122233'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230020')
);

-- Atualizando as fichas com CNES de unidades de saúde de Goiânia/GO

-- Ficha para Paciente 1 (Ana Clara Souza) - UBS Jardim América (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200032 WHERE protocolo = 'PROT20230001';

-- Ficha para Paciente 3 (Maria das Graças Oliveira) - UBS Bueno (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200040 WHERE protocolo = 'PROT20230003';

-- Ficha para Paciente 4 (Juliana Santos - Gestante) - UBS Campinas (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200059 WHERE protocolo = 'PROT20230004';

-- Ficha para Paciente 5 (Pedro Henrique Alves - Criança) - Hospital Infantil (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200067 WHERE protocolo = 'PROT20230005';

-- Ficha para Paciente 6 (Juan Carlos Mendez - Estrangeiro) - UBS Setor Sul (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200075 WHERE protocolo = 'PROT20230006';

-- Ficha para Paciente 7 (Aruã Pataxó - Indígena) - Casa de Saúde Indígena (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200083 WHERE protocolo = 'PROT20230007';

-- Ficha para Paciente 8 (Roberto Ferreira - Hipertenso) - UBS Setor Oeste (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200091 WHERE protocolo = 'PROT20230008';

-- Ficha para Paciente 9 (Sônia Regina Gomes - Diabética) - UBS Setor Leste (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200105 WHERE protocolo = 'PROT20230009';

-- Ficha para Paciente 10 (Francisco da Silva - Prioritário) - UBS Setor Norte (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200113 WHERE protocolo = 'PROT20230010';

-- Ficha para Paciente 11 (Fernanda Oliveira) - UBS Setor Marista (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200121 WHERE protocolo = 'PROT20230011';

-- Ficha para Paciente 12 (Ricardo Nunes) - UBS Setor Pedro Ludovico (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200130 WHERE protocolo = 'PROT20230012';

-- Ficha para Paciente 13 (Antônia Bezerra) - UBS Setor Coimbra (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200148 WHERE protocolo = 'PROT20230013';

-- Ficha para Paciente 14 (Camila Rocha - Gestante) - Maternidade Nascer Cidadão (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200156 WHERE protocolo = 'PROT20230014';

-- Ficha para Paciente 15 (Lucas Mendes - Adolescente) - UBS Setor Nova Vila (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200164 WHERE protocolo = 'PROT20230015';

-- Ficha para Paciente 16 (Sophie Martin - Estrangeira) - UBS Setor Bela Vista (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200172 WHERE protocolo = 'PROT20230016';

-- Ficha para Paciente 17 (Jaci Tupinambá - Indígena) - Polo Base de Saúde Indígena (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200180 WHERE protocolo = 'PROT20230017';

-- Ficha para Paciente 18 (Mauro Andrade - Hipertenso) - UBS Setor Criméia (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200199 WHERE protocolo = 'PROT20230018';

-- Ficha para Paciente 19 (Lúcia Pereira - Diabética) - UBS Setor Recanto do Bosque (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200202 WHERE protocolo = 'PROT20230019';

-- Ficha para Paciente 20 (Dona Maria - Prioritária) - UBS Setor Garavelo (Goiânia/GO)
UPDATE ficha SET unidade_id = 5200210 WHERE protocolo = 'PROT20230020';

INSERT INTO ficha (
    protocolo, data_coleta, responsavel_coleta, motivo_exame, observacoes, responsavel_id, paciente_id, unidade_id
) VALUES (
    'PROT20230089', '2025-01-10 08:30:00', 'Enfermeira Maria Silva', 'Rotina', 'Paciente assintomática', 
    (SELECT id FROM medico LIMIT 1), 
    (SELECT id FROM paciente WHERE cpf = '12345678901'),
    5200210
);

INSERT INTO exame (
    inspecao_colo, sinais_dst, observacoes, paciente_id, ficha_id
) VALUES (
    'Alterado', 'Não', 'Colo sem alterações',
    (SELECT id FROM paciente WHERE cpf = '12345678901'),
    (SELECT id FROM ficha WHERE protocolo = 'PROT20230089')
);