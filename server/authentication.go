package server

import "github.com/ledongthuc/admin-generator-go/services"

type authentication struct{}

var Authentication authentication

func (auth *authentication) Authenticate(email, password string) bool {
	return services.AdminUsers.Login(email, password)
}
