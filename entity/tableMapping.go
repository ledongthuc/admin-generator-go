package entity

// TableMapping uses to mapping setting from /conf/table_mappings.yml
type TableMapping struct {
	IsShow  bool `yaml:"is_show"`
	Columns map[string]struct {
		IsShow string `yaml:"is_show"`
	}
}
