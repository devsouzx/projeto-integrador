# SobreVidas

O **SobreVidas** é um sistema desenvolvido por estudantes do curso de Engenharia de Software da Universidade Federal de Goiás (UFG), com o objetivo de oferecer uma solução tecnológica eficiente para profissionais da saúde, atuando nos contextos da atenção básica e especializada. A plataforma facilita o rastreamento de casos, o acompanhamento clínico contínuo e o armazenamento seguro e estruturado de informações médicas, promovendo uma abordagem integrada para a prevenção, diagnóstico precoce e tratamento do câncer do colo do útero. Além disso, o sistema contribui para otimizar os fluxos de trabalho das unidades de saúde e aumentar a taxa de adesão das mulheres aos exames preventivos, ampliando o acesso a cuidados essenciais.

## 🎯 Público-Alvo

- Profissionais da saúde (médicos, enfermeiros, técnicos)
- Unidades Básicas de Saúde (UBS)
- Clínicas, hospitais e laboratórios
- Mulheres entre 25 e 64 anos

## ✅ Funcionalidades Implementadas

- Login de pacientes
- Login de profissionais da saúde
- Cadastro de pacientes
- Recuperação de senha e envio de e-mail de ativação de conta
- Áreas do profissionais de saúde

## 🛠️ Tecnologias e Dependências

- **Go** (versão 1.24.1)
- **HTML e CSS** (frontend)
- **Gin-Gonic** (framework web)
- **go-playground/validator** (validação de dados)
- **golang-jwt/jwt** (autenticação com JWT)
- **godotenv** (configuração via `.env`)
- **gomail** (envio de e-mails)
- **PostgreSQL** (banco de dados)
- Outras dependências estão listadas no `go.mod`

## 🚀 Como Executar o Projeto Localmente

### Pré-requisitos

- Go 1.24.1 ou superior
- PostgreSQL
- Git

### Passos

1. **Clone o repositório**
   ```bash
   git clone https://github.com/devsouzx/projeto-integrador.git
   cd projeto-integrador
   ```

2. **Crie um arquivo `.env`** com as variáveis necessárias:
   ```
   DB_HOST=localhost
   DB_USERNAME=seu_usuario
   DB_PASSWORD=sua_senha
   DB_NAME=sobrevidas
   DB_PORT=5432
   
   JWT_SECRET_KEY=crie_uma_secret_key
   
   SMTP_PORT=587
   SMTP_HOST=smtp.gmail.com
   SMTP_USER=seu_email
   SMTP_PASS=sua_senha
   SMTP_FROM=seu_email
   ```

3. **Instale as dependências**
   ```bash
   go mod tidy
   ```

4. **Inicialize o banco de dados** com as tabelas esperadas.

5. **Execute o servidor**
   ```bash
   go run main.go
   ```

6. **Acesse no navegador**
   ```
   http://localhost:8080
   ```

## 📁 Estrutura do Projeto

```
projeto-integrador/
├── main.go               # Arquivo principal
├── go.mod / go.sum       # Dependências
├── src/                  # Código backend (controladores, rotas, etc.)
  └── static/               # HTML, CSS e Javascript
```

## 👨‍💻 Contribuidores

- João Emanuel Marinho Sousa
- Henrique da Rocha Lima    
- Anderson Ramos Dias  
- Lucas Barretos Dias

## 📄 Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo `LICENSE` para mais detalhes.
