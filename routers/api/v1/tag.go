package v1

import (
	"github.com/GoldenLeeK/go-gin-blog/pkg/app"
	"github.com/GoldenLeeK/go-gin-blog/pkg/export"
	"github.com/GoldenLeeK/go-gin-blog/service"
	"net/http"

	"github.com/GoldenLeeK/go-gin-blog/models"
	"github.com/GoldenLeeK/go-gin-blog/pkg/e"
	"github.com/GoldenLeeK/go-gin-blog/pkg/setting"
	"github.com/GoldenLeeK/go-gin-blog/pkg/utils"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取标签列表
func GetTags(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	tagService := service.Tag{
		PageSize: setting.AppSetting.PageSize,
		PageNum:  utils.GetPage(c),
		Maps:     maps,
	}

	lists, err := tagService.GetAll()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}
	data["lists"] = lists
	data["total"] = tagService.Total()

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

//添加标签
func AddTag(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.PostForm("name")
	state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()
	createdBy := c.PostForm("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不得为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	tagService := service.Tag{
		Tag: &models.Tag{
			Name:      name,
			State:     state,
			CreatedBy: createdBy,
		},
	}

	exists, _ := tagService.ExistByName()

	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}

	_, err := tagService.Add()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CREATE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

//编辑标签
func EditTag(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.PostForm("name")
	modifiedBy := c.PostForm("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	tagService := service.Tag{
		ID:  id,
		Tag: &models.Tag{},
	}
	exists, _ := tagService.ExistByID()
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	tagService.Tag.ModifiedBy = modifiedBy
	if name != "" {
		tagService.Tag.Name = name
	}
	if state != -1 {
		tagService.Tag.State = state
	}
	_, err := tagService.Update()

	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_UPDATE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

//删除标签
func DeleteTag(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	tagService := service.Tag{
		ID: id,
	}
	exists, _ := tagService.ExistByID()
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	_, err := tagService.Delete()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

//导出标签
func ExportTag(c *gin.Context) {
	appG := app.Gin{C: c}

	tagService := service.Tag{
		PageSize: setting.AppSetting.PageSize,
		PageNum:  utils.GetPage(c),
		Maps:     map[string]interface{}{},
	}
	if name := c.PostForm("name"); name != "" {
		tagService.Maps["name"] = name
	}
	if state := c.PostForm("state"); state != "" {
		tagService.Maps["state"] = com.StrTo(state).MustInt()
	}

	fileName, err := tagService.Export()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_EXPORT_TAG_FAILE, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"export_url":       export.GetExcelFullUrl(fileName),
		"export_save_path": export.GetExcelPath() + fileName,
	})

}
