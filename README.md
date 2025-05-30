# s5n-auth-service
**S5N Auth Service** é o microserviço de autenticação e gerenciamento de usuários do projeto **S5N**.

## 🧰 Tecnologias utilizadas
- **Go 1.24**
- **Gorilla Mux** (roteador HTTP)
- **PostgreSQL**
- **JWT** (para geração e validação de tokens de autenticação)
- **Bcrypt** (para hashing de senhas)
- **Nano ID** (para geração de IDs randômicos, únicos e compactos)
- **golang-migrate** (para versionamento e execução de migrações no banco de dados)
- **Docker** (para containerização dos serviços, incluindo os bancos)
> - **Senhas** de usuários são hasheadas com **Bcrypt** e armazenadas de forma segura em um banco de dados isolado.
