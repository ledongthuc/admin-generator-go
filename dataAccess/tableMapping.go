package dataAccess

import (
	"io/ioutil"

	"github.com/jbrodriguez/mlog"
	"github.com/ledongthuc/admin-generator-go/entity"
	"gopkg.in/yaml.v2"
)

// TableDataAccess contains methods that used for access to `information_schema.tables`.
type tableMapping struct {
}

// Table instance of Table Data Access
var TableMapping tableMapping

// Load loads settings from /conf/settings.yml
func (dataAccess *tableMapping) Load() (map[string]entity.TableMapping, error) {
	data, _ := ioutil.ReadFile("conf/table_mappings.yml")

	tableMappings := map[string]entity.TableMapping{}
	err := yaml.Unmarshal([]byte(data), &tableMappings)
	return tableMappings, err
}

func (dataAccess *tableMapping) IsShowTable(tableName string) bool {
	mappings, err := dataAccess.Load()
	if err != nil {
		mlog.Error(err)
		return false
	}

	tableMapping, existed := mappings[tableName]
	if !existed || tableMapping.IsShow {
		return true
	}

	return false
}
