package config

type Model struct {
	Project    string   `yaml:"project"`
	Format     string   `yaml:"format"`
	Namespaces []string `yaml:"namespaces"`
	Resources  []string `yaml:"resources"`
}
