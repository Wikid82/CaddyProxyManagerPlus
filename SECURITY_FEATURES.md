# Security Features Documentation

CaddyProxyManager+ implements comprehensive security features designed for home-lab and self-hosted environments.

## 🔐 Authentication & Access Control

### Forward Authentication (SSO)
Enable single sign-on for any proxied service with pre-configured templates:

#### Authelia
- URL: `https://auth.example.com/api/verify?rd=https://service.example.com`
- Automatically copies headers: `Remote-User`, `Remote-Groups`, `Remote-Name`, `Remote-Email`

#### Authentik
- URL: `https://auth.example.com/outpost.goauthentik.io/auth/nginx`
- Automatically copies headers: `X-authentik-username`, `X-authentik-groups`, `X-authentik-email`, `X-authentik-name`

#### Pomerium
- URL: `https://auth.example.com/verify`
- Automatically copies headers: `X-Pomerium-Jwt-Assertion`, `X-Pomerium-Claim-Email`, `X-Pomerium-Claim-Groups`

### HTTP Basic Authentication
Simple username/password protection using bcrypt hashed passwords.

**Usage:**
1. Enable "Basic Authentication" in the Authentication tab
2. Enter username and password
3. Credentials are hashed before storage

### IP-based Access Control

#### Local Network Only (RFC1918)
Restricts access to private IP ranges:
- `10.0.0.0/8`
- `172.16.0.0/12`
- `192.168.0.0/16`

Perfect for admin panels and internal services.

#### IP Whitelist/Blacklist
- **Whitelist**: Allow only specific IPs or CIDR ranges
- **Blacklist**: Deny specific IPs or CIDR ranges
- Format: `192.168.1.0/24, 10.0.0.5/32`

#### Geo-blocking
Restrict access by country using ISO country codes.

**Example:**
- `US, CA, GB` - Only allow United States, Canada, and United Kingdom

## 🛡️ Threat Protection & Content Filtering

### Web Application Firewall (WAF)
Powered by Coraza with the OWASP Core Rule Set (CRS).

**Protection against:**
- SQL Injection
- Cross-Site Scripting (XSS)
- Remote Code Execution (RCE)
- Path Traversal
- Command Injection
- And 100+ other attack patterns

**Usage:**
Enable the "WAF" checkbox in the Security tab.

### Rate Limiting
Prevent brute-force attacks and DoS with preset configurations:

#### Presets:
- **Login Page**: 5 requests per minute
  - Use for login pages, admin panels
- **Standard API**: 60 requests per minute
  - Use for REST APIs, webhooks
- **Standard**: 100 requests per minute
  - Use for general web services
- **Custom**: Define your own rate and window

**How it works:**
Rate limits are applied per client IP address. Exceeding the limit returns a 429 (Too Many Requests) error.

### HTTP Security Headers
Automatically configured headers for enhanced security:

#### Strict-Transport-Security (HSTS)
- Default: `max-age=31536000` (1 year)
- Optional: `includeSubDomains; preload` for HSTS preloading

#### Content-Security-Policy (CSP)
- Configurable per-host
- Example: `default-src 'self'; script-src 'self' 'unsafe-inline'`

#### X-Frame-Options
- Default: `SAMEORIGIN` (prevents clickjacking)
- Options: `DENY`, `SAMEORIGIN`, or allow

#### Referrer-Policy
- Default: `strict-origin-when-cross-origin`
- Controls referrer information sent with requests

#### Additional Headers:
- `X-Content-Type-Options: nosniff`
- `X-XSS-Protection: 1; mode=block`

### CrowdSec Integration
Real-time protection against bots, scanners, and known bad actors.

**Features:**
- Automatic IP blocking based on CrowdSec decisions
- Community-powered threat intelligence
- Integration with CrowdSec Local API
- Dashboard showing active bans

**Setup:**
1. Install CrowdSec on your server
2. Enable "CrowdSec Protection" (enabled by default)
3. Monitor bans in the Monitoring tab

## 🚦 Traffic & TLS Management

### Automatic HTTPS
Let's Encrypt certificates are automatically obtained and renewed.

