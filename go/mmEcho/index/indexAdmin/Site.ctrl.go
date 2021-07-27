package indexAdmin

import (
	"app/config"
	"app/access"
	"app/index/indexModel"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)
/*解决import未使用*/
func SiteNull(c echo.Context) (err error){
	 
	now := time.Now()
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["now"]=now;
	 
	return c.JSON(http.StatusOK, reJson)
}

/*@@SiteIndex@@*/
func SiteIndex(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	var db = config.Db
	var list = []indexModel.SiteModel{}
	 
	where:=" status in(0,1,2) ";
	//统计数量
	start, err := strconv.Atoi(c.QueryParam("per_page"))
	if err!=nil {
		start=0;
	}
	limit, err2 := strconv.Atoi(c.QueryParam("limit"))
	if err2!=nil || limit==0 {
		limit=24;
	}
	var rscount int64;
	db.Model(&indexModel.SiteModel{}).Where(where).Count(&rscount);
	//获取列表
	res := db.Where(where).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []indexModel.SiteModel{}
	}
	//输出浏览器
	var per_page int64=int64(start+limit);
	if per_page>rscount {
		per_page=0;
	}
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["list"] = indexModel.SiteList(list)
	reJson["type"] = reflect.TypeOf(list)
	reJson["rscount"]=rscount;
	reJson["per_page"]=per_page;
	return c.JSON(http.StatusOK, reJson)
}

/*@@SiteAdd@@*/
func SiteAdd(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	siteid, err := strconv.Atoi(c.QueryParam("siteid"))
	var db = config.Db

	var data = indexModel.SiteModel{}
	if siteid != 0 {
		res := db.Where("siteid=?  AND status<4  ", siteid).First(&data)
		if res.Error != nil {
			return config.Success(c, 1, "数据不存在")
		}
		
	}

	//输出浏览器
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = data
	reJson["siteid"] = siteid
	return c.JSON(http.StatusOK, reJson)
}

/*@@SiteSave@@*/
func SiteSave(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	siteid, err := strconv.Atoi(c.FormValue("siteid"))
	var db = config.Db
	var data = indexModel.SiteModel{}
	if siteid != 0 {
		res := db.Where("siteid=?  AND status<4  ", siteid).First(&data)
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
	if siteid != 0 {
		db.Create(postData)
	} else {
		db.Model(indexModel.SiteModel{}).Where("siteid=?", siteid).Updates(postData)
	}

	//输出浏览器
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = postData
	return c.JSON(http.StatusOK, reJson)
}
/*@@SiteStatus@@*/
func SiteStatus(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	siteid := c.QueryParam("siteid")
	var db = config.Db
	data := new(indexModel.SiteModel)
	res := db.Where("siteid=?", siteid).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	
	status := 1
	if data.Status == 1 {
		status = 2
	}
	db.Model(indexModel.SiteModel{}).Where("siteid=?", siteid).Update("status", status)
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["status"] = status
	return c.JSON(http.StatusOK, reJson)

}

/*@@SiteDelete@@*/
func SiteDelete(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	siteid := c.QueryParam("siteid")
	var db = config.Db
	data := new(indexModel.SiteModel)
	res := db.Where("siteid=?", siteid).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	
	db.Model(indexModel.SiteModel{}).Where("siteid=?", siteid).Update("status", 11)
	return config.Success(c, 0, "删除成功")

}