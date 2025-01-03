package configs

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type rabbitMQ struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	VHost    string `yaml:"vhost"`
}

type Configs struct {
	RabbitMQ rabbitMQ `yaml:"rabbitmq"`
}

func NewConfigs() *Configs {
	file, err := os.Open("environment.yaml")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	var cfg Configs
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("Failed to decode YAML file: %v", err)
	}

	return &cfg
}
