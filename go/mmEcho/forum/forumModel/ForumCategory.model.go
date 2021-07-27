package forumModel

import (
	"app/config"
)

type ForumCategoryModel struct {
	Catid      uint   `gorm:"primaryKey";json:"catid"`
	Gid      uint   `json:"gid"` 
	Title      string   `json:"title"` 
	Description      string   `json:"description"` 
	Orderindex      uint   `json:"orderindex"` 
	Status      uint   `json:"status"` 
	Imgurl      string   `json:"imgurl"` 

}

func (ForumCategoryModel) TableName() string {
	return config.TablePre + "mod_forum_category"
}

func ForumCategoryList(list []ForumCategoryModel) []ForumCategoryModel {
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