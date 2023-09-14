package config

import (
	// ... (other imports)

	"os"

	"gopkg.in/yaml.v2"
)

func ReadYamlFromFile(filePath string) (*Config, error) {
	var configStruct Config
	yamlContent, err := os.ReadFile(filePath)
	if err != nil {

		return nil, err
	}

	err = yaml.Unmarshal(yamlContent, &configStruct)
	if err != nil {
		// gologger.Fatal().Msgf("Error unmarshaling YAML: %v", err)
		return nil, err
	}

	return &configStruct, nil
}
