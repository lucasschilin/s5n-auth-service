-- Remove as triggers
DROP TRIGGER IF EXISTS set_updated_at ON passwords;

-- Remove as tabelas
DROP TABLE IF EXISTS passwords;

-- Remove a função que era usada pelas triggers
DROP FUNCTION IF EXISTS update_updated_at_column();