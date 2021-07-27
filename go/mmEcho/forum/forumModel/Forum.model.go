package forumModel

import (
	"app/config"
)

type ForumModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Userid      uint   `json:"userid"` 
	Gid      uint   `json:"gid"` 
	Catid      uint   `json:"catid"` 
	Love_num      uint   `json:"love_num"` 
	Fav_num      uint   `json:"fav_num"` 
	Forward_num      uint   `json:"forward_num"` 
	Keywords      string   `json:"keywords"` 
	Description      string   `json:"description"` 
	Dateline      uint   `json:"dateline"` 
	Status      uint   `json:"status"` 
	Comment_num      uint   `json:"comment_num"` 
	Imgurl      string   `json:"imgurl"` 
	Last_time      uint   `json:"last_time"` 
	Grade      uint   `json:"grade"` 
	Isrecommend      uint   `json:"isrecommend"` 
	View_num      uint   `json:"view_num"` 
	Isnew      uint   `json:"isnew"` 
	Tags      string   `json:"tags"` 
	Videourl      string   `json:"videourl"` 
	Money      float64   `json:"money"` 
	Gold      uint   `json:"gold"` 
	Imgsdata      string   `json:"imgsdata"` 
	Createtime      string   `json:"createtime"` 

}

func (ForumModel) TableName() string {
	return config.TablePre + "mod_forum"
}

func ForumList(list []ForumModel) []ForumModel {
	slen := len(list)
	if slen == 0 {
		return list
	}

	for i := 0; i < slen; i++ {
		m := list[i]
		m.Imgurl = config.Image_site(m.Imgurl)
		list[i] = m
	}
	return list
}