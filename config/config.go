package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"path"
	"runtime"
)

type Configuration struct {
	Server struct {
		Port string `yaml:"port"    env:"SERVER_PORT"    env-default:"8080"`
	} `yaml:"server"`

	Database struct {
		Which    string `yaml:"d"     env:"DB_HOST"      env-default:"localhost"`
		Host     string `yaml:"host"     env:"DB_HOST"      env-default:"localhost"`
		Port     string `yaml:"port"     env:"DB_PORT"      env-default:"3306"`
		UserName string `yaml:"username" env:"DB_USERNAME"  env-default:"root"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		Name     string `yaml:"db-name"  env:"DB_NAME"`
	} `yaml:"database"`

	Profile string `yaml:"profile" env:"ACTIVE_PROFILE" env-default:"dev"`
}

func GetConfig() *Configuration {
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

func (cfg *Configuration) OpenDBConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.UserName,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	return db

}
