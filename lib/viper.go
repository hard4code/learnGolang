package lib

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}

type Mysql struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname   string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	LogMode  bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

func Yaml() {
	confFile := "config.yaml"
	vip := viper.New()
	vip.SetConfigFile(confFile)

	err := vip.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error config.yaml: %s \n", err))
	}

	var config Config
	if err := vip.Unmarshal(&config); err != nil {
		fmt.Println("error config",err)
	}

	fmt.Println(config.Mysql.Path)
}
