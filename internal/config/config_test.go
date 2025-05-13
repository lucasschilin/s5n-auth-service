package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	apiHostExpected := "127.0.0.1"
	apiPortExpected := "8080"

	os.Setenv("API_HOST", apiHostExpected)
	os.Setenv("API_PORT", apiPortExpected)
	defer os.Clearenv()

	config := Load()

	if config.Host != apiHostExpected {
		t.Errorf("Expected %v, got %v", apiHostExpected, config.Host)
	}

	if config.Port != apiPortExpected {
		t.Errorf("Expected %v, got %v", apiPortExpected, config.Port)
	}

}
