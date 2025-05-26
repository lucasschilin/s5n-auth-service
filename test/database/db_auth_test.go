package database_test

import (
	"context"
	"testing"

	"github.com/lucasschilin/schily-users-api/test/database"
)

func TestSetupTestConnectDBAuth(t *testing.T) {
	ctx := context.Background()

	// Setup do banco
	testDBAuth, err := database.SetupTestConnectDBAuth(ctx)
	if err != nil {
		t.Fatalf("Erro ao subir o banco 'db_auth': %v", err)
	}
	defer testDBAuth.Terminate()

	// Testar se a tabela "users" existe
	_, err = testDBAuth.DB.Query("SELECT * FROM schema_migrations")
	if err != nil {
		t.Fatalf("Erro na query em 'db_auth', provavelmente a tabela não existe: %v", err)
	}

	t.Log("Banco de teste 'db_auth' está funcionando e a tabela 'schema_migrations' existe!")
}
