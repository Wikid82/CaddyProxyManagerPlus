package main

import (
	"log"

	"github.com/Wikid82/CaddyProxyManagerPlus/internal/api"
	"github.com/Wikid82/CaddyProxyManagerPlus/internal/caddy"
	"github.com/Wikid82/CaddyProxyManagerPlus/internal/config"
	"github.com/Wikid82/CaddyProxyManagerPlus/internal/database"
	"github.com/Wikid82/CaddyProxyManagerPlus/internal/middleware"
	"github.com/Wikid82/CaddyProxyManagerPlus/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)



func main() {
	log.Println("🚀 Starting CaddyProxyManager+")

	// Load configuration
	cfg := config.Load()

	// Initialize database
	if err := database.Initialize(cfg.DataPath); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Set JWT secret for auth
	api.SetJWTSecret(cfg.JWTSecret)
	middleware.SetJWTSecret(cfg.JWTSecret)

	// Initialize Caddy client
	caddyClient := caddy.NewClient(cfg.CaddyAdminURL)
	api.SetCaddyClient(caddyClient)

	// Ensure default admin user exists
	ensureDefaultUser()

	// Set up Gin router
	router := setupRouter()

	// Start server
	addr := ":" + cfg.ServerPort
	log.Printf("✅ Server listening on http://localhost%s", addr)
	log.Printf("📝 Login with default credentials: admin / admin")
	
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./web/static")

	// Serve index.html for root
	router.StaticFile("/", "./web/templates/index.html")

	// Public routes
	publicAPI := router.Group("/api")
	{
		publicAPI.POST("/auth/login", api.Login)
	}

	// Protected routes
	protectedAPI := router.Group("/api")
	protectedAPI.Use(middleware.AuthMiddleware())
	{
		// Proxy hosts
		protectedAPI.GET("/proxy-hosts", api.ListProxyHosts)
		protectedAPI.GET("/proxy-hosts/:id", api.GetProxyHost)
		protectedAPI.POST("/proxy-hosts", api.CreateProxyHost)
		protectedAPI.PUT("/proxy-hosts/:id", api.UpdateProxyHost)
		protectedAPI.DELETE("/proxy-hosts/:id", api.DeleteProxyHost)
		protectedAPI.POST("/proxy-hosts/:id/toggle", api.ToggleProxyHost)
	}

	return router
}

func ensureDefaultUser() {
	var count int64
	database.DB.Model(&models.User{}).Count(&count)

	if count == 0 {
		log.Println("Creating default admin user...")
		
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		user := models.User{
			Username: "admin",
			Email:    "admin@localhost",
			Password: string(hashedPassword),
			IsAdmin:  true,
			Enabled:  true,
		}

		if err := database.DB.Create(&user).Error; err != nil {
			log.Fatalf("Failed to create default user: %v", err)
		}

		log.Println("✅ Default admin user created (username: admin, password: admin)")
		log.Println("⚠️  Please change the default password immediately!")
	}
}
