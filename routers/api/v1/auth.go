package v1

import (
	"github.com/GoldenLeeK/go-gin-blog/pkg/app"
	"github.com/GoldenLeeK/go-gin-blog/service/auth_service"
	"net/http"

	"github.com/GoldenLeeK/go-gin-blog/pkg/e"
	"github.com/GoldenLeeK/go-gin-blog/pkg/utils"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	username := c.PostForm("username")
	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{
		Username: username,
		Password: password,
	}
	isExist := authService.CheckAuth()
	if !isExist {
		appG.Response(http.StatusOK, e.ERROR_AUTH, nil)
		return
	}

	token, err := utils.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	appG.Response(http.StatusOK, e.SUCCESS, data)

}
