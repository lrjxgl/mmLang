package indexIndex

import (
	"app/config"
	"app/index/indexModel"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)
/*@@NavbarIndex@@*/
func NavbarIndex(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	var db = config.Db
	var list = []indexModel.NavbarModel{}
	res := db.Where("status=1").Find(&list)
	if res.Error != nil {
		list = []indexModel.NavbarModel{}
	}
	//输出浏览器
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["list"] = indexModel.NavbarList(list)
	reJson["type"] = reflect.TypeOf(list)
	return c.JSON(http.StatusOK, reJson)
}
/*@@NavbarShow@@*/
func NavbarShow(c echo.Context) (err error) {

	id := c.QueryParam("id")
	var db = config.Db
	data := new(indexModel.NavbarModel)
	res := db.Where("id=?  AND status=1  ", id).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	//输出浏览器
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = data
	return c.JSON(http.StatusOK, reJson)
}
/*@@NavbarAdd@@*/
func NavbarAdd(c echo.Context) (err error) {

	id, err := strconv.Atoi(c.QueryParam("id"))
	var db = config.Db

	var data = indexModel.NavbarModel{}
	if id != 0 {
		res := db.Where("id=?  AND status<4  ", id).First(&data)
		if res.Error != nil {
			return config.Success(c, 1, "数据不存在")
		}
	}

	//输出浏览器
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = data
	reJson["id"] = id
	return c.JSON(http.StatusOK, reJson)
}
/*@@NavbarSave@@*/
func NavbarSave(c echo.Context) (err error) {

	id, err := strconv.Atoi(c.FormValue("id"))
	var db = config.Db
	var data = indexModel.NavbarModel{}
	if id != 0 {
		res := db.Where("id=?  AND status<4  ", id).First(&data)
		if res.Error != nil {
			return config.Success(c, 1, "数据不存在")
		}
	}
	//新增数据

	postData := map[string]interface{}{}
	postData["title"] = c.FormValue("title")
	postData["description"] = c.FormValue("description")
	now := time.Now()
	postData["createtime"] = now.Format("2006-01-02 15:04:05")
	if id != 0 {
		db.Create(postData)
	} else {
		db.Model(indexModel.NavbarModel{}).Where("id=?", id).Updates(postData)
	}

	//输出浏览器
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = postData
	return c.JSON(http.StatusOK, reJson)
}
/*@@NavbarStatus@@*/
func NavbarStatus(c echo.Context) (err error) {
	id := c.QueryParam("id")
	var db = config.Db
	data := new(indexModel.NavbarModel)
	res := db.Where("id=?", id).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	status := 1
	if data.Status == 1 {
		status = 2
	}
	db.Model(indexModel.NavbarModel{}).Where("id=?", id).Update("status", status)
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["status"] = status
	return c.JSON(http.StatusOK, reJson)

}

/*@@NavbarDelete@@*/
func NavbarDelete(c echo.Context) (err error) {
	id := c.QueryParam("id")
	var db = config.Db
	data := new(indexModel.NavbarModel)
	res := db.Where("id=?", id).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	db.Model(indexModel.NavbarModel{}).Where("id=?", id).Update("status", 11)
	return config.Success(c, 0, "删除成功")

}