# Changelog

All notable changes to CaddyProxyManager+ will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial release of CaddyProxyManager+
- Web UI for managing Caddy reverse proxy configurations
- Comprehensive security features for home-lab and self-hosted environments

#### Authentication & Access Control
- Forward Authentication (SSO) integration with pre-configured templates:
  - Authelia
  - Authentik
  - Pomerium
  - Custom providers
- HTTP Basic Authentication support
- IP-based Access Control Lists (whitelist/blacklist)
- Geo-blocking by country (ISO codes)
- Local network only access (RFC1918)
- OAuth/OIDC server integration preparation

#### Threat Protection & Content Filtering
- Web Application Firewall (Coraza WAF) with OWASP Core Rule Set
- Rate limiting with smart presets:
  - Login pages (5 requests/minute)
  - Standard API (60 requests/minute)
  - Standard web (100 requests/minute)
  - Custom configurations
- HTTP Security Headers:
  - Strict-Transport-Security (HSTS) with preload support
  - Content-Security-Policy (CSP)
  - X-Frame-Options
  - Referrer-Policy
  - X-Content-Type-Options
  - X-XSS-Protection
- CrowdSec integration support

#### Traffic & TLS Management
- Automatic HTTPS with Let's Encrypt
- DNS Challenge support for wildcard certificates
- Supported DNS providers:
  - Cloudflare
  - AWS Route53
  - GoDaddy
  - Namecheap
- mTLS (Mutual TLS) client certificate authentication
- HSTS preload configuration
- Force HTTPS redirect option
- Automatic certificate renewal

#### User Interface
- Modern dark theme interface
- Responsive design for mobile and desktop
- Tabbed proxy host configuration:
  - Basic settings
  - SSL/TLS configuration
  - Authentication options
  - Security features
  - Advanced settings
- Visual security badges showing active features
- Real-time host management
- Enable/disable hosts without deletion
- Form validation and error handling

#### Backend Features
- REST API built with Gin framework
- SQLite database with GORM ORM
- Caddy admin API integration
- JWT authentication with middleware
- Modular configuration generator
- Secure credential storage
- Rate limiting per IP address
- User management system

#### Deployment
- Docker support with multi-stage builds
- Docker Compose configuration with Caddy
- Standalone binary (no external dependencies)
- Systemd service file example
- Environment variable configuration
- Volume persistence for data and config

#### Documentation
- Comprehensive README with feature overview
- Security features documentation (SECURITY_FEATURES.md)
- Detailed installation guide (INSTALLATION.md)
- Quick start guide (QUICKSTART.md)
- Real-world configuration examples (EXAMPLES.md)
- Contributing guidelines (CONTRIBUTING.md)
- MIT License

### Technical Details
- Go 1.21+ backend
- SQLite database
- Gin web framework
- JWT authentication
- Caddy 2.7+ admin API integration
- No build step required for frontend
- Single binary deployment

### Security Features
- Bcrypt password hashing
- JWT token authentication
- Input validation
- SQL injection protection (parameterized queries)
- XSS protection via headers
- CSRF protection
- Rate limiting
- IP-based access control
- TLS/HTTPS by default

## [0.1.0] - 2025-01-17

### Initial Release
- First public release of CaddyProxyManager+
- All core features implemented
- Production-ready with security features
- Docker and manual installation support
- Comprehensive documentation

---

## Future Roadmap

### Planned Features

#### v0.2.0
- [ ] Live log viewer in UI
- [ ] CrowdSec dashboard integration
- [ ] GoAccess integration for analytics
- [ ] Access log monitoring and filtering
- [ ] Statistics dashboard
- [ ] User management UI
- [ ] Settings page functionality

#### v0.3.0
- [ ] Certificate management UI
- [ ] Custom Caddy directive editor with syntax highlighting
- [ ] Backup/restore functionality
- [ ] Import/export configurations
- [ ] Configuration templates
- [ ] Bulk operations on proxy hosts

#### v0.4.0
- [ ] Multi-user support with roles
- [ ] Audit logging
- [ ] Email notifications for cert renewal
- [ ] Webhook support for events
- [ ] API documentation (OpenAPI/Swagger)
- [ ] Command-line interface (CLI)

#### v0.5.0
- [ ] Plugin system for custom handlers
- [ ] Theme customization
- [ ] Mobile app (PWA)
- [ ] Prometheus metrics export
- [ ] Grafana dashboard templates
- [ ] Health check monitoring

### Community Requests
- [ ] Let's Encrypt staging environment toggle
- [ ] Custom certificate upload
- [ ] HTTP/3 support configuration
- [ ] Load balancing configuration
- [ ] Failover support
- [ ] WebSocket handling configuration
- [ ] gRPC support

### Known Issues
- None reported yet

### Breaking Changes
- None

---

## How to Contribute

We welcome contributions! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Support

- GitHub Issues: https://github.com/Wikid82/CaddyProxyManagerPlus/issues
- Documentation: See repository docs
- Examples: See [EXAMPLES.md](EXAMPLES.md)

## License

MIT License - see [LICENSE](LICENSE) file for details.
