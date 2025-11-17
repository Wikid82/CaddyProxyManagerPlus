// Global state
let currentView = 'proxy-hosts';
let authToken = localStorage.getItem('authToken');
let currentEditingHost = null;

// Initialize app
document.addEventListener('DOMContentLoaded', function() {
    if (authToken) {
        showDashboard();
        loadProxyHosts();
    } else {
        showLogin();
    }

    setupEventListeners();
});

// Event Listeners
function setupEventListeners() {
    // Login form
    document.getElementById('login-form').addEventListener('submit', handleLogin);

    // Navigation
    document.querySelectorAll('.nav-btn').forEach(btn => {
        btn.addEventListener('click', handleNavigation);
    });

    // Logout
    document.getElementById('logout-btn').addEventListener('click', handleLogout);

    // Add host button
    document.getElementById('add-host-btn').addEventListener('click', showAddHostModal);

    // Modal controls
    document.getElementById('close-modal').addEventListener('click', closeModal);
    document.getElementById('cancel-modal').addEventListener('click', closeModal);
    document.getElementById('host-form').addEventListener('submit', handleHostSubmit);

    // Tab switching
    document.querySelectorAll('.tab-btn').forEach(btn => {
        btn.addEventListener('click', handleTabSwitch);
    });

    // Conditional form fields
    setupConditionalFields();
}

function setupConditionalFields() {
    // DNS Challenge
    const dnsChallenge = document.querySelector('input[name="dns_challenge"]');
    dnsChallenge?.addEventListener('change', function() {
        const groups = document.querySelectorAll('.dns-provider-group');
        groups.forEach(g => g.style.display = this.checked ? 'block' : 'none');
    });

    // Basic Auth
    const basicAuth = document.querySelector('input[name="basic_auth_enabled"]');
    basicAuth?.addEventListener('change', function() {
        const groups = document.querySelectorAll('.basic-auth-group');
        groups.forEach(g => g.style.display = this.checked ? 'block' : 'none');
    });

    // Forward Auth
    const forwardAuth = document.querySelector('input[name="forward_auth_enabled"]');
    forwardAuth?.addEventListener('change', function() {
        const groups = document.querySelectorAll('.forward-auth-group');
        groups.forEach(g => g.style.display = this.checked ? 'block' : 'none');
    });

    // Rate Limit
    const rateLimit = document.querySelector('input[name="rate_limit_enabled"]');
    rateLimit?.addEventListener('change', function() {
        const groups = document.querySelectorAll('.rate-limit-group');
        groups.forEach(g => g.style.display = this.checked ? 'block' : 'none');
    });

    // Geo Blocking
    const geoBlock = document.querySelector('input[name="geo_block_enabled"]');
    geoBlock?.addEventListener('change', function() {
        const groups = document.querySelectorAll('.geo-block-group');
        groups.forEach(g => g.style.display = this.checked ? 'block' : 'none');
    });

    // mTLS
    const mtls = document.querySelector('input[name="mtls_enabled"]');
    mtls?.addEventListener('change', function() {
        const groups = document.querySelectorAll('.mtls-group');
        groups.forEach(g => g.style.display = this.checked ? 'block' : 'none');
    });
}

// Authentication
async function handleLogin(e) {
    e.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('/api/auth/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password })
        });

        if (response.ok) {
            const data = await response.json();
            authToken = data.token;
            localStorage.setItem('authToken', authToken);
            showDashboard();
            loadProxyHosts();
        } else {
            const error = await response.json();
            showError('login-error', error.error || 'Login failed');
        }
    } catch (err) {
        showError('login-error', 'Network error: ' + err.message);
    }
}

function handleLogout() {
    authToken = null;
    localStorage.removeItem('authToken');
    showLogin();
}

// Navigation
function handleNavigation(e) {
    const view = e.target.dataset.view;
    if (!view) return;

    // Update active nav button
    document.querySelectorAll('.nav-btn').forEach(btn => btn.classList.remove('active'));
    e.target.classList.add('active');

    // Show corresponding view
    document.querySelectorAll('.view').forEach(v => v.classList.remove('active'));
    document.getElementById(view + '-view').classList.add('active');

    currentView = view;

    // Load view-specific data
    if (view === 'monitoring') {
        loadMonitoring();
    }
}

function handleTabSwitch(e) {
    const tab = e.target.dataset.tab;
    
    // Update active tab button
    document.querySelectorAll('.tab-btn').forEach(btn => btn.classList.remove('active'));
    e.target.classList.add('active');

    // Show corresponding tab content
    document.querySelectorAll('.tab-content').forEach(content => {
        if (content.dataset.tab === tab) {
            content.classList.add('active');
        } else {
            content.classList.remove('active');
        }
    });
}

// Screen Management
function showLogin() {
    document.getElementById('login-screen').style.display = 'block';
    document.getElementById('dashboard-screen').style.display = 'none';
}

function showDashboard() {
    document.getElementById('login-screen').style.display = 'none';
    document.getElementById('dashboard-screen').style.display = 'block';
}

// Proxy Hosts
async function loadProxyHosts() {
    try {
        const response = await fetch('/api/proxy-hosts', {
            headers: { 'Authorization': 'Bearer ' + authToken }
        });

        if (response.ok) {
            const hosts = await response.json();
            renderProxyHosts(hosts);
        } else if (response.status === 401) {
            handleLogout();
        }
    } catch (err) {
        console.error('Failed to load proxy hosts:', err);
    }
}

