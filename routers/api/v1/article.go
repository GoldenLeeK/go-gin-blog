package v1

import (
	"github.com/GoldenLeeK/go-gin-blog/pkg/app"
	"github.com/GoldenLeeK/go-gin-blog/pkg/logging"
	"github.com/GoldenLeeK/go-gin-blog/service/article_service"
	"net/http"

	"github.com/GoldenLeeK/go-gin-blog/models"
	"github.com/GoldenLeeK/go-gin-blog/pkg/e"
	"github.com/GoldenLeeK/go-gin-blog/pkg/setting"
	"github.com/GoldenLeeK/go-gin-blog/pkg/utils"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)

}

func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许为0或者1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{
		PageNum:  utils.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
		Maps:     maps,
	}
	lists, err := articleService.GetAll()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	data["lists"] = lists
	data["total"] = articleService.Total()

	appG.Response(http.StatusOK, e.SUCCESS, data)

}
func AddArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	coverImageUrl := c.PostForm("cover_image_url")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	createdBy := c.PostForm("created_by")
	state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(coverImageUrl, "cover_image_url").Message("封面图片不得为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	if !models.ExistTagById(tagId) {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
	}

	articleService := article_service.Article{
		Article: &models.Article{
			TagID:         tagId,
			Title:         title,
			CoverImageUrl: coverImageUrl,
			Desc:          desc,
			Content:       content,
			CreatedBy:     createdBy,
			State:         state,
		},
	}

	_, err := articleService.Add()

	if err != nil {
		logging.Info(err)
		appG.Response(http.StatusOK, e.ERROR_CREATE_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
func EditArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	coverImageUrl := c.PostForm("cover_image_url")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	modifiedBy := c.PostForm("modified_by")

	var state int = -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.Required(coverImageUrl, "cover_image_url").Message("封面图片不得为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{
		ID: id,
	}

	exists, _ := articleService.ExistByID()
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	if !models.ExistTagById(tagId) {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	articleService.Article = article
	if tagId > 0 {
		articleService.Article.TagID = tagId
	}
	if title != "" {
		articleService.Article.Title = title
	}
	if desc != "" {
		articleService.Article.Desc = desc
	}
	if content != "" {
		articleService.Article.Content = content
	}
	if coverImageUrl != "" {
		articleService.Article.CoverImageUrl = coverImageUrl
	}
	articleService.Article.ModifiedBy = modifiedBy

	_, err = articleService.Update()

	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_UPDATE_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}
func DeleteArticle(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{
		ID: id,
	}
	exists, _ := articleService.ExistByID()
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	_, err := articleService.Delete()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
