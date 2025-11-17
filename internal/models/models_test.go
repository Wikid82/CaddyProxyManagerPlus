package models

import (
	"testing"
)

func TestProxyHostTableName(t *testing.T) {
	host := ProxyHost{}
	if host.TableName() != "proxy_hosts" {
		t.Errorf("Expected table name to be proxy_hosts, got %s", host.TableName())
	}
}

func TestUserTableName(t *testing.T) {
	user := User{}
	if user.TableName() != "users" {
		t.Errorf("Expected table name to be users, got %s", user.TableName())
	}
}

func TestSettingsTableName(t *testing.T) {
	settings := Settings{}
	if settings.TableName() != "settings" {
		t.Errorf("Expected table name to be settings, got %s", settings.TableName())
	}
}

func TestAccessLogTableName(t *testing.T) {
	log := AccessLog{}
	if log.TableName() != "access_logs" {
		t.Errorf("Expected table name to be access_logs, got %s", log.TableName())
	}
}

func TestCrowdSecDecisionTableName(t *testing.T) {
	decision := CrowdSecDecision{}
	if decision.TableName() != "crowdsec_decisions" {
		t.Errorf("Expected table name to be crowdsec_decisions, got %s", decision.TableName())
	}
}

func TestProxyHostDefaults(t *testing.T) {
	host := ProxyHost{}
	
	// Test that zero values are appropriate
	if host.ID != 0 {
		t.Error("New ProxyHost should have ID 0")
	}
	
	if host.Enabled {
		t.Error("New ProxyHost should not be enabled by default")
	}
	
	if host.SSLEnabled {
		t.Error("New ProxyHost should not have SSL enabled by default")
	}
}

func TestUserDefaults(t *testing.T) {
	user := User{}
	
	if user.ID != 0 {
		t.Error("New User should have ID 0")
	}
	
	if user.IsAdmin {
		t.Error("New User should not be admin by default")
	}
	
	if user.Enabled {
		t.Error("New User should not be enabled by default")
	}
}
