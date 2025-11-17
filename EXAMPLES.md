# Configuration Examples

This document provides real-world examples for configuring CaddyProxyManager+ for common use cases.

## Basic Reverse Proxy

**Use Case:** Simple reverse proxy to an internal service

**Configuration:**
- Domain Names: `app.example.com`
- Scheme: `http`
- Forward Host: `192.168.1.100`
- Forward Port: `8080`
- SSL Enabled: ✅
- Force HTTPS: ✅

This creates a basic proxy with automatic HTTPS.

## Plex Media Server

**Use Case:** Secure Plex with authentication and geo-blocking

**Basic Tab:**
- Domain Names: `plex.example.com`
- Scheme: `http`
- Forward Host: `192.168.1.50`
- Forward Port: `32400`

**SSL/TLS Tab:**
- SSL Enabled: ✅
- Force HTTPS: ✅
- HSTS Enabled: ✅
- HSTS Preload: ✅

**Authentication Tab:**
- Forward Auth Enabled: ✅
- Provider: Authelia
- Forward Auth URL: `https://auth.example.com/api/verify?rd=https://plex.example.com`

**Security Tab:**
- CrowdSec Enabled: ✅
- Geo-blocking Enabled: ✅
- Allowed Countries: `US, CA` (restrict to your countries)
- Rate Limit Enabled: ✅
- Rate Limit Preset: Standard (100 req/min)

**Why:**
- Forward Auth adds SSO protection
- Geo-blocking prevents international bot traffic
- CrowdSec blocks known bad actors
- Rate limiting prevents abuse

## Jellyfin Media Server

**Use Case:** Jellyfin with WAF and rate limiting

**Basic Tab:**
- Domain Names: `jellyfin.example.com`
- Scheme: `http`
- Forward Host: `192.168.1.51`
- Forward Port: `8096`

**SSL/TLS Tab:**
- SSL Enabled: ✅
- Force HTTPS: ✅
- HSTS Enabled: ✅

**Security Tab:**
- WAF Enabled: ✅
- Rate Limit Enabled: ✅
- Rate Limit Preset: Standard
- CrowdSec Enabled: ✅
- Geo-blocking Enabled: ✅
- Allowed Countries: `US, CA, GB`

**Why:**
- WAF protects against scanner bots
- Rate limiting prevents brute force on login
- CrowdSec provides additional bot protection

## Sonarr/Radarr/Prowlarr

**Use Case:** Arr suite with SSO and local-only access for management

### Public API (for external access)
**Basic Tab:**
- Domain Names: `sonarr.example.com`
- Scheme: `http`
- Forward Host: `192.168.1.60`
- Forward Port: `8989`

**SSL/TLS Tab:**
- SSL Enabled: ✅
- Force HTTPS: ✅

**Authentication Tab:**
- Forward Auth Enabled: ✅
- Provider: Authentik
- Forward Auth URL: `https://auth.example.com/outpost.goauthentik.io/auth/nginx`

**Security Tab:**
- WAF Enabled: ✅
- Rate Limit Enabled: ✅
- Rate Limit Preset: API (60 req/min)
- CrowdSec Enabled: ✅

### Admin Panel (internal only)
**Basic Tab:**
- Domain Names: `sonarr-admin.local`
- Scheme: `http`
- Forward Host: `192.168.1.60`
- Forward Port: `8989`

**Security Tab:**
- Local Only Enabled: ✅ (RFC1918 only)
- No authentication needed (protected by network)

**Why:**
- Public API has full security stack
- Admin panel accessible only from local network
- WAF protects against automation bot attacks

## Internal Wiki/Documentation

**Use Case:** Internal documentation with basic auth

**Basic Tab:**
- Domain Names: `wiki.internal.example.com`
- Scheme: `http`
- Forward Host: `192.168.1.70`
- Forward Port: `3000`

