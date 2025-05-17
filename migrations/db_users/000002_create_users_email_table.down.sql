-- Remove as triggers
DROP TRIGGER IF EXISTS set_updated_at ON users_emails;
DROP TRIGGER IF EXISTS set_updated_at ON users;

-- Remove as tabelas
DROP TABLE IF EXISTS users_emails;
DROP TABLE IF EXISTS users;

-- Remove a função que era usada pelas triggers
DROP FUNCTION IF EXISTS update_updated_at_column();