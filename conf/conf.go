package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DSN   string
	MySQL struct {
		Root      string `yaml:"root"`
		Password  string `yaml:"password"`
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		DB        string `yaml:"db"`
		Charset   string `yaml:"charset"`
		ParseTime string `yaml:"parseTime"`
		Loc       string `yaml:"loc"`
	} `yaml:"mysql"`
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"DB"`
	} `yaml:"redis"`
}

func Conf() Config {
	file, err := os.ReadFile("conf/dbconf.yaml")

	if err != nil {
		panic(err)
	}
	var Config1 Config
	err = yaml.Unmarshal(file, &Config1)
	if err != nil {
		panic(err)
	}

	Config1.DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		Config1.MySQL.Root,
		Config1.MySQL.Password,
		Config1.MySQL.Host,
		Config1.MySQL.Port,
		Config1.MySQL.DB,
		Config1.MySQL.Charset,
		Config1.MySQL.ParseTime,
		Config1.MySQL.Loc)
	return Config1
}
