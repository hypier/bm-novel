package config

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/BurntSushi/toml"
)

// Config 配置
var Config = &tomlConfig{}

type tomlConfig struct {
	Server string   `toml:"listen"`
	DB     database `toml:"database"`
	Redis  redis    `toml:"redis"`
}

// database 配置
type database struct {
	UserName  string `toml:"userName"`
	Password  string `toml:"password"`
	IPAddress string `toml:"ipAddress"`
	Port      int    `toml:"port"`
	DBName    string `toml:"dbName"`
}

// Redis redis配置
type redis struct {
	IPAddress string `toml:"ipAddress"`
	Password  string `toml:"password"`
	DB        int    `toml:"db"`
}

// LoadConfig 载入配置文件
func LoadConfig() {
	if Config.Server != "" {
		return
	}

	fileName := "./configs/server/config.toml"
	if _, err := toml.DecodeFile(fileName, &Config); err != nil {
		_ = errors.WithStack(err)
		fmt.Println(err)
	}
}
