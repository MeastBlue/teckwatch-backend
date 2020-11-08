package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Database struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Name string `mapstructure:"name"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"password"`
}

type Config struct {
	Server Server   `mapstructure:"server"`
	Db     Database `mapstructure:"database"`
}

func InitConfig() (*Config, error) {
	v := viper.New()
	c := Config{}

	v.SetConfigType("yaml")
	v.SetConfigName("env")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		errString := fmt.Sprintf("Couldn't load config: %s", err)
		return nil, errors.New(errString)
	}

	if err := v.Unmarshal(&c); err != nil {
		errString := fmt.Sprintf("Couldn't read config: %s", err)
		return nil, errors.New(errString)
	}

	return &c, nil
}

func (c *Config) GetServerConf() *Server {
	return &c.Server
}

func (c *Config) GetDatabaseConf() *Database {
	return &c.Db
}
