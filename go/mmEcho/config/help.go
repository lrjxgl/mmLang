package config

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Success(c echo.Context, errno uint, message string) (err error) {
	reJson := make(map[string]interface{})
	reJson["error"] = errno
	reJson["message"] = message
	return c.JSON(http.StatusOK, reJson)
}

func Image_site(str string) string {

	if str != "" {
		var s string = AppConfig["images_site"].(string)
		return s + "/" + str
	} else {
		return str
	}
}

func Umd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	data = []byte(md5str)
	has = md5.Sum(data)
	md5str = fmt.Sprintf("%x", has)
	return md5str
}
