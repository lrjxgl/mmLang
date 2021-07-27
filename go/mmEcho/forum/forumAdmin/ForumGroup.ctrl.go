package forumAdmin

import (
	"app/config"
	"app/access"
	"app/forum/forumModel"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)
/*解决import未使用*/
func ForumGroupNull(c echo.Context) (err error){
	 
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

/*@@ForumGroupIndex@@*/
func ForumGroupIndex(c echo.Context) (err error) {
	fmt.Print("forumIndex")
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	var db = config.Db
	var list = []forumModel.ForumGroupModel{}
	 
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
	db.Model(&forumModel.ForumGroupModel{}).Where(where).Count(&rscount);
	//获取列表
	res := db.Where(where).Limit(limit).Offset(start).Find(&list)
	if res.Error != nil {
		list = []forumModel.ForumGroupModel{}
	}
	//输出浏览器
	var per_page int64=int64(start+limit);
	if per_page>rscount {
		per_page=0;
	}
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["list"] = forumModel.ForumGroupList(list)
	reJson["type"] = reflect.TypeOf(list)
	reJson["rscount"]=rscount;
	reJson["per_page"]=per_page;
	return c.JSON(http.StatusOK, reJson)
}

/*@@ForumGroupAdd@@*/
func ForumGroupAdd(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	gid, err := strconv.Atoi(c.QueryParam("gid"))
	var db = config.Db

	var data = forumModel.ForumGroupModel{}
	if gid != 0 {
		res := db.Where("gid=?  AND status<4  ", gid).First(&data)
		if res.Error != nil {
			return config.Success(c, 1, "数据不存在")
		}
		
	}

	//输出浏览器
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = data
	reJson["gid"] = gid
	return c.JSON(http.StatusOK, reJson)
}

/*@@ForumGroupSave@@*/
func ForumGroupSave(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	gid, err := strconv.Atoi(c.FormValue("gid"))
	var db = config.Db
	var data = forumModel.ForumGroupModel{}
	if gid != 0 {
		res := db.Where("gid=?  AND status<4  ", gid).First(&data)
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
	if gid != 0 {
		db.Create(postData)
	} else {
		db.Model(forumModel.ForumGroupModel{}).Where("gid=?", gid).Updates(postData)
	}

	//输出浏览器
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = postData
	return c.JSON(http.StatusOK, reJson)
}
/*@@ForumGroupStatus@@*/
func ForumGroupStatus(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	gid := c.QueryParam("gid")
	var db = config.Db
	data := new(forumModel.ForumGroupModel)
	res := db.Where("gid=?", gid).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	
	status := 1
	if data.Status == 1 {
		status = 2
	}
	db.Model(forumModel.ForumGroupModel{}).Where("gid=?", gid).Update("status", status)
	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["status"] = status
	return c.JSON(http.StatusOK, reJson)

}

/*@@ForumGroupRecommend@@*/
func ForumGroupRecommend(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	gid := c.QueryParam("gid")
	var db = config.Db
	data := new(forumModel.ForumGroupModel)
	res := db.Where("gid=?", gid).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	
	isrecommed := 1
	if data.Isrecommend == 1 {
		isrecommed = 0
	}
	db.Model(forumModel.ForumGroupModel{}).Where("gid=?", gid).Update("isrecommend", isrecommed)

	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["isrecommend"] = isrecommed
	return c.JSON(http.StatusOK, reJson)

}
/*@@ForumGroupDelete@@*/
func ForumGroupDelete(c echo.Context) (err error) {
	adminid := access.AdminCheckAccess(c)
	if adminid == 0 {
		return config.Success(c, 1000, "暂无权限")
	}

	gid := c.QueryParam("gid")
	var db = config.Db
	data := new(forumModel.ForumGroupModel)
	res := db.Where("gid=?", gid).First(&data)
	if res.Error != nil {
		return config.Success(c, 1, "数据不存在")
	}
	
	db.Model(forumModel.ForumGroupModel{}).Where("gid=?", gid).Update("status", 11)
	return config.Success(c, 0, "删除成功")

}