package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Server string
	DB     Database
	Redis  Redis
}

type Database struct {
	UserName  string
	Password  string
	IpAddress string
	Port      int
	DBName    string
}

type Redis struct {
	Addr     string
	Password string
	DB       int
}

func LoadDatabase() {
	var config TomlConfig
	if _, err := toml.DecodeFile("../../configs/server/config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

}
