package services

import (
	"crypto/md5"
	"fmt"
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

func (adminUsers *adminUsers) Login(username, password string) bool {
	users, err := adminUsers.Load()
	if err != nil {
		mlog.Error(err)
		return false
	}

	checkingUser := adminUsers.ComposeCheckingUser(username, password)
	for _, user := range users {
		if user.Equal(&checkingUser) {
			return true
		}
	}

	return false
}

func (adminUsers *adminUsers) ComposeCheckingUser(username, password string) entity.AdminUser {
	encryptedPassword := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	checkingUser := entity.AdminUser{
		Email:    username,
		Password: string(encryptedPassword),
	}

	return checkingUser
}
