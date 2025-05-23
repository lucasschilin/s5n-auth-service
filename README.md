# schily-users-api
**Schily Users API** é o microserviço de autenticação e manipulação de usuários da aplicação **Schily**.

## 🧰 Tecnologias utilizadas
- **Go 1.24**
- **Gorilla Mux** (roteador HTTP)
- **PostgreSQL 17**
- **JWT** (para geração e validação de tokens de autenticação)
- **OAuth 2.0** (para login com Google e Facebook)
- **Google Authenticator** (para 2FA)
- **AES** (para criptografia de tokens)
- **Bcrypt** (para hashing de senhas)
- **Nano ID** (para geração de IDs randômicos, únicos e compactos)
- **golang-migrate 4.18** (para versionamento e execução de migrações no banco de dados)
- **Docker** (para containerização dos serviços, incluindo os bancos)
> - **Tokens OAuth 2.0** serão armazenados em um banco de dados separado, garantindo maior segurança e controle.
> - **Senhas** de usuários serão hasheadas utilizando **Bcrypt** e armazenadas de forma segura no banco de dados principal.

## 🔗 Interface Web
Este repositório é apenas o backend. As interfaces web que farão uso desta aplicação estão em outro repositório.
