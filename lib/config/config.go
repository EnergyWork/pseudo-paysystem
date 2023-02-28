package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v7"
	"gopkg.in/yaml.v3"

	"github.com/energywork/pseudo-paysystem/lib/errs"
)

type Config struct {
	Service string `env:"-"`

	// Host and Port only for http
	HttpHost        string `yaml:"host" env:"HTTP_HOST" envDefault:"localhost"`
	HttpPort        string `yaml:"port" env:"HTTP_PORT" envDefault:"2222"`
	ReadTimeout     int    `yaml:"readTimeout" env:"HTTP_READ_TIMEOUT" envDefault:"5"`         // seconds
	WriteTimeout    int    `yaml:"writeTimeout" env:"HTTP_WRITE_TIMEOUT" envDefault:"5"`       // seconds
	ShutdownTimeout int    `yaml:"shutdownTimeout" env:"HTTP_SHUTDOWN_TIMEOUT" envDefault:"3"` // seconds

	// NatsHost and NatsPort ony for nats
	NatsHost string `yaml:"host" env:"NATS_HOST" envDefault:"localhost"`
	NatsPort string `yaml:"port" env:"NATS_PORT" envDefault:"4222"`

	// SQL parameters
	SqlType     string `yaml:"type" env:"SQL_TYPE" envDefault:"postgres"` // eg postges/sqlite/mysql
	SqlHost     string `yaml:"host" env:"SQL_HOST" envDefault:"localhost"`
	SqlPort     string `yaml:"port" env:"SQL_PORT" envDefault:"5432"`
	SqlUser     string `yaml:"user" env:"SQL_USER" envDefault:"postgres"`
	SqlPassword string `yaml:"password" env:"SQL_PASSWORD" envDefault:"postgres"`
	SqlDatabase string `yaml:"database" env:"SQL_DATABASE" envDefault:"postgres"`
	SqlSSLMode  string `yaml:"sslmode" env:"SQL_SSLMODE" envDefault:"disable"`

	DEV bool `yaml:"dev" env:"DEV" envDefault:"false"`
}

func New(service string) *Config {
	cfg := &Config{Service: service}
	if err := env.Parse(cfg); err != nil {
		return nil
	}
	return cfg
}

// LoadYAMLConfig returns a config from yaml file
//
// todo: implement a universal loader
func LoadYAMLConfig(path string) (*Config, *errs.Error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, errs.New().SetCode(errs.Internal).SetMsg(err.Error())
	}
	c := &Config{}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return nil, errs.New().SetCode(errs.Internal).SetMsg(err.Error())
	}
	return c, nil
}

func (c *Config) PostgresDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Moscow", c.SqlHost, c.SqlUser, c.SqlPassword, c.SqlDatabase, c.SqlPort, c.SqlSSLMode)
}
