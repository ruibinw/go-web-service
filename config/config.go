package config

import (
	"git.epam.com/ryan_wang/go-web-service/internal/models"
	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/sqlite"
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
	//open database connection
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	//create Record table
	db.AutoMigrate(&models.Record{})
	return db
}
