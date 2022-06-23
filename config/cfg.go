package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	ServerHost string
	DataBase   string
}

func SetConfig() (cfg Config, err error) {
	var cfgData []byte
	if cfgData, err = os.ReadFile("config/cfg.toml"); err == nil {
		_, err = toml.Decode(string(cfgData), &cfg)
		if err != nil {
			return
		}
	}
	return cfg, err
}
