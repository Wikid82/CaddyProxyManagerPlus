package caddy

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Wikid82/CaddyProxyManagerPlus/internal/models"
)

// GenerateConfig generates a Caddy configuration from proxy hosts
func GenerateConfig(hosts []models.ProxyHost) (map[string]interface{}, error) {
	config := map[string]interface{}{
		"apps": map[string]interface{}{
			"http": map[string]interface{}{
				"servers": map[string]interface{}{
					"srv0": generateServer(hosts),
				},
			},
		},
	}

	return config, nil
}

// generateServer creates the server configuration
func generateServer(hosts []models.ProxyHost) map[string]interface{} {
	routes := []map[string]interface{}{}

	for _, host := range hosts {
		if !host.Enabled {
			continue
		}

		route := generateRoute(host)
		routes = append(routes, route)
	}

	server := map[string]interface{}{
		"listen": []string{":80", ":443"},
		"routes": routes,
	}

	return server
}

// generateRoute creates a route configuration for a proxy host
func generateRoute(host models.ProxyHost) map[string]interface{} {
	domains := strings.Split(host.DomainNames, ",")
	for i, d := range domains {
		domains[i] = strings.TrimSpace(d)
	}

	route := map[string]interface{}{
		"match": []map[string]interface{}{
			{
				"host": domains,
			},
		},
		"handle": []map[string]interface{}{},
	}

	handlers := []map[string]interface{}{}

	// Add CrowdSec handler if enabled
	if host.CrowdSecEnabled {
		handlers = append(handlers, map[string]interface{}{
			"handler": "crowdsec",
		})
	}

	// Add IP ACL handlers
	if host.LocalOnlyEnabled {
		handlers = append(handlers, generateLocalOnlyHandler())
	}

	if host.IPWhitelist != "" || host.IPBlacklist != "" {
		handlers = append(handlers, generateIPACLHandler(host))
	}

	// Add geo-blocking handler
	if host.GeoBlockEnabled && host.GeoAllowedCountries != "" {
		handlers = append(handlers, generateGeoBlockHandler(host))
	}

	// Add basic auth handler
	if host.BasicAuthEnabled {
		handlers = append(handlers, generateBasicAuthHandler(host))
	}

	// Add forward auth handler
	if host.ForwardAuthEnabled {
		handlers = append(handlers, generateForwardAuthHandler(host))
	}

	// Add rate limiting handler
	if host.RateLimitEnabled {
		handlers = append(handlers, generateRateLimitHandler(host))
	}

	// Add WAF handler
	if host.WAFEnabled {
		handlers = append(handlers, generateWAFHandler(host))
	}

	// Add security headers handler
	handlers = append(handlers, generateSecurityHeadersHandler(host))

	// Add reverse proxy handler
	handlers = append(handlers, generateReverseProxyHandler(host))

	route["handle"] = handlers

	// Add TLS configuration if SSL is enabled
	if host.SSLEnabled {
		route["terminal"] = true
	}

	return route
}

// generateLocalOnlyHandler creates a handler that only allows RFC1918 IPs
func generateLocalOnlyHandler() map[string]interface{} {
	return map[string]interface{}{
		"handler": "vars",
		"root": map[string]interface{}{
			"trusted_proxy": []string{
				"10.0.0.0/8",
				"172.16.0.0/12",
				"192.168.0.0/16",
			},
		},
	}
}

// generateIPACLHandler creates an IP whitelist/blacklist handler
func generateIPACLHandler(host models.ProxyHost) map[string]interface{} {
	acl := map[string]interface{}{
		"handler": "acl",
	}

	if host.IPWhitelist != "" {
		whitelist := strings.Split(host.IPWhitelist, ",")
		for i, ip := range whitelist {
			whitelist[i] = strings.TrimSpace(ip)
		}
		acl["allow"] = whitelist
	}

	if host.IPBlacklist != "" {
		blacklist := strings.Split(host.IPBlacklist, ",")
		for i, ip := range blacklist {
			blacklist[i] = strings.TrimSpace(ip)
		}
		acl["deny"] = blacklist
	}

	return acl
}

// generateGeoBlockHandler creates a geo-blocking handler
func generateGeoBlockHandler(host models.ProxyHost) map[string]interface{} {
	countries := strings.Split(host.GeoAllowedCountries, ",")
	for i, c := range countries {
		countries[i] = strings.TrimSpace(strings.ToUpper(c))
	}

	return map[string]interface{}{
		"handler": "geoip",
		"allow":   countries,
	}
}

