package indexAdmin

import (
	"app/access"
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
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}
	var db = config.Db
	var list = []indexModel.NavbarModel{}

	where := " status in(0,1,2) "
	//统计数量
	start, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		start = 0
	}
	limit, err2 := strconv.Atoi(c.QueryParam("limit"))
	if err2 != nil || limit == 0 {
		limit = 24
	}
	var rscount int64
	db.Model(&indexModel.NavbarModel{}).Where(where).Count(&rscount)
	//获取列表
	res := db.Where(where).Find(&list)
	if res.Error != nil {
		list = []indexModel.NavbarModel{}
	} else {
		list = indexModel.NavbarList(list)
	}
	//输出浏览器
	var per_page int64 = int64(start + limit)
	if per_page > rscount {
		per_page = 0
	}
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["list"] = indexModel.NavbarChild(list)
	reJson["type"] = reflect.TypeOf(list)
	reJson["rscount"] = rscount
	reJson["per_page"] = per_page
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
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}
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
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}
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
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}
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
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}
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

/*@@NavbarIndex@@*/
func NavbarGet(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}
	group_id := c.QueryParam("group_id")
	var db = config.Db
	var list = []indexModel.NavbarModel{}
	res := db.Where("group_id=?", group_id).Find(&list)
	if res.Error != nil {
		list = []indexModel.NavbarModel{}
	}
	list = indexModel.NavbarList(list)
	//组装成 child
	pList := indexModel.NavbarChild(list)

	//输出浏览器
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["list"] = pList
	return c.JSON(http.StatusOK, reJson)
}
