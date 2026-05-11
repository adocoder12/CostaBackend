package handler

import (
	"github.com/adocoder12/Costabackend/internal/handler/middlewares"
	"github.com/gin-gonic/gin"
)

func (app *App) SetupRoutes() *gin.Engine {
	// 1. Initialize engine with default middleware (Logger & Recovery)
	r := gin.Default()

	// 2. Add Custom Global Middleware
	// It's a best practice to have a RequestID for tracing logs in slog
	r.Use(middlewares.CorsMiddleware())

	app.Logger.Info("API Routes initializing...")

	// 3. Public / System Routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "operational",
			"service": "costa-backend",
		})
	})

	// 4. API Version 1 Group
	v1 := r.Group("/api/v1")
	{
		// --- Apartment Routes ---
		apartments := v1.Group("/apartments")
		{
			// GET /api/v1/apartments -> List all
			apartments.GET("", app.GetApartmentsHandler)

			// POST /api/v1/apartments -> Create new
			apartments.POST("", app.CreateApartmentHandler)

			// GET /api/v1/apartments/:id -> Single Detail
			//apartments.GET("/:id", app.GetApartmentByIDHandler)

			// PATCH /api/v1/apartments/:id -> Partial Update
			// apartments.PATCH("/:id", app.UpdateApartmentHandler)
		}

		// --- Cleaner Routes ---
		cleaners := v1.Group("/cleaners")
		{
			cleaners.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Cleaners module active"})
			})
			// Add your cleaner handlers here as you build them
		}
	}

	return r
}