**Default behavior:**
- HTTP-01 challenge for public servers
- Automatic certificate renewal

### DNS Challenge
For internal servers and wildcard certificates.

**Supported Providers:**
- Cloudflare
- AWS Route53
- GoDaddy
- Namecheap
- And many more...

**Setup:**
1. Enable "Use DNS Challenge"
2. Select your DNS provider
3. Enter API credentials (usually an API token)
4. Wildcard certificates (*.example.com) will be automatically obtained

**Example (Cloudflare):**
```json
{
  "api_token": "your-cloudflare-api-token"
}
```

### mTLS (Mutual TLS)
Client certificate authentication for zero-trust setups.

**Setup:**
1. Enable "Require Client Certificate" in Advanced tab
2. Paste your CA certificate (PEM format)
3. Only clients with certificates signed by your CA can connect

**Use cases:**
- Internal APIs requiring machine-to-machine authentication
- Highly sensitive admin panels
- Compliance requirements

### HSTS Preload
Enable your domain for the HSTS preload list.

**Requirements:**
- HTTPS must be enabled
- HSTS max-age must be at least 31536000 seconds (1 year)
- Enable "HSTS Preload" checkbox

**Note:** Submitting to the preload list is permanent and requires manual removal.

## 📊 Monitoring & Logging

### CrowdSec Dashboard
View real-time ban decisions:
- Banned IP addresses
- Reason for ban
- Scenario that triggered the ban
- Ban duration and expiry

### Access Logs
Real-time monitoring of all proxied traffic:
- Request method and path
- Status codes
- Client IP addresses
- User agents
- Geo-location data
- Security events (WAF blocks, rate limits)

### GoAccess Integration (Planned)
Beautiful visual analytics from Caddy access logs:
- Visitor statistics
- Request patterns
- Top URLs
- Browser and OS breakdown
- Geographic distribution

## Security Best Practices

### For Home-Lab Users:
1. **Use Forward Auth** for services without built-in authentication
2. **Enable WAF** for all publicly-exposed services
3. **Enable CrowdSec** to block automated attacks
4. **Use DNS Challenge** if your server is not directly accessible
5. **Enable Rate Limiting** on login pages

### For Plex/Jellyfin/Arr Suite:
1. **Enable HSTS** to force HTTPS
2. **Use Geo-blocking** to restrict access to your country
3. **Enable WAF** to protect against scanner bots
4. **Consider Forward Auth** for additional security layer
5. **Enable Rate Limiting** to prevent abuse

### For Internal Services:
1. **Use "Local Network Only"** for admin panels
2. **Consider mTLS** for API-to-API communication
3. **Use IP Whitelist** if you have a static IP
4. **Disable public exposure** when not needed

## Advanced Configuration

### Custom Caddy Directives
For advanced users, you can add custom Caddy directives in the "Advanced" tab.

**Example:**
```
# Custom logging
log {
    output file /var/log/caddy/access.log
}

# Custom headers
header {
    Custom-Header "Custom Value"
}
```

## Troubleshooting

### WAF Blocking Legitimate Traffic
If the WAF is too aggressive:
1. Check access logs for blocked requests
2. Adjust CSP or disable specific rules (requires custom config)
3. Consider using "Learning Mode" during initial setup

### Rate Limiting Too Restrictive
If users are getting rate limited:
1. Increase the rate limit preset
2. Use "Custom" preset with higher values
3. Consider whitelisting specific IPs

### Forward Auth Not Working
Common issues:
1. Verify the Forward Auth URL is correct
2. Check that headers are being copied correctly
3. Ensure the SSO provider is accessible from Caddy
4. Check SSO provider logs for errors

### Certificate Issues
If certificates aren't being issued:
1. Verify DNS records are correct
2. For DNS challenge: check API credentials
3. Check Caddy logs for specific errors
4. Ensure ports 80/443 are accessible (for HTTP challenge)

## Security Updates

Always keep CaddyProxyManager+ and Caddy updated to receive the latest security patches.

For security issues, please report to the project maintainers.
