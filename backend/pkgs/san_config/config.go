package san_config

import (
	"os"
	"pkl/backend/configs"

	"gopkg.in/yaml.v3"
)

func RenderYamlConfig() (configs.AppConfig, error) {
	yamlConfig, err := os.ReadFile("config.yaml")
	if err != nil {
		return configs.AppConfig{}, err
	}

	var config configs.AppConfig
	err = yaml.Unmarshal(yamlConfig, &config)

	return config, nil
}
