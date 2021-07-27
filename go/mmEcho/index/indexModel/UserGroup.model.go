package indexModel

import (
	"app/config"
)

type UserGroupModel struct {
	Groupid      uint   `gorm:"primaryKey";json:"groupid"`
	Groupname      string   `json:"groupname"` 
	Access      string   `json:"access"` 

}

func (UserGroupModel) TableName() string {
	return config.TablePre + "user_group"
}

func UserGroupList(list []UserGroupModel) []UserGroupModel {
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