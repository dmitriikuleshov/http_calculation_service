package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tempFile := createTempConfigFile(t)
	defer os.Remove(tempFile.Name()) 

	config := loadAndValidateConfig(t, tempFile.Name())

	assertEqual(t, "test_host", config.Server.Host, "host")
	assertEqual(t, "9999", config.Server.Port, "port")
}

// createTempConfigFile creates a temporary configuration file with test data
func createTempConfigFile(t *testing.T) *os.File {
	tempFile, err := os.CreateTemp("", "config.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}

	configData := `{
		"server": {
			"host": "test_host",
			"port": "9999"
		}
	}`

	if _, err := tempFile.WriteString(configData); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	if err := tempFile.Close(); err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}

	return tempFile
}

// loadAndValidateConfig loads the configuration and validates any unexpected errors
func loadAndValidateConfig(t *testing.T, fileName string) *Config {
	config, err := LoadConfig(fileName)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
	return config
}

// assertEqual compares the expected and actual values and logs an error if they differ
func assertEqual(t *testing.T, expected, actual, field string) {
	if expected != actual {
		t.Errorf("Expected %s '%s', got '%s'", field, expected, actual)
	}
}
