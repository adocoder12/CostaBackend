package main

import (
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/adocoder12/Costabackend/internal/config"
	"github.com/adocoder12/Costabackend/internal/db"
	"github.com/adocoder12/Costabackend/internal/handler"
	"github.com/adocoder12/Costabackend/internal/repository/postgres"
	"github.com/adocoder12/Costabackend/internal/services" // Add this
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using system env: %v", err)
	}

	// 2. Initialize Structured Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	// 3. Load Config
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	// 4. Initialize DB Pool
	pool, err := db.NewPool(cfg.Database)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()
	logger.Info("Database connection established")

	// 5. Run Migrations
	if err := db.Migrate(cfg.Database.MigrationURL()); err != nil {
		logger.Error("failed to run migrations", "error", err)
		os.Exit(1)
	}
	logger.Info("✅ database migrations verified")

	// 6. WIRE DEPENDENCIES (The "Magic" Step)
	// Initialize Repository
	aptRepo := &postgres.ApartmentRepository{Pool: pool}

	// Initialize Service (Inject the repo and logger)
	aptService := services.NewApartmentService(aptRepo, logger)

	// Initialize App Handler (Inject the service)
	// Assuming NewApp(logger, cfg, apartmentService, cleanerService)
	app := handler.NewApp(logger, cfg, aptService, nil)

	// 7. Setup Gin routes
	engine := app.SetupRoutes()

	// 8. Start HTTP server
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
