package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server  ServerConfig
	Logger  LoggerConfig
	Swagger SwaggerConfig
	MySQL   MySQLConfig
}

type ServerConfig struct {
	AppVersion   string
	Port         string
	PprofPort    string
	Mode         string
	JwtSecretKey string
	// JWTExpiredTime time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	SSL          bool
	Debug        bool
	CSRF         bool
}

type LoggerConfig struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

type SwaggerConfig struct {
	Title       string
	Description string
	Version     string
	BasePath    string
	Host        string
}

type MySQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Driver   string
	SSLMode  bool
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found, ")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfigDefault(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
