package api

import (
	"net/http"
	"strconv"

	"github.com/Wikid82/CaddyProxyManagerPlus/internal/caddy"
	"github.com/Wikid82/CaddyProxyManagerPlus/internal/database"
	"github.com/Wikid82/CaddyProxyManagerPlus/internal/models"
	"github.com/gin-gonic/gin"
)

var caddyClient *caddy.Client

// SetCaddyClient sets the Caddy client for the API
func SetCaddyClient(client *caddy.Client) {
	caddyClient = client
}

// ListProxyHosts returns all proxy hosts
func ListProxyHosts(c *gin.Context) {
	var hosts []models.ProxyHost
	if err := database.DB.Find(&hosts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hosts)
}

// GetProxyHost returns a single proxy host
func GetProxyHost(c *gin.Context) {
	id := c.Param("id")
	var host models.ProxyHost
	if err := database.DB.First(&host, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Proxy host not found"})
		return
	}

	c.JSON(http.StatusOK, host)
}

// CreateProxyHost creates a new proxy host
func CreateProxyHost(c *gin.Context) {
	var host models.ProxyHost
	if err := c.ShouldBindJSON(&host); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&host).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Reload Caddy configuration
	if err := reloadCaddyConfig(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reload Caddy: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, host)
}

// UpdateProxyHost updates an existing proxy host
func UpdateProxyHost(c *gin.Context) {
	id := c.Param("id")
	var host models.ProxyHost
	if err := database.DB.First(&host, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Proxy host not found"})
		return
	}

	if err := c.ShouldBindJSON(&host); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse ID from param
	idInt, _ := strconv.ParseUint(id, 10, 32)
	host.ID = uint(idInt)

	if err := database.DB.Save(&host).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Reload Caddy configuration
	if err := reloadCaddyConfig(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reload Caddy: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, host)
}

// DeleteProxyHost deletes a proxy host
func DeleteProxyHost(c *gin.Context) {
	id := c.Param("id")
	var host models.ProxyHost
	if err := database.DB.First(&host, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Proxy host not found"})
		return
	}

	if err := database.DB.Delete(&host).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Reload Caddy configuration
	if err := reloadCaddyConfig(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reload Caddy: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Proxy host deleted"})
}

// ToggleProxyHost enables or disables a proxy host
func ToggleProxyHost(c *gin.Context) {
	id := c.Param("id")
	var host models.ProxyHost
	if err := database.DB.First(&host, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Proxy host not found"})
		return
	}

	host.Enabled = !host.Enabled
	if err := database.DB.Save(&host).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Reload Caddy configuration
	if err := reloadCaddyConfig(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reload Caddy: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, host)
}

// reloadCaddyConfig reloads the Caddy configuration with all active proxy hosts
func reloadCaddyConfig() error {
	var hosts []models.ProxyHost
	if err := database.DB.Find(&hosts).Error; err != nil {
		return err
	}

	config, err := caddy.GenerateConfig(hosts)
	if err != nil {
		return err
	}

	return caddyClient.LoadConfig(config)
}
