package forumModel

import (
	"app/config"
)

type ForumTagsIndexModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Tagid      uint   `json:"tagid"` 
	Objectid      uint   `json:"objectid"` 
	Orderindex      uint   `json:"orderindex"` 

}

func (ForumTagsIndexModel) TableName() string {
	return config.TablePre + "mod_forum_tags_index"
}

func ForumTagsIndexList(list []ForumTagsIndexModel) []ForumTagsIndexModel {
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