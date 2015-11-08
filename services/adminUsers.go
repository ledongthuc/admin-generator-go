package services

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/jbrodriguez/mlog"
	"github.com/ledongthuc/admin-generator-go/entity"
)

// settings is defination of service settings. Use to load and manage setting
type adminUsers struct{}

// Settings instance of Settings services
var AdminUsers adminUsers

// Load loads settings from /conf/settings.yml
func (adminUsers *adminUsers) Load() ([]entity.AdminUser, error) {
	data, _ := ioutil.ReadFile("conf/admin_users.yml")

	var users []entity.AdminUser
	err := yaml.Unmarshal([]byte(data), &users)
	return users, err
}

func (adminUsers *adminUsers) Login(checkingUser entity.AdminUser) bool {
	users, err := adminUsers.Load()
	if err != nil {
		mlog.Error(err)
		return false
	}

	for _, user := range users {
		if user.Equal(&checkingUser) {
			return true
		}
	}

	return false
}
