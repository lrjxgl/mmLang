package forumModel

import (
	"app/config"
)

type ForumGroupAdminModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Gid      uint   `json:"gid"` 
	Dateline      uint   `json:"dateline"` 
	Status      uint   `json:"status"` 

}

func (ForumGroupAdminModel) TableName() string {
	return config.TablePre + "mod_forum_group_admin"
}

func ForumGroupAdminList(list []ForumGroupAdminModel) []ForumGroupAdminModel {
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