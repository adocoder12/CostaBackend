package handler

import (
	"log/slog"
	"net/http"

	"github.com/adocoder12/Costabackend/internal/config"
	"github.com/adocoder12/Costabackend/internal/services"
	"github.com/gin-gonic/gin"
)

type App struct {
	Logger *slog.Logger
	Config *config.Config
	// Services added as interfaces in Phase 3:
	ApartmentService services.ApartmentServiceInterface
	// CleanerService   service.CleanerServiceInterface
	// BookingService   service.BookingServiceInterface
	// GuestService     service.GuestServiceInterface
	// TaskService      service.TaskServiceInterface
}

func NewApp(
	logger *slog.Logger,
	cfg *config.Config,
	aptService services.ApartmentServiceInterface,
) *App {
	return &App{
		Logger:           logger,
		Config:           cfg,
		ApartmentService: aptService,
	}
}

// serverError logs the full error detail (never exposed to client)
func (app *App) serverError(c *gin.Context, err error) {
	app.Logger.Error("internal server error",
		"path", c.Request.URL.Path,
		"method", c.Request.Method,
		"error", err.Error(),
	)

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": http.StatusText(http.StatusInternalServerError),
	})
}

// clientError sends a structured 4xx response with a human-readable message.
func (app *App) clientError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error":   http.StatusText(status),
		"message": message,
	})
}

// notFound is a 404 shorthand.
func (app *App) notFound(c *gin.Context) {
	app.clientError(c, http.StatusNotFound, "the requested resource was not found")
}
