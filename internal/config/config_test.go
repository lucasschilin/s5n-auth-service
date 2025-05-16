package config_test

import (
	"os"
	"testing"

	"github.com/lucasschilin/schily-users-api/internal/config"
)

func TestLoad(t *testing.T) {
	apiHostExpected := os.Getenv("API_HOST")
	apiPortExpected := os.Getenv("API_PORT")

	os.Setenv("API_HOST", apiHostExpected)
	os.Setenv("API_PORT", apiPortExpected)
	defer os.Clearenv()

	config := config.Load()

	// API
	if config.API.Host != apiHostExpected {
		t.Errorf("Expected %v, got %v", apiHostExpected, config.API.Host)
	}
	if config.API.Port != apiPortExpected {
		t.Errorf("Expected %v, got %v", apiPortExpected, config.API.Port)
	}

}
