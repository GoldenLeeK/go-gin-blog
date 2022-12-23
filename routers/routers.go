package routers

import (
	"github.com/GoldenLeeK/go-gin-blog/middleware/jwt"
	"github.com/GoldenLeeK/go-gin-blog/pkg/setting"
	"github.com/GoldenLeeK/go-gin-blog/pkg/upload"
	"github.com/GoldenLeeK/go-gin-blog/routers/api"
	v1 "github.com/GoldenLeeK/go-gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	//获取授权token
	r.POST("/auth", v1.GetAuth)
	//上传图片
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取帖子列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定的帖子
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建帖子
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定帖子
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定帖子
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	}
	return r
}
