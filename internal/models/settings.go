package models

import (
	"time"
)

// Settings represents global application settings
type Settings struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Key   string `json:"key" gorm:"uniqueIndex;not null"`
	Value string `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for Settings
func (Settings) TableName() string {
	return "settings"
}

// AccessLog represents an access log entry
type AccessLog struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Host       string    `json:"host"`
	Method     string    `json:"method"`
	Path       string    `json:"path"`
	Status     int       `json:"status"`
	IP         string    `json:"ip"`
	UserAgent  string    `json:"user_agent"`
	Country    string    `json:"country"`
	Blocked    bool      `json:"blocked"`
	BlockedBy  string    `json:"blocked_by"` // waf, crowdsec, rate_limit, ip_acl
	CreatedAt  time.Time `json:"created_at"`
}

// TableName specifies the table name for AccessLog
func (AccessLog) TableName() string {
	return "access_logs"
}

// CrowdSecDecision represents a CrowdSec ban decision
type CrowdSecDecision struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	IP        string    `json:"ip" gorm:"index"`
	Reason    string    `json:"reason"`
	Scenario  string    `json:"scenario"`
	Duration  string    `json:"duration"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName specifies the table name for CrowdSecDecision
func (CrowdSecDecision) TableName() string {
	return "crowdsec_decisions"
}
