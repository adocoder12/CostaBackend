package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/adocoder12/Costabackend/internal/config"
	"github.com/adocoder12/Costabackend/internal/db"
	"github.com/adocoder12/Costabackend/internal/handler"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load environment variables from .env

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using system env: %v", err)
	}

	// 2. Initialize Structured Logger (2026 Best Practice)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug, // Change to Info for production
	}))
	slog.SetDefault(logger)

	// 3. Load the central Config struct
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	// 4. Initialize Database Connection (pgxpool)
	// Use the DSN() helper we added to your config earlier
	pool, err := db.NewPool(cfg.Database)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()
	logger.Info("Database connection established")

	// 5. Run migrations — safe to call on every startup
	if err := db.Migrate(cfg.Database.MigrationURL()); err != nil {
		logger.Error("failed to run migrations", "error", err)
		os.Exit(1)
	}
	logger.Info("✅ database migrations verified")
	// 6. Wire dependencies into App
	// Repositories, services, and AWS clients added in Phase 3

	var count int
	errC := pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM apartments").Scan(&count)
	if errC != nil {
		logger.Error("Table does not exist!", "error", errC)
	} else {
		logger.Info("Table verified", "row_count", count)
	}

	app := handler.NewApp(logger, cfg)

	// 7. Setup Gin routes
	// Full routing added in Phase 4
	engine := app.SetupRoutes()

	// 8. Start HTTP server with timeouts
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second, // longer for S3 photo uploads
		IdleTimeout:  time.Minute,
	}

	logger.Info("🚀 Costa PMS API starting", "port", cfg.Server.Port)
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Error("server failed", "error", err)
		os.Exit(1)
	}
}
