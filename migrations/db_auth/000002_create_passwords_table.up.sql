-- Cria ou substitui a função que atualiza automaticamente o campo 'updated_at' com o timestamp atual
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Cria a tabela 'passwords'
CREATE TABLE passwords (
    "user" VARCHAR(21) NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("user")
);
COMMENT ON COLUMN passwords.user IS 'table PK and "FK" from db_users.public.users.id';

-- Cria a trigger 'set_updated_at' que chama a função para atualizar o campo 'updated_at'
CREATE TRIGGER set_updated_at
BEFORE UPDATE ON passwords
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();