// generateBasicAuthHandler creates a basic authentication handler
func generateBasicAuthHandler(host models.ProxyHost) map[string]interface{} {
	return map[string]interface{}{
		"handler": "authentication",
		"providers": map[string]interface{}{
			"http_basic": map[string]interface{}{
				"realm": "Restricted",
				"accounts": []map[string]interface{}{
					{
						"username": host.BasicAuthUsername,
						"password": host.BasicAuthPassword,
					},
				},
			},
		},
	}
}

// generateForwardAuthHandler creates a forward authentication handler
func generateForwardAuthHandler(host models.ProxyHost) map[string]interface{} {
	config := map[string]interface{}{
		"handler": "forward_auth",
		"uri":     host.ForwardAuthURL,
	}

	// Add provider-specific configuration
	switch host.ForwardAuthType {
	case "authelia":
		config["copy_headers"] = []string{
			"Remote-User",
			"Remote-Groups",
			"Remote-Name",
			"Remote-Email",
		}
	case "authentik":
		config["copy_headers"] = []string{
			"X-authentik-username",
			"X-authentik-groups",
			"X-authentik-email",
			"X-authentik-name",
		}
	case "pomerium":
		config["copy_headers"] = []string{
			"X-Pomerium-Jwt-Assertion",
			"X-Pomerium-Claim-Email",
			"X-Pomerium-Claim-Groups",
		}
	}

	return config
}

// generateRateLimitHandler creates a rate limiting handler
func generateRateLimitHandler(host models.ProxyHost) map[string]interface{} {
	var rate, window string

	switch host.RateLimitPreset {
	case "login":
		rate = "5"
		window = "1m"
	case "api":
		rate = "60"
		window = "1m"
	case "standard":
		rate = "100"
		window = "1m"
	case "custom":
		if host.RateLimitCustom != "" {
			var custom map[string]interface{}
			if err := json.Unmarshal([]byte(host.RateLimitCustom), &custom); err == nil {
				return map[string]interface{}{
					"handler": "rate_limit",
					"rate":    custom["rate"],
					"window":  custom["window"],
				}
			}
		}
		rate = "100"
		window = "1m"
	default:
		rate = "100"
		window = "1m"
	}

	return map[string]interface{}{
		"handler": "rate_limit",
		"rate":    rate,
		"window":  window,
	}
}

// generateWAFHandler creates a WAF handler using Coraza
func generateWAFHandler(host models.ProxyHost) map[string]interface{} {
	return map[string]interface{}{
		"handler": "coraza",
		"rule_set": host.WAFRuleSet,
	}
}

// generateSecurityHeadersHandler creates a security headers handler
func generateSecurityHeadersHandler(host models.ProxyHost) map[string]interface{} {
	headers := map[string][]string{}

	// HSTS
	if host.HSTSEnabled {
		hstsValue := fmt.Sprintf("max-age=%d", host.HSTSMaxAge)
		if host.HSTSPreload {
			hstsValue += "; includeSubDomains; preload"
		}
		headers["Strict-Transport-Security"] = []string{hstsValue}
	}

	// CSP
	if host.CSPEnabled && host.CSPDirective != "" {
		headers["Content-Security-Policy"] = []string{host.CSPDirective}
	}

	// X-Frame-Options
	if host.XFrameOptions != "" {
		headers["X-Frame-Options"] = []string{host.XFrameOptions}
	}

	// Referrer-Policy
	if host.ReferrerPolicy != "" {
		headers["Referrer-Policy"] = []string{host.ReferrerPolicy}
	}

	// Additional security headers
	headers["X-Content-Type-Options"] = []string{"nosniff"}
	headers["X-XSS-Protection"] = []string{"1; mode=block"}

	return map[string]interface{}{
		"handler": "headers",
		"response": map[string]interface{}{
			"set": headers,
		},
	}
}

// generateReverseProxyHandler creates a reverse proxy handler
func generateReverseProxyHandler(host models.ProxyHost) map[string]interface{} {
	proxy := map[string]interface{}{
		"handler": "reverse_proxy",
		"upstreams": []map[string]interface{}{
			{
				"dial": fmt.Sprintf("%s:%d", host.ForwardHost, host.ForwardPort),
			},
		},
	}

	// Add mTLS configuration if enabled
	if host.MTLSEnabled && host.MTLSClientCA != "" {
		proxy["transport"] = map[string]interface{}{
			"protocol": "http",
			"tls": map[string]interface{}{
				"client_certificate_authorities": []string{host.MTLSClientCA},
			},
		}
	}

	return proxy
}
