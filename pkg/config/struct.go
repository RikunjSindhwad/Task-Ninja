package config

type Task struct {
	Name           string   `yaml:"name"`
	Image          string   `yaml:"image"`
	DockerHive     string   `yaml:"dockerhive"`
	Cmds           []string `yaml:"cmds"`
	Silent         bool     `yaml:"silent"`
	Parallel       bool     `yaml:"parallel"`
	UseHostNetwork bool     `yaml:"hostnetwork"`
	Required       []string `yaml:"required"`
	Timeout        int      `yaml:"timeout"`
	StoponError    bool     `yaml:"stoponerr"`
	Type           string   `yaml:"type"`
	DynamicFile    string   `yaml:"dynamicFile"`
	DynamicRange   string   `yaml:"dynamicRange"`
	MaxThreads     int      `yaml:"threads"`
	InputMouts     []string `yaml:"mounts"`
	Inputs         []string `yaml:"inputs"`
}

type WorkflowConfig struct {
	Name               string `yaml:"name"`
	Author             string `yaml:"author"`
	Usage              string `yaml:"usage"`
	Shell              string `yaml:"shell"`
	DefaultDockerimage string `yaml:"defaultimage"`
	DefaultHive        string `yaml:"hive"`
	EnableLogs         bool   `yaml:"logs"`
}

type Config struct {
	WorkflowConfig WorkflowConfig    `yaml:"config"`
	Vars           map[string]string `yaml:"vars"`
	Tasks          []Task            `yaml:"tasks"`
}
