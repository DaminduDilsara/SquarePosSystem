package configurations

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	AppConfig    *AppConfigurations    `yaml:"app_config"`
	SquareConfig *SquareConfigurations `yaml:"square_config"`
}

func LoadConfigurations() *Config {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Loading config yaml err: %v ", err)
	}

	var configs Config
	err = yaml.Unmarshal(yamlFile, &configs)
	if err != nil {
		log.Fatalf("Unmarshal error: %v", err)
	}

	return &configs
}
