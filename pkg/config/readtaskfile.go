package config

import (
	"os"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/visuals"
	"gopkg.in/yaml.v2"
)

func ReadYamlFromFile(filePath string) *Config {
	var configStruct Config

	yamlContent, err := os.ReadFile(filePath)
	if err != nil {
		visuals.PrintState("workflow-error", "", "Reading YAML File: "+err.Error())
		return nil
	}

	err = yaml.Unmarshal(yamlContent, &configStruct)
	if err != nil {
		visuals.PrintState("workflow-error", "", "Parsing YAML File: "+err.Error())
		return nil
	}

	return &configStruct
}
