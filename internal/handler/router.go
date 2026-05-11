package handler

import (
	"github.com/adocoder12/Costabackend/internal/handler/middlewares"
	"github.com/gin-gonic/gin"
)

func (a *App) SetupRoutes() *gin.Engine {
	r := gin.Default()

	// 1. Global Middleware (CORS, Recovery, Logging)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CorsMiddleware())
	a.Logger.Info("Routes initialized")

	// 2. Public Routes (Health check, etc.)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 3. API Version 1 Group
	v1 := r.Group("/api/v1")
	{
		// --- Apartment Routes ---
		apartments := v1.Group("/apartments")
		{
			apartments.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Apartments endpoint is live"})
			})
			apartments.POST("", a.CreateApartmentHandler) // POST /api/v1/apartments
			//apartments.GET("", a.ListApartments)        // GET /api/v1/apartments
			//apartments.GET("/:id", a.GetApartment)      // GET /api/v1/apartments/:id
			//apartments.PATCH("/:id", a.UpdateApartment) // PATCH /api/v1/apartments/:id
		}

		// --- Cleaner Routes ---
		cleaners := v1.Group("/cleaners")
		{
			cleaners.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Cleaners endpoint is live"})
			})
		}

	}

	return r
}
