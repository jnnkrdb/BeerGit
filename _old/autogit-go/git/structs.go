package git

type Git struct {
	Collection []GitCollection `yaml:"git"`
}

type GitCollection struct {
	Collection string          `yaml:"collection"`
	Directory  string          `yaml:"dir"`
	Repos      []GitRepository `yaml:"repos"`
}

type GitRepository struct {
	Name   string `yaml:"name"`
	URL    string `yaml:"url"`
	Branch string `yaml:"branch"`
	Path   string `yaml:"path"`
}
