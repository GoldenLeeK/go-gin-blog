package auth_service

import "github.com/GoldenLeeK/go-gin-blog/models"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) CheckAuth() bool {
	return models.CheckAuth(a.Username, a.Password)
}
