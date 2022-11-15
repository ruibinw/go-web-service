package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"path"
	"runtime"
)

type Configuration struct {
	Server struct {
		Port string `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"     env:"DB_HOST"      env-default:"localhost"`
		Port     string `yaml:"port"     env:"DB_PORT"      env-default:"3306"`
		UserName string `yaml:"username" env:"DB_USERNAME"  env-default:"root"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		Name     string `yaml:"db-name"  env:"DB_NAME"`
	} `yaml:"database"`
}

func LoadConfig() *Configuration {
	var config Configuration
	configFile := getSourcePath() + "/../config.yml"
	if err := cleanenv.ReadConfig(configFile, &config); err != nil {
		panic(err)
	}
	return &config
}

func getSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
