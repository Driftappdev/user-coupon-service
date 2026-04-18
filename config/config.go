package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {

	// ------------------------------------------------
	// App
	// ------------------------------------------------

	App struct {
		Name string `yaml:"name"`
	} `yaml:"app"`

	// ------------------------------------------------
	// HTTP
	// ------------------------------------------------

	HTTP struct {
		Port string `yaml:"port"`
	} `yaml:"http"`

	// ------------------------------------------------
	// GRPC
	// ------------------------------------------------

	GRPC struct {
		Address string `yaml:"address"`
	} `yaml:"grpc"`

	// ------------------------------------------------
	// Database
	// ------------------------------------------------

	Database struct {
		DSN             string        `yaml:"dsn"`
		MaxOpenConns    int           `yaml:"max_open_conns"`
		MaxIdleConns    int           `yaml:"max_idle_conns"`
		ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
	} `yaml:"database"`

	// ------------------------------------------------
	// NATS JetStream
	// ------------------------------------------------

	NATS struct {
		URL     string `yaml:"url"`
		Stream  string `yaml:"stream"`
		Subject string `yaml:"subject"`
	} `yaml:"nats"`
}

func LoadConfig() *Config {

	cfg := &Config{}

	//------------------------------------------------
	// Load YAML
	//------------------------------------------------

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "config.yaml"
	}

	data, err := os.ReadFile(path)
	if err != nil {

		log.Printf(
			"[Config] cannot read %s : %v (fallback env/default)",
			path,
			err,
		)

	} else {

		if err := yaml.Unmarshal(data, cfg); err != nil {
			log.Fatalf("[Config] yaml parse error: %v", err)
		}

	}

	//------------------------------------------------
	// ENV override
	//------------------------------------------------

	if v := os.Getenv("HTTP_PORT"); v != "" {
		cfg.HTTP.Port = v
	}

	if v := os.Getenv("GRPC_ADDRESS"); v != "" {
		cfg.GRPC.Address = v
	}

	if v := os.Getenv("DATABASE_DSN"); v != "" {
		cfg.Database.DSN = v
	}

	if v := os.Getenv("NATS_URL"); v != "" {
		cfg.NATS.URL = v
	}

	if v := os.Getenv("NATS_STREAM"); v != "" {
		cfg.NATS.Stream = v
	}

	if v := os.Getenv("NATS_SUBJECT"); v != "" {
		cfg.NATS.Subject = v
	}

	//------------------------------------------------
	// Defaults
	//------------------------------------------------

	if cfg.App.Name == "" {
		cfg.App.Name = "user-coupon-service"
	}

	if cfg.HTTP.Port == "" {
		cfg.HTTP.Port = "8080"
	}

	if cfg.GRPC.Address == "" {
		cfg.GRPC.Address = ":50051"
	}

	if cfg.Database.MaxOpenConns == 0 {
		cfg.Database.MaxOpenConns = 50
	}

	if cfg.Database.MaxIdleConns == 0 {
		cfg.Database.MaxIdleConns = 10
	}

	if cfg.Database.ConnMaxLifetime == 0 {
		cfg.Database.ConnMaxLifetime = 30 * time.Minute
	}

	if cfg.NATS.URL == "" {
		cfg.NATS.URL = "nats://localhost:4222"
	}

	if cfg.NATS.Stream == "" {
		cfg.NATS.Stream = "COUPON_STREAM"
	}

	if cfg.NATS.Subject == "" {
		cfg.NATS.Subject = "coupon.events"
	}

	//------------------------------------------------
	// Log Summary
	//------------------------------------------------

	log.Printf(
		"[Config] app=%s http=%s grpc=%s db=%s nats=%s stream=%s subject=%s",
		cfg.App.Name,
		cfg.HTTP.Port,
		cfg.GRPC.Address,
		cfg.Database.DSN,
		cfg.NATS.URL,
		cfg.NATS.Stream,
		cfg.NATS.Subject,
	)

	return cfg
}
