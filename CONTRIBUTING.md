# Contributing to CaddyProxyManager+

Thank you for your interest in contributing to CaddyProxyManager+! This document provides guidelines and instructions for contributing.

## Code of Conduct

Please be respectful and constructive in all interactions. We're building a welcoming community for everyone.

## How to Contribute

### Reporting Bugs

If you find a bug, please open an issue with:
- Clear description of the problem
- Steps to reproduce
- Expected behavior
- Actual behavior
- Screenshots (if applicable)
- Environment details (OS, Go version, Caddy version)
- Relevant logs

### Suggesting Features

Feature requests are welcome! Please include:
- Clear description of the feature
- Use case and motivation
- How it would work (UI mockups help!)
- Any security implications

### Contributing Code

1. **Fork the repository**
2. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Make your changes**
   - Follow the existing code style
   - Add tests for new functionality
   - Update documentation as needed
   - Keep commits focused and atomic

4. **Test your changes**
   ```bash
   # Build the application
   go build -o caddyproxymanager ./cmd/caddyproxymanager
   
   # Run tests
   go test ./...
   
   # Test manually
   ./caddyproxymanager
   ```

5. **Commit your changes**
   ```bash
   git commit -m "Add feature: description of feature"
   ```

6. **Push to your fork**
   ```bash
   git push origin feature/your-feature-name
   ```

7. **Open a Pull Request**
   - Describe what the PR does
   - Reference any related issues
   - Include screenshots for UI changes
   - Ensure CI passes

## Development Setup

### Prerequisites
- Go 1.21 or later
- Git
- A code editor (VS Code, GoLand, etc.)

### Local Development

1. **Clone your fork**
   ```bash
   git clone https://github.com/YOUR_USERNAME/CaddyProxyManagerPlus.git
   cd CaddyProxyManagerPlus
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Build and run**
   ```bash
   go build -o caddyproxymanager ./cmd/caddyproxymanager
   ./caddyproxymanager
   ```

4. **Access the UI**
   - Open http://localhost:8080
   - Login with admin/admin

### Project Structure

```
CaddyProxyManagerPlus/
├── cmd/
│   └── caddyproxymanager/     # Main application entry point
│       └── main.go
├── internal/
│   ├── api/                   # HTTP API handlers
│   │   ├── auth.go           # Authentication endpoints
│   │   └── proxy_hosts.go    # Proxy host CRUD
│   ├── caddy/                 # Caddy integration
│   │   ├── client.go         # Caddy admin API client
│   │   └── config_generator.go  # Config generation
│   ├── config/                # Application configuration
│   │   └── config.go
│   ├── database/              # Database layer
│   │   └── database.go
│   ├── middleware/            # HTTP middleware
│   │   └── auth.go           # JWT authentication
│   └── models/                # Data models
│       ├── proxy_host.go
│       ├── settings.go
│       └── user.go
├── web/
│   ├── static/
│   │   ├── css/              # Stylesheets
│   │   └── js/               # JavaScript
│   └── templates/            # HTML templates
├── go.mod
├── go.sum
├── Dockerfile
└── docker-compose.yml
```

## Coding Guidelines

### Go Code Style

Follow standard Go conventions:
- Use `gofmt` for formatting
- Use `golint` for linting
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Add comments for exported functions and types

Example:
```go
// CreateProxyHost creates a new proxy host configuration
// and reloads the Caddy server with the new configuration.
func CreateProxyHost(c *gin.Context) {
    // Implementation
}
```

### JavaScript Code Style

- Use ES6+ features
- Use `const` and `let`, not `var`
- Add JSDoc comments for complex functions
- Keep functions small and focused

Example:
```javascript
/**
 * Loads proxy hosts from the API and renders them
 * @returns {Promise<void>}
 */
async function loadProxyHosts() {
    // Implementation
}
```

### CSS Code Style

- Use CSS custom properties for theming
- Keep selectors specific but not overly nested
- Group related styles together
- Add comments for complex layouts

### Commit Messages

Follow conventional commits:
- `feat: Add new feature`
- `fix: Fix bug description`
- `docs: Update documentation`
- `style: Format code`
- `refactor: Refactor component`
- `test: Add tests`
- `chore: Update dependencies`

## Testing

### Unit Tests

Add unit tests for new functionality:

```go
func TestCreateProxyHost(t *testing.T) {
    // Setup
    // Test
    // Assert
}
```

Run tests:
```bash
go test ./...
```

### Integration Tests

For features that interact with Caddy:
1. Set up a test Caddy instance
2. Test the full flow
3. Verify the Caddy configuration

### Manual Testing

Before submitting a PR:
1. Test the happy path
2. Test error cases
3. Test with invalid input
4. Test UI responsiveness
5. Check browser console for errors

## Documentation

Update documentation when adding features:
- Update README.md for major features
- Update SECURITY_FEATURES.md for security features
- Update INSTALLATION.md for setup changes
- Add examples to EXAMPLES.md
- Update inline code comments

## Security

### Reporting Security Issues

**Do not open public issues for security vulnerabilities.**

Instead:
1. Email the maintainers privately
2. Include detailed description
3. Include steps to reproduce
4. Allow time for fix before disclosure

### Security Best Practices

When contributing code:
- Validate all user input
- Use parameterized queries for database
- Hash passwords with bcrypt
- Use HTTPS for external connections
- Follow OWASP guidelines
- Never log sensitive data
- Use constant-time comparisons for secrets

## Adding Security Features

When adding new security features:

1. **Research**: Understand the security mechanism thoroughly
2. **Design**: Plan the UI and configuration storage
3. **Implement**: Add to the model, config generator, and UI
4. **Document**: Add to SECURITY_FEATURES.md with examples
5. **Test**: Verify it actually provides the claimed protection

Example checklist for adding a new feature:
- [ ] Add fields to ProxyHost model
- [ ] Add UI controls in index.html
- [ ] Add config generation in config_generator.go
- [ ] Add to security features documentation
- [ ] Add example configuration
- [ ] Test the feature works
- [ ] Add unit tests

## UI Development

### Adding a New Tab

1. Add tab button in HTML:
```html
<button type="button" class="tab-btn" data-tab="newtab">New Tab</button>
```

2. Add tab content:
```html
<div class="tab-content" data-tab="newtab">
    <!-- Form fields -->
</div>
```

3. Tab switching is handled automatically by existing JavaScript

### Adding Form Fields

1. Add to HTML form
2. Add to ProxyHost model in Go
3. Handle in config_generator.go
4. Add database migration if needed

### Styling

- Use existing CSS custom properties
- Follow the dark theme design
- Ensure mobile responsiveness
- Test in Chrome, Firefox, and Safari

## Pull Request Process

1. **Keep PRs focused**: One feature per PR
2. **Update documentation**: Include doc updates in the same PR
3. **Add tests**: New features should include tests
4. **Ensure CI passes**: All checks must pass
5. **Respond to feedback**: Address review comments promptly
6. **Squash commits**: Before merging, squash into logical commits

## Release Process

(For maintainers)

1. Update version number
2. Update CHANGELOG.md
3. Create release tag
4. Build binaries
5. Create GitHub release
6. Update Docker images

## Questions?

If you have questions:
- Open a discussion on GitHub
- Check existing issues and PRs
- Read the documentation
- Ask in the community chat (if available)

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

## Recognition

Contributors will be recognized in:
- GitHub contributors page
- Release notes
- Project documentation (for major contributions)

Thank you for contributing to CaddyProxyManager+! 🎉
