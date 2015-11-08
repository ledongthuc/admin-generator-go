package helpers

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Setting struct {
	Database struct {
		ConnectionString string `yaml:"connection-string"`
		Type             string `yaml:"type"`
	} `yaml:"database"`
}

// LoadSettings loads settings from /conf/settings.yml
func LoadSettings() (Setting, error) {
	data, _ := ioutil.ReadFile("conf/settings.yml")

	setting := Setting{}
	err := yaml.Unmarshal([]byte(data), &setting)
	return setting, err
}
