# s5n-auth-service

<p align="left">
  <img src="https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white" alt="Go 1.24">
  <img src="https://img.shields.io/badge/Redis-DC382D?logo=redis&logoColor=white" alt="Redis">
  <img src="https://img.shields.io/badge/PostgreSQL-316192?logo=postgresql&logoColor=white" alt="PostgreSQL">
  <img src="https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white" alt="Docker">
  <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="MIT License">
</p>

Este é um **projeto de autenticação desenvolvido em Go** para **estudo de conceitos** como autenticação com JWT, hashing de senhas com Bcrypt, versionamento de banco com `golang-migrate`, roteamento com Gorilla Mux e envio de emails via SMTP.

### **O objetivo foi praticar e entender:**
- Estrutura de projetos em Go
- Criação de APIs REST em Go
- Segurança com JWT e Bcrypt em Go
- Geração de IDs randômicos com Nano ID
- Versionamento de banco de dados
- Dockerização de aplicações

> **Nota:** Este projeto é para fins de estudo e não é recomendado para uso direto em produção sem revisões de segurança e ajustes adicionais.

---


## Tecnologias utilizadas
- **Go 1.24**
- **Gorilla Mux** (roteador HTTP)
- **PostgreSQL**
- **JWT** (para geração e validação de tokens de autenticação)
- **Bcrypt** (para hashing de senhas)
- **Nano ID** (para geração de IDs randômicos, únicos e compactos)
- **golang-migrate** (para versionamento e execução de migrações no banco de dados)
- **Docker** (para containerização dos serviços, incluindo os bancos)
- **net/smtp** (biblioteca nativa do Go para envio de emails via SMTP)
> - **Senhas** de usuários são hasheadas com **Bcrypt** e armazenadas de forma segura em um banco de dados isolado.

## Como Executar
1. **Clonar o repositório**
    ```bash
    git clone https://github.com/lucasschilin/s5n-auth-service.git
    cd s5n-auth-service
    ```

2. **Criar aquivo .env**
    *Usar como base o arquivo .env.example*

3. **Subir os serviços com Docker**
    ```bash
    make docker-compose-up
    ```
4. **Rodar *migrations***
    Para o banco de autenticação:
    ```bash
    make migrate-auth-up
    ```
    Para o banco de usuários:
    ```bash
    make migrate-users-up
    ```
5. **Rodar o serviço**
    ```bash
    make run
    ```


## Endpoints
**GET '/'** – Retorna se o serviço está online e a API está acessível.  

**POST '/auth/signup'** – Cria um novo usuário no sistema.  

**POST '/auth/login'** – Realiza login e retorna um token JWT válido.  

**POST '/auth/refresh'** – Gera um novo token de acesso usando um token de refresh válido.  

**POST '/auth/forgot-password'** – Envia um email com instruções para redefinição de senha.  

**POST '/auth/reset-password'** – Redefine a senha do usuário a partir de um token de redefinição recebido por email.  

**GET '/auth/validate'** – Valida se o token JWT informado é válido e retorna informações do usuário autenticado.  

## Licença
Este projeto está licenciado sob os termos da [MIT License](LICENSE).