**SSL/TLS Tab:**
- SSL Enabled: ✅
- DNS Challenge: ✅
- DNS Provider: Cloudflare
- API Credentials: `{"api_token": "your-token"}`

**Authentication Tab:**
- Basic Auth Enabled: ✅
- Username: `admin`
- Password: `your-secure-password`

**Security Tab:**
- Local Only Enabled: ✅
- HSTS Enabled: ✅

**Why:**
- DNS challenge allows internal-only server to get certificates
- Basic auth is sufficient for internal wiki
- Local-only prevents external access
- Still uses HTTPS for encryption on local network

## Public API with mTLS

**Use Case:** Machine-to-machine API requiring client certificates

**Basic Tab:**
- Domain Names: `api.example.com`
- Scheme: `http`
- Forward Host: `192.168.1.80`
- Forward Port: `8000`

**SSL/TLS Tab:**
- SSL Enabled: ✅
- Force HTTPS: ✅
- HSTS Enabled: ✅
- HSTS Preload: ✅

**Security Tab:**
- WAF Enabled: ✅
- Rate Limit Enabled: ✅
- Rate Limit Preset: API (60 req/min)
- CrowdSec Enabled: ✅

**Advanced Tab:**
- mTLS Enabled: ✅
- Client CA Certificate: 
```
-----BEGIN CERTIFICATE-----
MIIDXTCCAkWgAwIBAgIJAKL0UG+mRnijMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV
[... your CA certificate ...]
-----END CERTIFICATE-----
```

**Why:**
- Only clients with valid certificates can connect
- Perfect for service-to-service authentication
- WAF and rate limiting provide additional protection
- No password-based authentication needed

## WordPress Site

**Use Case:** WordPress with aggressive security

**Basic Tab:**
- Domain Names: `blog.example.com, www.blog.example.com`
- Scheme: `http`
- Forward Host: `192.168.1.90`
- Forward Port: `80`

**SSL/TLS Tab:**
- SSL Enabled: ✅
- Force HTTPS: ✅
- HSTS Enabled: ✅
- HSTS Max Age: `31536000`
- HSTS Preload: ✅

**Security Tab:**
- WAF Enabled: ✅ (critical for WordPress)
- Rate Limit Enabled: ✅
- Custom Rate Limit:
```json
{
  "rate": "20",
  "window": "1m"
}
```
- CrowdSec Enabled: ✅
- Geo-blocking Enabled: ✅
- Allowed Countries: `US, CA, GB, AU` (your target audience)

**Advanced Tab:**
- CSP Directive: `default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval'; style-src 'self' 'unsafe-inline'`
- X-Frame-Options: `SAMEORIGIN`
- Referrer-Policy: `strict-origin-when-cross-origin`

**Why:**
- WordPress is a common target for attacks
- WAF blocks most automated attacks
- Rate limiting prevents brute force on wp-admin
- CrowdSec blocks known WordPress scanners
- Geo-blocking reduces attack surface

## Home Assistant

**Use Case:** Home automation with SSO and local network priority

**Basic Tab:**
- Domain Names: `home.example.com`
- Scheme: `http`
- Forward Host: `192.168.1.100`
- Forward Port: `8123`

**SSL/TLS Tab:**
- SSL Enabled: ✅
- Force HTTPS: ✅
- HSTS Enabled: ✅

**Authentication Tab:**
- Forward Auth Enabled: ✅ (for external access)
- Provider: Authelia
- Forward Auth URL: `https://auth.example.com/api/verify?rd=https://home.example.com`

**Security Tab:**
- IP Whitelist: `192.168.1.0/24` (local network always allowed)
- Rate Limit Enabled: ✅
- Rate Limit Preset: Login (5 req/min)
- CrowdSec Enabled: ✅

**Why:**
- Local network can access directly
- External access requires SSO
- Rate limiting prevents brute force
- CrowdSec blocks automated attacks

## Development Server (Staging)

