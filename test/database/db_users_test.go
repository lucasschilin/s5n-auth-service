package database_test

import (
	"context"
	"testing"

	"github.com/lucasschilin/schily-users-api/test/database"
)

func TestSetupTestConnectDBUsers(t *testing.T) {
	ctx := context.Background()

	// Setup do banco
	testDBUsers, err := database.SetupTestConnectDBAuth(ctx)
	if err != nil {
		t.Fatalf("Erro ao subir o banco 'db_users': %v", err)
	}
	defer testDBUsers.Terminate()

	// Testar se a tabela "users" existe
	_, err = testDBUsers.DB.Query("SELECT * FROM schema_migrations")
	if err != nil {
		t.Fatalf("Erro na query em 'db_users', provavelmente a tabela não existe: %v", err)
	}

	t.Log("Banco de teste 'db_users' está funcionando e a tabela 'schema_migrations' existe!")
}
