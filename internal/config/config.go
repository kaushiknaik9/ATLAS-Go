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
	Host    string        `yaml:"Host" env-required:"true" env-default:"localhost:2007"`
	Timeout time.Duration `yaml:"Timeout" env-default:"10s"`
}

type Config struct {
	Env         string `yaml:"ENV" env-default:"production"`
	StoragePath string `yaml:"Storage_Path" env-required:"true" env-default:"storage/storage.db"`
	HttpServer  `yaml:"Http_Server" env-required:"true"`
}

func ConfigImported() *Config {
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
