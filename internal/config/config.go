package config

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
	"strings"
)

type Config struct {
	DB DB
}

type DB struct {
	Host     string
	Port     int64
	Database string
	User     string
	Password string
}

func (d *DB) Dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", d.Host, d.User, d.Password, d.Database, d.Port)
}

func Load(cfgLocation string) Config {
	var k = koanf.New(".")
	if err := k.Load(file.Provider(cfgLocation), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	k.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "")), "_", ".", -1)
	}), nil)

	var cfg Config
	k.Unmarshal("", &cfg)
	return cfg
}
