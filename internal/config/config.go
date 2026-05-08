package config

import (
	"ecom/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ServerConfig struct {
	RunEnv types.Env `mapstructure:"run_env" validate:"required,oneof=dev stg prod"`
	Port   string    `mapstructure:"port" validate:"required,numeric"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required,numeric"`
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Database string `mapstructure:"database" validate:"required"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address" validate:"required"`
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Database int    `mapstructure:"database" validate:"gte=0"`
}

type ElasticSearchConfig struct {
	Address  string `mapstructure:"address" validate:"required"`
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
}

type HashingConfig struct {
	Memory      uint32 `mapstructure:"memory" validate:"required,min=1"`
	Iterations  uint32 `mapstructure:"iterations" validate:"required,min=1"`
	Parallelism uint8  `mapstructure:"parallelism" validate:"required,min=1"`
	SaltLength  uint32 `mapstructure:"salt_length" validate:"required,min=16"`
	KeyLength   uint32 `mapstructure:"key_length" validate:"required,min=16"`
}

type JWTConfig struct {
	SecretKey string `mapstructure:"secret_key" validate:"required"`
}

type Config struct {
	Server        *ServerConfig        `mapstructure:"server" validate:"required"`
	Postgres      *PostgresConfig      `mapstructure:"postgres" validate:"required"`
	Redis         *RedisConfig         `mapstructure:"redis" validate:"required"`
	ElasticSearch *ElasticSearchConfig `mapstructure:"elastic_search" validate:"required"`
	Hashing       *HashingConfig       `mapstructure:"hashing" validate:"required"`
	JWT           *JWTConfig           `mapstructure:"jwt" validate:"required"`
}

func (c *Config) Validate() error {
	return validator.New().Struct(c)
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		zap.L().Error(
			"failed to read config file",
			zap.Error(err),
		)
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		zap.L().Error(
			"failed to unmarshal config",
			zap.Error(err),
		)
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		zap.L().Error(
			"config validation failed",
			zap.Error(err),
		)
		return nil, err
	}

	zap.L().Info("successfully loaded config")

	return &cfg, nil
}
