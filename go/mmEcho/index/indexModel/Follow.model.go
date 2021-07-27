package indexModel

import (
	"app/config"
)

type FollowModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	T_userid      uint   `json:"t_userid"` 
	Status      uint   `json:"status"` 
	Dateline      uint   `json:"dateline"` 

}

func (FollowModel) TableName() string {
	return config.TablePre + "follow"
}

func FollowList(list []FollowModel) []FollowModel {
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