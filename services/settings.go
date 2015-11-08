package services

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/ledongthuc/admin-generator-go/entity"
)

// settings is defination of service settings. Use to load and manage setting
type settings struct{}

// Settings instance of Settings services
var Settings settings

// Load loads settings from /conf/settings.yml
func (settings *settings) Load() (entity.Setting, error) {
	data, _ := ioutil.ReadFile("conf/settings.yml")

	setting := entity.Setting{}
	err := yaml.Unmarshal([]byte(data), &setting)
	return setting, err
}
