package main

import (
	"time"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/config"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/executors"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/utils"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/visuals"
)

func main() {
	visuals.SetLevelInfo()
	args := utils.ParseArgs()
	if !args.NoBanner {
		visuals.PrintBanner()
	}

	if args.Json {
		visuals.JsonView()
	}
	if args.Detailed {
		visuals.SetLevelDebug()
	}

	yamlFilePath := args.Workflow

	configStruct := config.ReadYamlFromFile(yamlFilePath)
	startTime := time.Now()
	if configStruct.WorkflowConfig.Author != "" {

		visuals.PrintCredit(configStruct.WorkflowConfig.Author, configStruct.WorkflowConfig.Name, "start", "")
	}

	variables := args.YamlVars
	utils.UpdateConfigStruct(configStruct, variables)
	utils.ReplacePlaceholders(configStruct)
	executors.ExecHelper(configStruct)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	visuals.PrintCredit(configStruct.WorkflowConfig.Author, configStruct.WorkflowConfig.Name, "end", duration.String())
}
