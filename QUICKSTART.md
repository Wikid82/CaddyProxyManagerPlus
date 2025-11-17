# Quick Start Guide

Get CaddyProxyManager+ running in 5 minutes!

## The Fastest Way: Docker Compose

### Step 1: Get Docker
If you don't have Docker installed:
- **Linux**: `curl -fsSL https://get.docker.com | sh`
- **macOS/Windows**: Install [Docker Desktop](https://www.docker.com/products/docker-desktop)

### Step 2: Clone and Start
```bash
git clone https://github.com/Wikid82/CaddyProxyManagerPlus.git
cd CaddyProxyManagerPlus
docker-compose up -d
```

### Step 3: Access the UI
- Open your browser to http://localhost:8080
- Login with:
  - **Username**: `admin`
  - **Password**: `admin`
- ⚠️ **Change the password immediately!**

## Your First Proxy Host

Let's proxy a local service running on port 8000.

### Step 1: Add a Proxy Host
1. Click "➕ Add Proxy Host"
2. Fill in the **Basic** tab:
   - **Domain Names**: `app.local` (or your domain)
   - **Scheme**: `http`
   - **Forward Host**: `host.docker.internal` (or your service IP)
   - **Forward Port**: `8000`

3. Go to **SSL/TLS** tab:
   - ✅ Enable SSL/TLS
   - ✅ Force HTTPS Redirect

4. Go to **Security** tab:
   - ✅ Enable CrowdSec Protection
   - ✅ Enable Rate Limiting
   - Select preset: "Standard (100 req/min)"

5. Click **Save**

### Step 2: Update Your Hosts File (for testing)
If using `app.local`:

**Linux/macOS:**
```bash
echo "127.0.0.1 app.local" | sudo tee -a /etc/hosts
```

**Windows** (as Administrator):
```bash
echo 127.0.0.1 app.local >> C:\Windows\System32\drivers\etc\hosts
```

### Step 3: Access Your Service
- Open https://app.local
- You'll see a self-signed certificate warning (this is normal for local development)
- Your service is now proxied with security features!

## Adding Real SSL Certificates

For production with a real domain:

### Step 1: Point Your Domain to Your Server
Update DNS A record:
```
app.example.com → YOUR_SERVER_IP
```

### Step 2: Create Proxy Host
1. Click "➕ Add Proxy Host"
2. **Basic** tab:
   - **Domain Names**: `app.example.com`
   - **Forward Host**: `192.168.1.100` (your internal server)
   - **Forward Port**: `8080`

3. **SSL/TLS** tab:
   - ✅ Enable SSL/TLS
   - ✅ Force HTTPS
   - ✅ Enable HSTS

4. **Security** tab:
   - ✅ Enable WAF (important for public sites!)
   - ✅ Enable Rate Limiting
   - ✅ Enable CrowdSec

5. Click **Save**

Let's Encrypt will automatically issue a certificate!

## Adding Authentication (SSO)

Protect any service with single sign-on:

### Using Authelia (example)

1. Set up Authelia (separate service)
2. Edit your proxy host
3. Go to **Authentication** tab:
   - ✅ Enable Forward Authentication
   - **Provider**: Authelia
   - **Forward Auth URL**: `https://auth.example.com/api/verify?rd=https://app.example.com`
4. Save

Now users must authenticate with Authelia before accessing the service!

## Common Configurations

### WordPress Site
**Security Tab:**
- ✅ WAF (critical!)
- ✅ Rate Limiting (Login preset)
- ✅ CrowdSec
- ✅ Geo-blocking (optional, but recommended)

### Plex/Jellyfin
**Security Tab:**
- ✅ Forward Auth (optional, for extra security)
- ✅ Rate Limiting (Standard preset)
- ✅ CrowdSec
- ✅ Geo-blocking (restrict to your country)

### Internal Admin Panel
**Security Tab:**
- ✅ Local Network Only (RFC1918)
- No other auth needed - network provides security

### Public API
**Security Tab:**
- ✅ WAF
- ✅ Rate Limiting (API preset)
- ✅ CrowdSec
- Consider: mTLS for machine-to-machine

## Wildcard Certificates (*.example.com)

For internal servers or multiple subdomains:

### Step 1: Get API Token
Get an API token from your DNS provider (Cloudflare, etc.)

### Step 2: Configure Proxy Host
1. **SSL/TLS** tab:
   - ✅ Enable SSL/TLS
   - ✅ Use DNS Challenge
   - **DNS Provider**: Cloudflare (or yours)
   - **API Credentials**:
   ```json
   {
     "api_token": "your-cloudflare-api-token"
   }
   ```

2. **Domain Names**: `*.example.com` or `subdomain.example.com`

3. Save

Caddy will use DNS challenge to get certificates!

## Troubleshooting

### Cannot access the UI
```bash
# Check if service is running
docker-compose ps

# View logs
docker-compose logs caddyproxymanager

# Restart
docker-compose restart
```

### Service not proxied correctly
1. Check if the forward host is accessible from the container
2. For localhost services, use `host.docker.internal` instead of `localhost`
3. Check Caddy logs: `docker-compose logs caddy`

### SSL certificate not issued
1. Verify DNS points to your server
2. Check ports 80/443 are accessible
3. For DNS challenge: verify API credentials
4. Check Caddy logs for specific errors

### WAF blocking legitimate requests
1. Check access logs in Monitoring tab
2. Temporarily disable WAF to verify
3. May need custom WAF rules (advanced)

## Next Steps

Now that you're running:

1. **Secure your installation**
   - Change default password
   - Set strong JWT secret
   - Restrict admin API access

2. **Add your services**
   - Start with non-critical services
   - Test thoroughly before production
   - Enable security features incrementally

3. **Set up monitoring**
   - Check CrowdSec dashboard
   - Monitor access logs
   - Set up alerts (external tool)

4. **Learn advanced features**
   - Read [SECURITY_FEATURES.md](SECURITY_FEATURES.md)
   - Check out [EXAMPLES.md](EXAMPLES.md)
   - Explore [INSTALLATION.md](INSTALLATION.md)

## Getting Help

- **Documentation**: Check the docs in this repository
- **Issues**: [GitHub Issues](https://github.com/Wikid82/CaddyProxyManagerPlus/issues)
- **Examples**: See [EXAMPLES.md](EXAMPLES.md) for real-world configs

## Common Questions

**Q: Can I use this with existing Caddy config?**
A: Yes! CaddyProxyManager+ uses the admin API and won't interfere with static configs.

**Q: Is this production-ready?**
A: The core features are solid, but always test thoroughly. Community feedback will improve it!

**Q: What's the performance impact?**
A: Minimal. Most overhead is from the security features you enable (WAF, rate limiting).

**Q: Can I backup my configuration?**
A: Yes! The SQLite database in `/data` contains everything. Regular backups recommended.

**Q: How do I update?**
A: `docker-compose pull && docker-compose up -d`

**Q: Why CaddyProxyManager+ instead of Nginx Proxy Manager?**
A: 
- Automatic HTTPS (way easier!)
- Modern architecture
- Better security features built-in
- Active development of Caddy
- No need for certbot

**Q: Why not just use Caddy directly?**
A: You can! But this gives you:
- Easy web UI
- Visual security configuration
- No JSON/Caddyfile editing
- Perfect for home-lab users

## Tips for Success

1. **Start Simple**: Add basic proxy first, then add security features
2. **Test Locally**: Use .local domains for testing before production
3. **Document Your Setup**: Keep notes on your configuration
4. **Monitor Logs**: Check logs regularly, especially after changes
5. **Update Regularly**: Keep CaddyProxyManager+ and Caddy updated

## Welcome to CaddyProxyManager+! 🎉

You're now running a modern, secure reverse proxy with enterprise-grade security features, all managed through a simple web UI.

Happy proxying! 🚀
