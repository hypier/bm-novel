package config

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

// Config 配置
var Config = &tomlConfig{}

type tomlConfig struct {
	Server   string   `toml:"listen"`
	LogLevel string   `toml:"logLevel"`
	DB       database `toml:"database"`
	Redis    rdb      `toml:"rdb"`
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
type rdb struct {
	IPAddress string `toml:"ipAddress"`
	Password  string `toml:"password"`
	DB        int    `toml:"db"`
}

// LoadConfig 载入配置文件
func LoadConfig(fileName string) {
	if Config.Server != "" {
		return
	}

	if _, err := toml.DecodeFile(fileName, &Config); err != nil {
		err = errors.WithStack(err)
		panic(err)
		//fmt.Printf("%+v", err)
	}
}

// LoadConfigForTest 测试配置加载
func LoadConfigForTest() {
	LoadConfig("E:\\GoCode\\src\\bm-novel\\configs\\server\\config.toml")
}
