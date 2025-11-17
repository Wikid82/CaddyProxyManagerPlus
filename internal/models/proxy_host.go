package models

import (
	"time"
)

// ProxyHost represents a proxy configuration
type ProxyHost struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	DomainNames string    `json:"domain_names" gorm:"not null"` // Comma-separated list
	Scheme      string    `json:"scheme" gorm:"default:http"`
	ForwardHost string    `json:"forward_host" gorm:"not null"`
	ForwardPort int       `json:"forward_port" gorm:"not null"`
	
	// TLS Settings
	SSLEnabled         bool   `json:"ssl_enabled" gorm:"default:false"`
	SSLForced          bool   `json:"ssl_forced" gorm:"default:false"`
	HTTPSPort          int    `json:"https_port" gorm:"default:443"`
	DNSChallenge       bool   `json:"dns_challenge" gorm:"default:false"`
	DNSProvider        string `json:"dns_provider"`
	DNSCredentials     string `json:"dns_credentials"` // Encrypted JSON
	
	// Authentication
	BasicAuthEnabled   bool   `json:"basic_auth_enabled" gorm:"default:false"`
	BasicAuthUsername  string `json:"basic_auth_username"`
	BasicAuthPassword  string `json:"basic_auth_password"` // Hashed
	ForwardAuthEnabled bool   `json:"forward_auth_enabled" gorm:"default:false"`
	ForwardAuthURL     string `json:"forward_auth_url"`
	ForwardAuthType    string `json:"forward_auth_type"` // authelia, authentik, pomerium, custom
	
	// Access Control
	IPWhitelist        string `json:"ip_whitelist"` // Comma-separated CIDR list
	IPBlacklist        string `json:"ip_blacklist"` // Comma-separated CIDR list
	GeoBlockEnabled    bool   `json:"geo_block_enabled" gorm:"default:false"`
	GeoAllowedCountries string `json:"geo_allowed_countries"` // Comma-separated ISO codes
	LocalOnlyEnabled   bool   `json:"local_only_enabled" gorm:"default:false"`
	
	// Security Features
	WAFEnabled         bool   `json:"waf_enabled" gorm:"default:false"`
	WAFRuleSet         string `json:"waf_rule_set" gorm:"default:owasp-crs"`
	RateLimitEnabled   bool   `json:"rate_limit_enabled" gorm:"default:false"`
	RateLimitPreset    string `json:"rate_limit_preset"` // login, api, standard, custom
	RateLimitCustom    string `json:"rate_limit_custom"` // JSON configuration
	
	// Security Headers
	HSTSEnabled        bool   `json:"hsts_enabled" gorm:"default:true"`
	HSTSMaxAge         int    `json:"hsts_max_age" gorm:"default:31536000"`
	HSTSPreload        bool   `json:"hsts_preload" gorm:"default:false"`
	CSPEnabled         bool   `json:"csp_enabled" gorm:"default:false"`
	CSPDirective       string `json:"csp_directive"`
	XFrameOptions      string `json:"x_frame_options" gorm:"default:SAMEORIGIN"`
	ReferrerPolicy     string `json:"referrer_policy" gorm:"default:strict-origin-when-cross-origin"`
	
	// mTLS
	MTLSEnabled        bool   `json:"mtls_enabled" gorm:"default:false"`
	MTLSClientCA       string `json:"mtls_client_ca"` // PEM encoded CA certificate
	
	// CrowdSec
	CrowdSecEnabled    bool   `json:"crowdsec_enabled" gorm:"default:true"`
	
	// Advanced
	CustomCaddyConfig  string    `json:"custom_caddy_config"` // Custom Caddy directives
	Enabled            bool      `json:"enabled" gorm:"default:true"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// TableName specifies the table name for ProxyHost
func (ProxyHost) TableName() string {
	return "proxy_hosts"
}
