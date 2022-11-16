package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"path"
	"runtime"
)

var config *Configuration

type Configuration struct {
	Server struct {
		Host string `yaml:"host"    env:"SERVER_HOST"    env-default:"localhost"`
		Port string `yaml:"port"    env:"SERVER_PORT"    env-default:"8080"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"     env:"DB_HOST"      env-default:"localhost"`
		Port     string `yaml:"port"     env:"DB_PORT"      env-default:"3306"`
		UserName string `yaml:"username" env:"DB_USERNAME"  env-default:"root"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		Name     string `yaml:"db-name"  env:"DB_NAME"`
	} `yaml:"database"`

	Swagger struct {
		Host string `yaml:"host"    env:"SWAGGER_HOST"    env-default:""`
	} `yaml:"swagger"`
}

func GetConfig() *Configuration {
	if config == nil {
		config = &Configuration{}
		configFile := getSourcePath() + "/../config.yml"
		if err := cleanenv.ReadConfig(configFile, config); err != nil {
			panic(err)
		}
	}
	return config
}

func getSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func NewDBConnection() *gorm.DB {
	cfg := GetConfig()
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
