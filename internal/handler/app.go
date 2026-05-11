package handler

import (
	"log/slog"
	"net/http"

	"github.com/adocoder12/Costabackend/internal/config"
	"github.com/gin-gonic/gin"
)

// All services are injected here as interfaces — never concrete types.
// This makes handlers testable without a real DB or S3.
type App struct {
	Logger *slog.Logger
	Config *config.Config
	// Services added as interfaces in Phase 3:
	// ApartmentService service.ApartmentServiceInterface
	// CleanerService   service.CleanerServiceInterface
	// BookingService   service.BookingServiceInterface
	// GuestService     service.GuestServiceInterface
	// TaskService      service.TaskServiceInterface
}

// NewApp wires all dependencies into the App struct.
func NewApp(logger *slog.Logger, cfg *config.Config) *App {
	return &App{
		Logger: logger,
		Config: cfg,
	}
}

// serverError logs the full error detail (never exposed to client)
func (a *App) serverError(c *gin.Context, err error) {
	a.Logger.Error("internal server error",
		"path", c.Request.URL.Path,
		"method", c.Request.Method,
		"error", err.Error(),
	)

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": http.StatusText(http.StatusInternalServerError),
	})
}

// clientError sends a structured 4xx response with a human-readable message.
func (a *App) clientError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error":   http.StatusText(status),
		"message": message,
	})
}

// notFound is a 404 shorthand.
func (a *App) notFound(c *gin.Context) {
	a.clientError(c, http.StatusNotFound, "the requested resource was not found")
}
