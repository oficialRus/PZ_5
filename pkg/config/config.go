package config

import (
	"github.com/spf13/viper"
)

type (
	DBCfg struct {
		Driver string
		DSN    string
	}
	ServerCfg struct {
		Port int
	}
	Config struct {
		DB     DBCfg
		Server ServerCfg
	}
)

func Load(path string) (Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	// allow override via env vars: WEBAPP_DB_DSN, WEBAPP_SERVER_PORT …
	v.SetEnvPrefix("WEBAPP")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
