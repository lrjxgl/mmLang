package forumModel

import (
	"app/config"
)

type ForumImgModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Objectid      uint   `json:"objectid"` 
	Shopid      uint   `json:"shopid"` 
	Orderindex      uint   `json:"orderindex"` 
	Createtime      string   `json:"createtime"` 
	Imgurl      string   `json:"imgurl"` 

}

func (ForumImgModel) TableName() string {
	return config.TablePre + "mod_forum_img"
}

func ForumImgList(list []ForumImgModel) []ForumImgModel {
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