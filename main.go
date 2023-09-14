package main

import (
	"Robensive-TaskNinja/pkg/config"
	"Robensive-TaskNinja/pkg/executors"
	"Robensive-TaskNinja/pkg/utils"
	"Robensive-TaskNinja/pkg/visuals"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/formatter"
	"github.com/projectdiscovery/gologger/levels"
)

func main() {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	args := utils.ParseArgs()
	if !args.NoBanner {
		visuals.PrintBanner()
	}

	if args.Json {
		gologger.DefaultLogger.SetFormatter(&formatter.JSON{})
	}

	yamlFilePath := args.Workflow

	configStruct, err := config.ReadYamlFromFile(yamlFilePath)
	if err != nil {
		gologger.Fatal().Label("ERROR").Msgf("Error Reading YAML: %v", err)

	}

	if configStruct.WorkflowConfig.Author != "" {
		// Format string in advance to add author name in string
		gologger.Info().Label("Workflow-Credit").Str("Workflow-Author", configStruct.WorkflowConfig.Author).Msgf("Tasked Workflow '%s'", configStruct.WorkflowConfig.Name)
	}
	// Replace Placeholders {{#}}
	variables := args.YamlVars
	utils.UpdateConfigStruct(configStruct, variables)
	utils.ReplacePlaceholders(configStruct)
	// updatedConfigData, err := yaml.Marshal(&configStruct)
	// if err != nil {

	// }
	// println(string(updatedConfigData))
	executors.ExecHelper(configStruct)
	gologger.Info().Label("Workflow-Complete").Str("Workflow-Author", configStruct.WorkflowConfig.Author).Msgf("Workflow '%s' Execution Complete", configStruct.WorkflowConfig.Name)

}
