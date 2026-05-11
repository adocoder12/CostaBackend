package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all environment-driven configuration for the application.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Supabase SupabaseConfig
	AWS      AWSConfig
	Upload   UploadConfig
	CORS     CORSConfig
}

type ServerConfig struct {
	Port    string
	GinMode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	MaxConns int32
	MinConns int32
}

type SupabaseConfig struct {
	JWTSecret string // from Supabase dashboard → Settings → API → JWT Secret
	URL       string // your Supabase project URL
}

type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	S3Bucket        string
	S3Endpoint      string // LocalStack endpoint in dev, empty in prod
	SQSQueueURL     string // SES.HOSPEDAJES transmission queue
}

type UploadConfig struct {
	MaxFileSize    int64  // bytes
	UploadProvider string // "s3" or "local"
}

type CORSConfig struct {
	AllowedOrigin string
}

// Load reads all environment variables and returns a Config.
// Tolerates a missing .env file (uses system env instead).
func Load() (*Config, error) {
	// Tolerant load — no error if .env is missing (production uses system env)
	_ = godotenv.Load()

	maxConns, _ := strconv.ParseInt(getEnv("DB_MAX_CONNS", "25"), 10, 32)
	minConns, _ := strconv.ParseInt(getEnv("DB_MIN_CONNS", "5"), 10, 32)
	maxUploadSize, _ := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "10485760"), 10, 64) // 10MB default

	cfg := &Config{
		Server: ServerConfig{
			Port:    getEnv("PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "admin"),
			Password: getEnv("DB_PASSWORD", "password"),
			Name:     getEnv("DB_NAME", "costaBackend"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			MaxConns: int32(maxConns),
			MinConns: int32(minConns),
		},
		Supabase: SupabaseConfig{
			JWTSecret: getEnv("SUPABASE_JWT_SECRET", ""),
			URL:       getEnv("SUPABASE_URL", ""),
		},
		AWS: AWSConfig{
			Region:          getEnv("AWS_REGION", "eu-central-1"),
			AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", "test"),
			SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", "test"),
			S3Bucket:        getEnv("S3_BUCKET", "costa-pms-uploads"),
			S3Endpoint:      getEnv("AWS_ENDPOINT", ""), // set to http://localstack:4566 in dev
			SQSQueueURL:     getEnv("SQS_QUEUE_URL", ""),
		},
		Upload: UploadConfig{
			MaxFileSize:    maxUploadSize,
			UploadProvider: getEnv("UPLOAD_PROVIDER", "s3"),
		},
		CORS: CORSConfig{
			AllowedOrigin: getEnv("ALLOWED_ORIGIN", "http://localhost:3000"),
		},
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// DSN returns the PostgreSQL connection string from the database config.
func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode,
	)
}

func (d DatabaseConfig) MigrationURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		d.User, d.Password, d.Host, d.Port, d.Name, d.SSLMode,
	)
}

// validate checks that all required variables are set.
// Fails fast at startup rather than panicking mid-request.
func (c *Config) validate() error {
	if c.Supabase.JWTSecret == "" {
		return fmt.Errorf("config: SUPABASE_JWT_SECRET is required")
	}
	if c.Database.Host == "" {
		return fmt.Errorf("config: DB_HOST is required")
	}
	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
