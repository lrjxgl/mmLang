package indexAdmin

import (
	"app/access"
	"app/config"
	"app/index/indexModel"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LoginIndex(c echo.Context) (err error) {
	return config.Success(c, 0, "success")
}
func LoginSave(c echo.Context) (err error) {
	var db = config.Db
	username := c.FormValue("username")
	password := c.FormValue("password")
	r := indexModel.AdminModel{}
	res := db.First(&r)

	if res.Error != nil {
		//生成默认管理员
		salt := rand.Intn(10000)
		saltStr := strconv.Itoa(salt)
		password = config.Umd5("admin" + saltStr)
		inData := indexModel.AdminModel{
			Username:  "admin",
			Salt:      uint(salt),
			Password:  password,
			Isfounder: 1,
		}
		db.Create(inData)
		return config.Success(c, 1, "登录失败")
	}
	row := indexModel.AdminModel{}
	res2 := db.Where("username=?", username).First(&row)
	if res2.Error != nil {
		return config.Success(c, 1, "登录失败")
	}
	password = config.Umd5("admin" + strconv.Itoa(int(row.Salt)))
	if password != row.Password {
		return config.Success(c, 1, "登录失败")
	}
	reJson := access.GetAdminToken(row.Id, password)
	reJson["error"] = 0
	reJson["message"] = "success"

	return c.JSON(http.StatusOK, reJson)
}
