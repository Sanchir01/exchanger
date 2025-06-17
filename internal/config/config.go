package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env     string   `yaml:"env"`
	Domain  string   `yaml:"domain"`
	GRPC    GRPC     `yaml:"grpc" required:"true"`
	RedisDB Redis    `yaml:"redis"`
	DB      DataBase `yaml:"database"`
}
type Kafka struct {
	Outbox Producer `yaml:"outbox"`
}

type Producer struct {
	Retries int      `yaml:"retries"`
	Topic   []string `yaml:"topic"`
	Broke   []string `yaml:"brokers"`
}
type Redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Retries  int    `yaml:"retries"`
	DBNumber int    `yaml:"dbnumber"`
}

type DataBase struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	User        string `yaml:"user"`
	Database    string `yaml:"dbname"`
	SSL         string `yaml:"ssl"`
	MaxAttempts int    `yaml:"max_attempts"`
}
type GRPC struct {
	Port    string        `yaml:"port" required:"true"`
	Host    string        `yaml:"host" required:"true"`
	Timeout time.Duration `yaml:"timeout" required:"true"`
}

func InitConfig() *Config {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env.dev"
	}
	fmt.Println("env name", envFile)
	if err := godotenv.Load(envFile); err != nil {
		slog.Error("ошибка при инициализации переменных окружения", err.Error())
	}
	configPath := os.Getenv("CONFIG_PATH")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist:%s", configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	return &cfg
}
