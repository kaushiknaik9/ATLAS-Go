package config

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Host    string        `yaml:"host" env-required:"true" env-default:"localhost:2007"`
	Timeout time.Duration `yaml:"timeout" env-default:"10s"`
}

type Config struct {
	Env         string `yaml:"env" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true" env-default:"storage/storage.db"`
	HttpServer  `yaml:"http_server" env-required:"true"`
}

func LoadConfig() *Config {
	var ConfigPath string
	ConfigPath = os.Getenv("CONFIG_PATH")

	if ConfigPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()
		ConfigPath = *flags

		if ConfigPath == "" {
			log.Fatal("Config file not found in env or wasnt provided during the start of server")
		}
	}

	_, err := os.Stat(ConfigPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File Does not exists")
		}
		log.Fatal("Error While Locating File")
	}

	var cfg Config
	err = cleanenv.ReadConfig(ConfigPath, &cfg)
	if err != nil {
		log.Fatal("Error while reading the file. There nothing exists in the config file")
	}

	slog.Info("Config File Found and it is Readable")

	return &cfg
}
