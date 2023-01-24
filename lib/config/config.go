package config

import (
	"fmt"
	"os"

	"github.com/energywork/pseudo-paysystem/lib/errs"
	"gopkg.in/yaml.v3"
)

type Config struct {
	API struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"api"`
	NATS struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"nats"`
	SQL struct {
		Type     string `yaml:"type"` // eg postgers/sqlite/mysql
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"sql"`

	DEV bool `yaml:"dev"`
}

func LoadConfig(path string) (*Config, *errs.Error) {
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
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=false TimeZone=Europe/Moscow", c.SQL.Host, c.SQL.User, c.SQL.Password, c.SQL.Database, c.SQL.Port)
}
