# schily-users-api
**Schily Users API** é o microserviço de autenticação e manipulação de usuários do **Schily**, responsável pelo cadastro, operações de autenticação e manipulação de usuário.  
Ele suporta autenticação via **OAuth 2.0** (Google e Facebook), **autenticação de dois fatores (2FA)** via Google Authenticator e armazena **tokens e senhas de forma segura** com criptografia AES e hash Bcrypt em banco paralelo.

Este serviço é consumido pelo frontend do módulo de Autenticação e Usuário e outros microserviços do **Schily**.


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
> - **Tokens OAuth 2.0** serão armazenados em um banco de dados separado, garantindo maior segurança e controle.
> - **Senhas** de usuários serão hasheadas utilizando **Bcrypt** e armazenadas de forma segura no banco de dados principal.

## 🔗 Interface Web
Este repositório é apenas o backend. As interfaces web que farão uso desta aplicação estão em outro repositório.
