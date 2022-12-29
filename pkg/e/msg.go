package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_EXIST_TAG:                 "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:             "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "图片保存失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "图片检查失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "图片格式或者大小不合法",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "检查文件存在错误",
	ERROR_GET_ARTICLE_FAIL:          "查询文章失败",
	ERROR_GET_ARTICLES_FAIL:         "查询文章列表失败",
	ERROR_CREATE_ARTICLE_FAIL:       "创建文章失败",
	ERROR_UPDATE_ARTICLE_FAIL:       "更新文章失败",
	ERROR_DELETE_ARTICLE_FAIL:       "删除文章失败",
	ERROR_GET_TAG_FAIL:              "查询标签失败",
	ERROR_GET_TAGS_FAIL:             "查询标签列表失败",
	ERROR_CREATE_TAG_FAIL:           "创建标签失败",
	ERROR_UPDATE_TAG_FAIL:           "更新标签失败",
	ERROR_DELETE_TAG_FAIL:           "删除标签失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
