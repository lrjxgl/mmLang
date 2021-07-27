package indexModel

import (
	"app/config"
)

type UserGroupPeopleModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Gid      uint   `json:"gid"` 
	Userid      uint   `json:"userid"` 
	Dateline      uint   `json:"dateline"` 

}

func (UserGroupPeopleModel) TableName() string {
	return config.TablePre + "user_group_people"
}

func UserGroupPeopleList(list []UserGroupPeopleModel) []UserGroupPeopleModel {
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