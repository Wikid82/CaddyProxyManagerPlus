# CaddyProxyManager+

A modern web UI for Caddy Server with advanced security features, designed for home-lab and self-hosted environments.

## Features

### 🔐 Authentication & Access Control
- **Forward Auth Integration**: Pre-configured templates for Authelia, Authentik, and Pomerium
- **HTTP Basic Authentication**: Simple username/password protection
- **IP-based Access Control**: Geo-blocking and local network restrictions
- **OAuth/OIDC Server**: Built-in authentication server using caddy-security

### 🛡️ Threat Protection & Content Filtering
- **Web Application Firewall**: Coraza WAF with OWASP Core Rule Set
- **Rate Limiting**: Pre-configured presets for different service types
- **Security Headers**: Automatic HSTS, CSP, X-Frame-Options, and more
- **CrowdSec Integration**: Modern bot and scanner protection

### 🚦 Traffic & TLS Management
- **Automatic HTTPS**: Set it and forget it with Let's Encrypt
- **DNS Challenge Support**: For internal servers and wildcard certificates
- **mTLS Support**: Client certificate authentication for zero-trust setups
- **HSTS Preload**: Easy configuration for enhanced security

### 📊 Monitoring & Logging
- **CrowdSec Dashboard**: View banned IPs and decisions
- **GoAccess Integration**: Beautiful visual analytics
- **Live Log Viewer**: Real-time access and error logs

## Quick Start

### Prerequisites
- Go 1.21 or later
- Caddy 2.7 or later
- SQLite (included)

### Installation

```bash
# Clone the repository
git clone https://github.com/Wikid82/CaddyProxyManagerPlus.git
cd CaddyProxyManagerPlus

# Build the application
go build -o caddyproxymanager ./cmd/caddyproxymanager

# Run the application
./caddyproxymanager
```

The web UI will be available at `http://localhost:8080`

### Docker

```bash
docker run -d \
  --name caddyproxymanager \
  -p 8080:8080 \
  -p 80:80 \
  -p 443:443 \
  -v caddy_data:/data \
  -v caddy_config:/config \
  wikid82/caddyproxymanager:latest
```

## Configuration

CaddyProxyManager+ stores its configuration in an SQLite database. The first time you run the application, you'll be prompted to create an admin account.

## Security Features in Detail

### Forward Authentication
Enable SSO for any proxied service with a single toggle. Pre-configured templates make it easy to integrate with:
- Authelia
- Authentik
- Pomerium

### Web Application Firewall
One-click WAF protection using Coraza with the OWASP Core Rule Set protects against:
- SQL Injection
- Cross-Site Scripting (XSS)
- Remote Code Execution
- Path Traversal
- And many more...

### Rate Limiting
Prevent brute-force attacks with easy presets:
- Login pages: 5 requests/minute
- Standard API: 60 requests/minute
- Custom configurations available

### DNS Challenge Providers
Supported providers for automatic wildcard certificates:
- Cloudflare
- GoDaddy
- Namecheap
- Route53
- And many more...

## Architecture

CaddyProxyManager+ is built with:
- **Backend**: Go with Gin web framework
- **Frontend**: HTML, CSS, JavaScript (no build step required)
- **Database**: SQLite for simple deployment
- **Proxy**: Caddy Server with dynamic configuration

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details

## Acknowledgments

- Inspired by [Nginx Proxy Manager](https://nginxproxymanager.com/)
- Built on [Caddy Server](https://caddyserver.com/)
- Security powered by [CrowdSec](https://www.crowdsec.net/), [Coraza](https://coraza.io/), and more
