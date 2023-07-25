package models

import (
	"github.com/spf13/viper"
)

type Config struct {
	SqlServer SqlServer `yaml:"sqlserver"`
	Postgre   Postgre   `yaml:"postgress"`
	Api       Api       `yaml:"api"`
}

type SqlServer struct {
	Server   string `yaml:"server"`
	DbName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Postgre struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	DbName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Api struct {
	Port string `yaml:"port"`
}

func (config *Config) GetConfigValues() {
	v := viper.New()
	v.SetTypeByDefaultValue(true)
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./resources")
	err := v.ReadInConfig()
	if err != nil {
		panic("Config not found :\r\n" + err.Error())
	}
	sub := v.Sub("local")
	unMarshallErr := sub.Unmarshal(config)

	if unMarshallErr != nil {
		panic("Unmarshall error")
	}
}
