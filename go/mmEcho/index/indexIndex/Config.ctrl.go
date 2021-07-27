package indexIndex

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
func ConfigNull(c echo.Context) (err error){
	 
	now := time.Now()
	
			userid := access.UserCheckAccess(c)
			if userid == 0 {
				return config.Success(c, 1000, "请先登录")
			}
		
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["now"]=now;
	 
	return c.JSON(http.StatusOK, reJson)
}

/*@@ConfigIndex@@*/
func ConfigIndex(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	
	var db = config.Db
	var list = []indexModel.ConfigModel{}
	 
	where:=" 1 ";
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
	db.Model(&indexModel.ConfigModel{}).Where(where).Count(&rscount);
	//获取列表
	res := db.Where(where).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []indexModel.ConfigModel{}
	}
	//输出浏览器
	var per_page int64=int64(start+limit);
	if per_page>rscount {
		per_page=0;
	}
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["list"] = indexModel.ConfigList(list)
	reJson["type"] = reflect.TypeOf(list)
	reJson["rscount"]=rscount;
	reJson["per_page"]=per_page;
	return c.JSON(http.StatusOK, reJson)
}

/*@@ConfigList@@*/
func ConfigList(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	
	var db = config.Db
	var list = []indexModel.ConfigModel{}
	 
	where:=" 1 ";
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
	db.Model(&indexModel.ConfigModel{}).Where(where).Count(&rscount);
	//获取列表
	res := db.Where(where).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []indexModel.ConfigModel{}
	}
	//输出浏览器
	var per_page int64=int64(start+limit);
	if per_page>rscount {
		per_page=0;
	}
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["list"] = indexModel.ConfigList(list)
	reJson["type"] = reflect.TypeOf(list)
	reJson["rscount"]=rscount;
	reJson["per_page"]=per_page;
	return c.JSON(http.StatusOK, reJson)
}

/*@@ConfigShow@@*/
func ConfigShow(c echo.Context) (err error) {
	
	id := c.QueryParam("id")
	var db = config.Db
	data := new(indexModel.ConfigModel)
	res := db.Where("id=?  ", id).First(&data)
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