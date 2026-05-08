package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adocoder12/Costabackend/internal/config"
	"github.com/adocoder12/Costabackend/internal/handler"
	"github.com/adocoder12/Costabackend/internal/middleware"
	"github.com/adocoder12/Costabackend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/joho/godotenv"
)

func main() {

	//1. load env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// 2. Setup loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//3. load config
	cfg := config.Load()
	if err != nil {
		errorLog.Fatalf("Error loading config: %v", err)
	}

	//4. initialize db
	ctxPool, cancelCtxPool := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancelCtxPool()
	pool, err := pgxpool.New(ctxPool, cfg.DBDSN)
	if err != nil {
		errorLog.Fatalf("Error connecting to database: %v", err)
	}
	defer pool.Close()

	//5. check if db connection is alive
	ctxPing, cancelPing := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelPing()
	if err := pool.Ping(ctxPing); err != nil {
		errorLog.Fatalf("Database ping failed: %v", err)
	}
	infoLog.Println("🗄️  Successfully connected to PostgreSQL!")
}
