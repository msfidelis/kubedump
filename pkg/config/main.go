package config

// Model is a configuration schema from kubedump files
type Model struct {
	Project    string   `yaml:"project"`
	Format     string   `yaml:"format"`
	Namespaces []string `yaml:"namespaces"`
	Resources  []string `yaml:"resources"`
}
