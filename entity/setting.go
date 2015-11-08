package entity

// Use to mapping setting from /conf/setting.yml
type Setting struct {
	Database struct {
		ConnectionString string `yaml:"connection-string"`
		Type             string `yaml:"type"`
	} `yaml:"database"`
}
