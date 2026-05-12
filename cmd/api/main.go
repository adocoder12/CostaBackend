package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/adocoder12/Costabackend/internal/config"
	"github.com/adocoder12/Costabackend/internal/db"
	"github.com/adocoder12/Costabackend/internal/handler"
	"github.com/adocoder12/Costabackend/internal/repository/postgres"
	"github.com/adocoder12/Costabackend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Structured logger — slog is stdlib since Go 1.21
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	// 2. Load .env — tolerant, production uses system env
	if err := godotenv.Load(); err != nil {
		// FIX: use slog consistently — not log.Printf
		logger.Warn("no .env file found, using system environment")
	}

	// 3. Load config — fails fast if required vars are missing
	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	// 4. Set Gin mode before creating the engine
	gin.SetMode(cfg.Server.GinMode)

	// 5. Initialize DB pool
	pool, err := db.NewPool(cfg.Database)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()
	logger.Info("🗄️  database connection established")

	// 6. Run migrations — FIX: use DSN() not MigrationURL() which doesn't exist
	if err := db.Migrate(cfg.Database.DSN()); err != nil {
		logger.Error("failed to run migrations", "error", err)
		os.Exit(1)
	}
	logger.Info("✅ database migrations verified")

	// 7. Wire dependencies bottom-up: repo → service → handler
	aptRepo := &postgres.ApartmentRepository{Pool: pool}
	aptService := services.NewApartmentService(aptRepo, logger)
	app := handler.NewApp(logger, cfg, aptService)

	// 8. Setup Gin routes
	engine := app.SetupRoutes()

	// 9. Start HTTP server with timeouts
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	logger.Info("🚀 Costa PMS API starting", "port", cfg.Server.Port)
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Error("server failed", "error", err)
		os.Exit(1)
	}
}
