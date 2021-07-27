package forumModel

import (
	"app/config"
)

type ForumDataModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Content      string   `json:"content"` 
	Dateline      uint   `json:"dateline"` 

}

func (ForumDataModel) TableName() string {
	return config.TablePre + "mod_forum_data"
}

func ForumDataList(list []ForumDataModel) []ForumDataModel {
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