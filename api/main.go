package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

const ParamError = "ParamError"
const Success = "success"
const Error = "error"

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Total int         `json:"total"`
	Error string      `json:"error,omitempty"`
}

// 成功信息
func SuccessResponse(msg string, total int, data ...interface{}) Response {
	var newData interface{}
	if len(data) > 0 {
		newData = data[0]
	} else {
		newData = ""
	}
	return Response{
		Code:  10000,
		Data:  newData,
		Msg:   msg,
		Total: total,
		Error: "",
	}
}

// 失败信息
func ErrorResponse(msg string) Response {
	return Response{
		Code:  10001,
		Data:  nil,
		Msg:   msg,
		Error: "",
	}
}

// 上传图片
func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileName := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	path := "static/upload/" + fileName
	if err := c.SaveUploadedFile(file, path); err != nil {
		panic(err)
	}
	ret := make(map[string]string)
	ret["url"] = GetSite(c) + "/" + path
	c.JSON(200, SuccessResponse("上传成功", 0, ret))
}


// 获取Scheme  http 或https
func GetScheme(c *gin.Context) string {
	if scheme := c.Request.Header.Get("X-Forwarded-Proto"); scheme != "" {
		return scheme
	}
	if c.Request.URL.Scheme != "" {
		return c.Request.URL.Scheme
	}
	if c.Request.TLS == nil {
		return "http"
	}
	return "https"
}

// 获取site http://www.aaa.com
func GetSite(c *gin.Context) string {
	return GetScheme(c) + "://" + c.Request.Host
}

// 判断字符串是否为数字
func StringIsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}