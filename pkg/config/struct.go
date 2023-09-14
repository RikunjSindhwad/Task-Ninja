package config

type Task struct {
	Name         string   `yaml:"name"`
	Cmds         []string `yaml:"cmds"`
	Silent       bool     `yaml:"silent"`
	Parallel     bool     `yaml:"parallel"`
	Required     []string `yaml:"required"`
	Timeout      int      `yaml:"timeout"`
	StoponError  bool     `yaml:"stoponerr"`
	Type         string   `yaml:"type"`
	DynamicFile  string   `yaml:"dynamicFile"`
	DynamicRange string   `yaml:"dynamicRange"`
	MaxThreads   int      `yaml:"threads"`
}

type WorkflowConfig struct {
	Name       string `yaml:"name"`
	Author     string `yaml:"author"`
	Usage      string `yaml:"usage"`
	Shell      string `yaml:"shell"`
	StdeoutDir string `yaml:"stdoutDir"`
	StderrDir  string `yaml:"stderrDir"`
}

type Config struct {
	WorkflowConfig WorkflowConfig    `yaml:"config"`
	Vars           map[string]string `yaml:"vars"`
	Tasks          []Task            `yaml:"tasks"`
}
