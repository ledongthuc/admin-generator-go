package entity

// AdminUser uses to mapping setting from /conf/setting.yml
type AdminUser struct {
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}

// Equal compares 2 user
func (user *AdminUser) Equal(anotherUser *AdminUser) bool {
	if user.Email == anotherUser.Email &&
		user.Password == anotherUser.Password {
		return true
	}

	return false
}
