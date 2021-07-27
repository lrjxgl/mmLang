package forumModel

import (
	"app/config"
)

type ForumFeedsModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Objectid      uint   `json:"objectid"` 
	Fuserid      uint   `json:"fuserid"` 
	Dateline      uint   `json:"dateline"` 

}

func (ForumFeedsModel) TableName() string {
	return config.TablePre + "mod_forum_feeds"
}

func ForumFeedsList(list []ForumFeedsModel) []ForumFeedsModel {
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