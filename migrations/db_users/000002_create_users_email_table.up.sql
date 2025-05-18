-- Cria ou substitui a função que atualiza automaticamente o campo 'updated_at' com o timestamp atual
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Cria a tabela 'users'
CREATE TABLE users (
    "id" VARCHAR(21) NOT NULL,
    "username" VARCHAR(24) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMPTZ,
    PRIMARY KEY ("id")
);
COMMENT ON COLUMN users.id IS 'Nano ID';

-- Cria a trigger 'set_updated_at' que chama a função para atualizar o campo 'updated_at'
CREATE TRIGGER set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Cria a tabela 'users_emails'
CREATE TABLE users_emails (
    "id" VARCHAR(21) NOT NULL,
    "user" VARCHAR(21) NOT NULL,
    "address" VARCHAR(254) NOT NULL,
    "verify_token" VARCHAR(50),
    "verified_at" TIMESTAMPTZ,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMPTZ,
    PRIMARY KEY ("id")
);

COMMENT ON COLUMN users_emails.id IS 'Nano ID';

ALTER TABLE users_emails
ADD CONSTRAINT FK_users_TO_users_emails FOREIGN KEY ("user") REFERENCES users ("id") ON DELETE CASCADE;

-- Cria uma trigger que atualiza automaticamente o campo 'updated_at' da tabela 'users_emails' antes de cada UPDATE
CREATE TRIGGER set_updated_at
BEFORE UPDATE ON users_emails
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();