package forumModel

import (
	"app/config"
)

type ForumCommentModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Status      uint   `json:"status"` 
	Gid      uint   `json:"gid"` 
	Objectid      uint   `json:"objectid"` 
	Pid      uint   `json:"pid"` 
	Createtime      string   `json:"createtime"` 
	Content      string   `json:"content"` 
	Ip      string   `json:"ip"` 
	Ip_city      string   `json:"ip_city"` 

}

func (ForumCommentModel) TableName() string {
	return config.TablePre + "mod_forum_comment"
}

func ForumCommentList(list []ForumCommentModel) []ForumCommentModel {
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