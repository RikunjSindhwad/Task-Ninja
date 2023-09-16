package utils

import (
	"strings"

	"github.com/RikunjSindhwad/Task-Ninja/pkg/config"
	"github.com/RikunjSindhwad/Task-Ninja/pkg/visuals"

	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
)

type Args struct {
	Workflow string
	YamlVars goflags.StringSlice
	Json     bool
	NoBanner bool
}

func ParseArgs() *Args {
	options := &Args{}
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`Task-Ninja is workflow based tasks execution framework.`)
	flagSet.StringVarP(&options.Workflow, "workflow", "w", "", "YAML Workflow Path")
	flagSet.BoolVar(&options.Json, "json", false, "Json Log (default false)")
	flagSet.BoolVarP(&options.NoBanner, "noBanner", "nb", false, "Do not print banner (default false)")
	flagSet.StringSliceVarP(&options.YamlVars, "vars", "v", nil, "yaml variables and values '<var=value,var2=value2>'", goflags.FileCommaSeparatedStringSliceOptions)
	_ = flagSet.Parse()

	if options.Workflow == "" {
		visuals.PrintBanner()
		gologger.Fatal().Msg("Usage: TaskNinja -workflow <workflow.yaml> -debug")
	}

	return options
}

func UpdateConfigStruct(configStruct *config.Config, keyValueList []string) error {

	for _, pair := range keyValueList {
		// Split the pair into key and value
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			// Invalid pair, skip
			continue
		}

		key := kv[0]
		value := kv[1]

		// Update the Vars field in the Config struct with the new key-value pair
		configStruct.Vars[key] = value
	}
	isempty, emptyVar := checkEmptyVars(configStruct)
	if !isempty {
		// visuals.PrintState("FETAL")
		gologger.Error().Label("ERROR").Str("Variables", strings.Join(emptyVar, ",")).Msg("Missing Variable Value")
		gologger.Fatal().Label("USAGE").Msg(configStruct.WorkflowConfig.Usage)
	}

	return nil
}

func checkEmptyVars(configStruct *config.Config) (bool, []string) {
	emptyVars := make([]string, 0)

	for key, value := range configStruct.Vars {
		if value == "" {
			emptyVars = append(emptyVars, key)
		}
	}

	if len(emptyVars) > 0 {
		return false, emptyVars
	}
	return true, emptyVars
}
