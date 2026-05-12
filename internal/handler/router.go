package handler

import (
	"net/http"

	"github.com/adocoder12/Costabackend/internal/handler/middlewares"
	"github.com/gin-gonic/gin"
)

func (app *App) SetupRoutes() *gin.Engine {
	// Use gin.New() instead of gin.Default() so we control all middleware.
	// gin.Default() adds its own logger which conflicts with our slog setup.
	r := gin.New()

	// ── Global middleware ─────────────────────────────────────────────────────
	r.Use(gin.Recovery()) // recover from panics, return 500
	r.Use(middlewares.CorsMiddleware(app.Config.CORS.AllowedOrigin))

	app.Logger.Info("initializing API routes")

	// ── System routes (public) ────────────────────────────────────────────────
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "operational",
			"service": "costa-backend",
		})
	})

	// ── API v1 ────────────────────────────────────────────────────────────────
	v1 := r.Group("/api/v1")

	// ── Apartments ───────────────────────────────────────────────────────────
	apartments := v1.Group("/apartments")
	{
		apartments.GET("", app.GetApartmentsHandler)
		apartments.POST("", app.CreateApartmentHandler)
		apartments.GET("/:id", app.GetApartmentByIDHandler)
		apartments.PUT("/:id", app.UpdateApartmentHandler)
		apartments.DELETE("/:id", app.DeleteApartmentHandler)
	}

	// ── Cleaners ─────────────────────────────────────────────────────────────
	cleaners := v1.Group("/cleaners")
	{
		cleaners.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "cleaners module coming soon"})
		})
	}

	// ── Bookings ─────────────────────────────────────────────────────────────
	bookings := v1.Group("/bookings")
	{
		bookings.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "bookings module coming soon"})
		})
	}

	// ── Cleaning Tasks ───────────────────────────────────────────────────────
	tasks := v1.Group("/tasks")
	{
		tasks.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "tasks module coming soon"})
		})
	}

	// ── Dashboard ────────────────────────────────────────────────────────────
	v1.GET("/dashboard/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "dashboard module coming soon"})
	})

	return r
}
