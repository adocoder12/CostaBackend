package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	DBDSN             string
	SupabaseJWTSecret string
	AWSEndpoint       string
	AWSRegion         string
	S3Bucket          string
	SQSQueueURL       string
}

func Load() *Config {
	// We ignore error here because in production .env might not exist
	// (variables would be set via Docker/Kubernetes directly)
	_ = godotenv.Load()

	cfg := &Config{
		Port:              getEnv("PORT", "8080"),
		SupabaseJWTSecret: getEnv("SUPABASE_JWT_SECRET", ""),
		AWSEndpoint:       getEnv("AWS_ENDPOINT", "http://localhost:4566"),
		AWSRegion:         getEnv("AWS_REGION", "eu-central-1"),
		S3Bucket:          getEnv("S3_BUCKET", "costa-pms-uploads"),
		SQSQueueURL:       getEnv("SQS_QUEUE_URL", ""),
	}

	cfg.DBDSN = buildDSN()

	if cfg.SupabaseJWTSecret == "" {
		log.Fatal("SUPABASE_JWT_SECRET is required for authentication")
	}

	return cfg
}

func buildDSN() string {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "admin")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "costaBackend")
	sslmode := getEnv("DB_SSLMODE", "disable")

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
