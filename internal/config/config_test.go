package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Test default configuration
	cfg := Load()
	
	if cfg.ServerPort == "" {
		t.Error("ServerPort should have a default value")
	}
	
	if cfg.DataPath == "" {
		t.Error("DataPath should have a default value")
	}
	
	if cfg.CaddyAdminURL == "" {
		t.Error("CaddyAdminURL should have a default value")
	}
}

func TestLoadWithEnvVars(t *testing.T) {
	// Set environment variables
	os.Setenv("SERVER_PORT", "9090")
	os.Setenv("DATA_PATH", "/custom/data")
	os.Setenv("CADDY_ADMIN_URL", "http://caddy:2019")
	defer func() {
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("DATA_PATH")
		os.Unsetenv("CADDY_ADMIN_URL")
	}()
	
	cfg := Load()
	
	if cfg.ServerPort != "9090" {
		t.Errorf("Expected ServerPort to be 9090, got %s", cfg.ServerPort)
	}
	
	if cfg.DataPath != "/custom/data" {
		t.Errorf("Expected DataPath to be /custom/data, got %s", cfg.DataPath)
	}
	
	if cfg.CaddyAdminURL != "http://caddy:2019" {
		t.Errorf("Expected CaddyAdminURL to be http://caddy:2019, got %s", cfg.CaddyAdminURL)
	}
}

func TestGetEnv(t *testing.T) {
	// Test default value
	value := getEnv("NONEXISTENT_VAR", "default")
	if value != "default" {
		t.Errorf("Expected default value, got %s", value)
	}
	
	// Test environment variable
	os.Setenv("TEST_VAR", "test_value")
	defer os.Unsetenv("TEST_VAR")
	
	value = getEnv("TEST_VAR", "default")
	if value != "test_value" {
		t.Errorf("Expected test_value, got %s", value)
	}
}
