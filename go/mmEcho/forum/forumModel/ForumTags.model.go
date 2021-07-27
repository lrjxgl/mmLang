package forumModel

import (
	"app/config"
)

type ForumTagsModel struct {
	Tagid      uint   `gorm:"primaryKey";json:"tagid"`
	Title      string   `json:"title"` 
	Status      uint   `json:"status"` 
	Total_num      uint   `json:"total_num"` 
	View_num      uint   `json:"view_num"` 
	Gkey      string   `json:"gkey"` 
	Gnum      uint   `json:"gnum"` 

}

func (ForumTagsModel) TableName() string {
	return config.TablePre + "mod_forum_tags"
}

func ForumTagsList(list []ForumTagsModel) []ForumTagsModel {
	slen := len(list)
	if slen == 0 {
		return list
	}

	for i := 0; i < slen; i++ {
		m := list[i]
		
		list[i] = m
	}
	return list
}