**Use Case:** Development/staging environment with IP whitelist

**Basic Tab:**
- Domain Names: `dev.example.com`
- Scheme: `http`
- Forward Host: `192.168.1.110`
- Forward Port: `3000`

**SSL/TLS Tab:**
- SSL Enabled: ✅
- Force HTTPS: ✅

**Authentication Tab:**
- Basic Auth Enabled: ✅
- Username: `dev`
- Password: `dev-password`

**Security Tab:**
- IP Whitelist: `203.0.113.0/24, 192.168.1.0/24` (office and local IPs)
- Local Only Enabled: ❌ (need remote access)

**Advanced Tab:**
- Custom Caddy Config:
```
# Disable caching for development
header {
    Cache-Control "no-store, no-cache, must-revalidate"
    -Server
}
```

**Why:**
- IP whitelist restricts to known developers
- Basic auth provides simple protection
- Custom headers prevent caching issues in development

## Multi-Domain Service

**Use Case:** One service, multiple domains (with and without www)

**Basic Tab:**
- Domain Names: `example.com, www.example.com, example.net, www.example.net`
- Scheme: `http`
- Forward Host: `192.168.1.120`
- Forward Port: `8080`

**SSL/TLS Tab:**
- SSL Enabled: ✅
- Force HTTPS: ✅
- HSTS Enabled: ✅

**Security Tab:**
- WAF Enabled: ✅
- Rate Limit Enabled: ✅
- Rate Limit Preset: Standard
- CrowdSec Enabled: ✅

**Why:**
- Multiple domains handled by single config
- All domains get same security posture
- Automatic certificates for all domains

## Configuration Tips

### Rate Limiting Guidelines
- **Login Pages**: 5 requests/minute (prevents brute force)
- **APIs**: 60 requests/minute (balances functionality and protection)
- **Static Content**: 100+ requests/minute (or no limit)
- **Public Forms**: 10 requests/minute (prevents spam)

### When to Use Each Auth Method
- **Forward Auth (SSO)**: Multi-service authentication, user management
- **Basic Auth**: Simple single-service protection, internal tools
- **mTLS**: Machine-to-machine, API authentication, zero-trust
- **No Auth + IP Whitelist**: Internal services, trusted networks

### WAF Best Practices
- Enable for all public-facing services
- Essential for WordPress, Joomla, Drupal
- Important for custom web applications
- May need tuning for false positives

### Geo-blocking Considerations
- Use for services with known geographic audience
- Reduces automated attack traffic by 80%+
- Keep monitoring for legitimate international users
- Consider VPN users might be blocked

### HSTS Configuration
- Always enable for production services
- Use 1-year max-age for production
- Enable preload only if you're committed
- Remember: HSTS preload is permanent

## Monitoring Your Configuration

After setting up, monitor:
1. Access logs for blocked requests
2. CrowdSec dashboard for banned IPs
3. Rate limit hits in logs
4. WAF blocks in security logs
5. Certificate renewal status

## Testing Your Configuration

Before going live:
1. Test SSL certificate is valid
2. Verify rate limiting works (exceed limit)
3. Test authentication (try wrong credentials)
4. Check WAF blocks common attacks (SQL injection attempts)
5. Verify geo-blocking works (use VPN)
6. Test from both internal and external networks

## Common Mistakes to Avoid

1. **Too Restrictive Rate Limiting**: Start conservative, increase if needed
2. **Wrong IP Whitelist Format**: Use CIDR notation (e.g., `192.168.1.0/24`)
3. **Forgetting DNS Challenge for Internal Servers**: HTTP challenge won't work
4. **Not Testing Authentication**: Always test before deploying
5. **Enabling HSTS Preload Too Early**: Test HTTPS thoroughly first

## Getting Help

If you run into issues:
1. Check the security logs
2. Review Caddy logs
3. Verify DNS records
4. Test without security features first
5. Open an issue on GitHub with logs
