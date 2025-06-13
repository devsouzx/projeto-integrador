INSERT INTO medico (
  nomecompleto, 
  cpf, 
  email,
  telefone, 
  crm, 
  especialidade, 
  senha, 
  unidade_saude_id
)
VALUES (
  'Dra. Maria Oliveira',
  '12345678901',
  'medico@ufg.br',
  '62981887777',
  '12345GO',
  'Ginecologista',
  crypt('@Teste123', gen_salt('bf')),
  (SELECT id FROM unidade_saude WHERE cnes = '1234567')
);
