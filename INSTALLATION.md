# Installation Guide

## Quick Start with Docker Compose (Recommended)

The easiest way to get started with CaddyProxyManager+ is using Docker Compose.

### Prerequisites
- Docker and Docker Compose installed
- Ports 80, 443, and 8080 available

### Steps

1. **Clone the repository**
```bash
git clone https://github.com/Wikid82/CaddyProxyManagerPlus.git
cd CaddyProxyManagerPlus
```

2. **Create environment file** (optional)
```bash
cat > .env <<EOF
JWT_SECRET=$(openssl rand -base64 32)
EOF
```

3. **Start the services**
```bash
docker-compose up -d
```

4. **Access the UI**
- Open http://localhost:8080 in your browser
- Login with default credentials:
  - Username: `admin`
  - Password: `admin`
- **⚠️ Change the default password immediately!**

## Manual Installation

### Prerequisites
- Go 1.21 or later
- Caddy 2.7 or later
- SQLite (included with Go)

### Steps

1. **Clone and build**
```bash
git clone https://github.com/Wikid82/CaddyProxyManagerPlus.git
cd CaddyProxyManagerPlus
go build -o caddyproxymanager ./cmd/caddyproxymanager
```

2. **Install Caddy**

**Linux:**
```bash
sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list
sudo apt update
sudo apt install caddy
```

**macOS:**
```bash
brew install caddy
```

3. **Configure Caddy**

Enable the Caddy admin API in `/etc/caddy/Caddyfile`:
```
{
    admin 0.0.0.0:2019
}
```

Restart Caddy:
```bash
sudo systemctl restart caddy
```

4. **Run CaddyProxyManager+**
```bash
./caddyproxymanager
```

5. **Access the UI**
- Open http://localhost:8080
- Login with default credentials (admin/admin)

## Production Deployment

### Security Hardening

1. **Change default credentials immediately**

2. **Use a strong JWT secret**
```bash
export JWT_SECRET=$(openssl rand -base64 32)
```

3. **Restrict admin API access**
   
Edit Caddy configuration to restrict admin API:
```
{
    admin 127.0.0.1:2019
}
```

4. **Use environment variables for secrets**
```bash
export JWT_SECRET="your-secret-key"
export CADDY_ADMIN_URL="http://localhost:2019"
```

### Systemd Service

Create `/etc/systemd/system/caddyproxymanager.service`:

```ini
[Unit]
Description=CaddyProxyManager+ Service
After=network.target caddy.service
Requires=caddy.service

[Service]
Type=simple
User=caddy
Group=caddy
WorkingDirectory=/opt/caddyproxymanager
ExecStart=/opt/caddyproxymanager/caddyproxymanager
Restart=on-failure
RestartSec=5s

Environment="DATA_PATH=/var/lib/caddyproxymanager"
Environment="SERVER_PORT=8080"
Environment="CADDY_ADMIN_URL=http://localhost:2019"
Environment="JWT_SECRET=change-this-in-production"

[Install]
WantedBy=multi-user.target
```

Enable and start:
```bash
sudo systemctl enable caddyproxymanager
sudo systemctl start caddyproxymanager
```

### Reverse Proxy for the UI

To access the UI securely via HTTPS, add to Caddyfile:

```
admin.example.com {
    reverse_proxy localhost:8080
}
```

## Integration with Existing Caddy Setup

If you already have Caddy running:

1. **Ensure admin API is enabled**
```
{
    admin localhost:2019
}
```

2. **Run CaddyProxyManager+ with existing Caddy**
```bash
export CADDY_ADMIN_URL="http://localhost:2019"
./caddyproxymanager
```

3. **CaddyProxyManager+ will manage your Caddy configuration**
   - Existing static configurations in Caddyfile will be preserved
   - Dynamic proxy hosts are managed via the admin API

## CrowdSec Integration

### Install CrowdSec

**Ubuntu/Debian:**
```bash
curl -s https://packagecloud.io/install/repositories/crowdsec/crowdsec/script.deb.sh | sudo bash
sudo apt install crowdsec
```

### Install Caddy Bouncer

```bash
sudo apt install crowdsec-caddy-bouncer
```

### Configure Bouncer

1. Generate bouncer API key:
```bash
sudo cscli bouncers add caddy-bouncer
```

2. Configure in Caddy:
```
{
    order crowdsec first
}

(crowdsec) {
    crowdsec {
        api_url http://localhost:8080
        api_key <your-bouncer-key>
    }
}
```

3. Enable CrowdSec in CaddyProxyManager+ (enabled by default)

## DNS Challenge Setup

For wildcard certificates or internal servers:

### Cloudflare

1. Get API token from Cloudflare dashboard
2. In CaddyProxyManager+ UI:
   - Enable "DNS Challenge"
   - Select "Cloudflare"
   - Enter API token:
   ```json
   {
     "api_token": "your-cloudflare-api-token"
   }
   ```

### AWS Route53

1. Create IAM user with Route53 permissions
2. In CaddyProxyManager+ UI:
   - Enable "DNS Challenge"
   - Select "AWS Route53"
   - Enter credentials:
   ```json
   {
     "aws_access_key_id": "your-key-id",
     "aws_secret_access_key": "your-secret-key"
   }
   ```

## Backup and Restore

### Backup

The database contains all configuration:
```bash
# Backup database
cp /data/caddyproxymanager.db /backup/caddyproxymanager-$(date +%Y%m%d).db

# Backup Caddy certificates
sudo tar czf /backup/caddy-certs-$(date +%Y%m%d).tar.gz /var/lib/caddy
```

### Restore

```bash
# Restore database
cp /backup/caddyproxymanager-YYYYMMDD.db /data/caddyproxymanager.db

# Restore certificates
sudo tar xzf /backup/caddy-certs-YYYYMMDD.tar.gz -C /
```

## Troubleshooting

### Cannot access UI

1. Check if service is running:
```bash
sudo systemctl status caddyproxymanager
```

2. Check logs:
```bash
sudo journalctl -u caddyproxymanager -f
```

3. Verify port 8080 is not in use:
```bash
sudo lsof -i :8080
```

### Cannot connect to Caddy admin API

1. Verify Caddy is running:
```bash
sudo systemctl status caddy
```

2. Test admin API:
```bash
curl http://localhost:2019/config/
```

3. Check Caddy admin API is enabled in Caddyfile

### Certificates not being issued

1. Check Caddy logs:
```bash
sudo journalctl -u caddy -f
```

2. Verify DNS records are correct
3. Ensure ports 80/443 are accessible
4. For DNS challenge, verify API credentials

### Database locked error

This happens when multiple instances are running:
```bash
# Find and stop all instances
ps aux | grep caddyproxymanager
sudo kill <pid>
```

## Upgrading

### Docker

```bash
docker-compose pull
docker-compose up -d
```

### Manual

```bash
git pull
go build -o caddyproxymanager ./cmd/caddyproxymanager
sudo systemctl restart caddyproxymanager
```

## Support

For issues and questions:
- GitHub Issues: https://github.com/Wikid82/CaddyProxyManagerPlus/issues
- Documentation: https://github.com/Wikid82/CaddyProxyManagerPlus/wiki

## Next Steps

After installation:
1. Change default admin password
2. Add your first proxy host
3. Configure security features (WAF, rate limiting)
4. Set up CrowdSec for enhanced protection
5. Configure DNS challenge for wildcard certificates
