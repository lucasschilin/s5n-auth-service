# schily-auth-api
**Schily Auth API** é o microserviço de autenticação do **Schily**, responsável pelo cadastro, autenticação e validação de usuários.  
Ele suporta autenticação via **OAuth 2.0** (Google e Facebook), **autenticação de dois fatores (2FA)** via Google Authenticator e armazena **tokens e senhas de forma segura** com criptografia AES e hash Bcrypt.

Este serviço é consumido pelo frontend do módulo de Autenticação e outros microserviços do **Schily**.
Toda autenticação é feita via JWT.


## 🧰 Tecnologias utilizadas
- Go 1.24
- Gorilla Mux (roteador HTTP)
- PostgreSQL
- JWT (para geração e validação de tokens de autenticação)
- OAuth 2.0 (para login com Google e Facebook)
- Google Authenticator (para 2FA)
- AES (para criptografia de tokens)
- Bcrypt (para hashing de senhas)
- Docker (opcional, para containerização)
> **Importante:**  
> - **Tokens OAuth 2.0** serão armazenados em um banco de dados separado, garantindo maior segurança e controle.
> - **Senhas** de usuários serão hasheadas utilizando **Bcrypt** e armazenadas de forma segura no banco de dados principal.

## 🔗 Interface Web
Este repositório é apenas o backend. A interface web está, ou estará, em [schily-auth](https://github.com/lucasschilin/schily-auth).