function renderProxyHosts(hosts) {
    const hostsList = document.getElementById('hosts-list');
    
    if (hosts.length === 0) {
        hostsList.innerHTML = `
            <div class="card">
                <p style="text-align: center; padding: 40px;">
                    No proxy hosts configured yet. Click "Add Proxy Host" to get started.
                </p>
            </div>
        `;
        return;
    }

    hostsList.innerHTML = hosts.map(host => `
        <div class="host-card ${!host.enabled ? 'disabled' : ''}">
            <div class="host-info">
                <h3>${host.domain_names}</h3>
                <p>→ ${host.scheme}://${host.forward_host}:${host.forward_port}</p>
                <div class="host-badges">
                    ${host.ssl_enabled ? '<span class="badge badge-success">🔒 SSL</span>' : ''}
                    ${host.waf_enabled ? '<span class="badge badge-success">🛡️ WAF</span>' : ''}
                    ${host.rate_limit_enabled ? '<span class="badge badge-info">🚦 Rate Limited</span>' : ''}
                    ${host.forward_auth_enabled ? '<span class="badge badge-info">🔐 SSO</span>' : ''}
                    ${host.basic_auth_enabled ? '<span class="badge badge-info">🔑 Basic Auth</span>' : ''}
                    ${host.crowdsec_enabled ? '<span class="badge badge-success">🤖 CrowdSec</span>' : ''}
                    ${host.local_only_enabled ? '<span class="badge badge-warning">🏠 Local Only</span>' : ''}
                    ${host.geo_block_enabled ? '<span class="badge badge-info">🌍 Geo-blocking</span>' : ''}
                    ${!host.enabled ? '<span class="badge badge-warning">⏸️ Disabled</span>' : ''}
                </div>
            </div>
            <div class="host-actions">
                <button class="btn btn-secondary" onclick="editHost(${host.id})">✏️ Edit</button>
                <button class="btn ${host.enabled ? 'btn-warning' : 'btn-success'}" onclick="toggleHost(${host.id})">
                    ${host.enabled ? '⏸️ Disable' : '▶️ Enable'}
                </button>
                <button class="btn btn-danger" onclick="deleteHost(${host.id})">🗑️ Delete</button>
            </div>
        </div>
    `).join('');
}

// Modal Management
function showAddHostModal() {
    currentEditingHost = null;
    document.getElementById('modal-title').textContent = 'Add Proxy Host';
    document.getElementById('host-form').reset();
    document.getElementById('host-modal').style.display = 'flex';
}

async function editHost(id) {
    try {
        const response = await fetch(`/api/proxy-hosts/${id}`, {
            headers: { 'Authorization': 'Bearer ' + authToken }
        });

        if (response.ok) {
            const host = await response.json();
            currentEditingHost = host;
            populateHostForm(host);
            document.getElementById('modal-title').textContent = 'Edit Proxy Host';
            document.getElementById('host-modal').style.display = 'flex';
        }
    } catch (err) {
        console.error('Failed to load host:', err);
    }
}

function populateHostForm(host) {
    const form = document.getElementById('host-form');
    
    // Populate all form fields
    Object.keys(host).forEach(key => {
        const input = form.querySelector(`[name="${key}"]`);
        if (input) {
            if (input.type === 'checkbox') {
                input.checked = host[key];
                // Trigger change event to show/hide conditional fields
                input.dispatchEvent(new Event('change'));
            } else {
                input.value = host[key] || '';
            }
        }
    });
}

function closeModal() {
    document.getElementById('host-modal').style.display = 'none';
    currentEditingHost = null;
}

async function handleHostSubmit(e) {
    e.preventDefault();
    const form = e.target;
    const formData = new FormData(form);
    const data = {};

    // Convert form data to object
    formData.forEach((value, key) => {
        if (form.querySelector(`[name="${key}"]`).type === 'checkbox') {
            data[key] = form.querySelector(`[name="${key}"]`).checked;
        } else if (form.querySelector(`[name="${key}"]`).type === 'number') {
            data[key] = parseInt(value) || 0;
        } else {
            data[key] = value;
        }
    });

    try {
        const url = currentEditingHost 
            ? `/api/proxy-hosts/${currentEditingHost.id}`
            : '/api/proxy-hosts';
        
        const method = currentEditingHost ? 'PUT' : 'POST';

        const response = await fetch(url, {
            method,
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + authToken
            },
            body: JSON.stringify(data)
        });

        if (response.ok) {
            closeModal();
            loadProxyHosts();
        } else {
            const error = await response.json();
            alert('Error: ' + (error.error || 'Failed to save proxy host'));
        }
    } catch (err) {
        alert('Network error: ' + err.message);
    }
}

async function toggleHost(id) {
    try {
        const response = await fetch(`/api/proxy-hosts/${id}/toggle`, {
            method: 'POST',
            headers: { 'Authorization': 'Bearer ' + authToken }
        });

        if (response.ok) {
            loadProxyHosts();
        }
    } catch (err) {
        console.error('Failed to toggle host:', err);
    }
}

async function deleteHost(id) {
    if (!confirm('Are you sure you want to delete this proxy host?')) {
        return;
    }

    try {
        const response = await fetch(`/api/proxy-hosts/${id}`, {
            method: 'DELETE',
            headers: { 'Authorization': 'Bearer ' + authToken }
        });

        if (response.ok) {
            loadProxyHosts();
        }
    } catch (err) {
        console.error('Failed to delete host:', err);
    }
}

// Monitoring
async function loadMonitoring() {
    // Placeholder for monitoring data
    document.getElementById('crowdsec-bans').textContent = '0';
    document.getElementById('log-content').textContent = 'No recent logs';
}

// Utilities
function showError(elementId, message) {
    const element = document.getElementById(elementId);
    element.textContent = message;
    element.style.display = 'block';
    setTimeout(() => {
        element.style.display = 'none';
    }, 5000);
}

// Global functions for inline onclick handlers
window.editHost = editHost;
window.toggleHost = toggleHost;
window.deleteHost = deleteHost;
