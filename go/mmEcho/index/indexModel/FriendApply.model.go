package indexModel

import (
	"app/config"
)

type FriendApplyModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Touserid      uint   `json:"touserid"` 
	Dateline      uint   `json:"dateline"` 
	Description      string   `json:"description"` 

}

func (FriendApplyModel) TableName() string {
	return config.TablePre + "friend_apply"
}

func FriendApplyList(list []FriendApplyModel) []FriendApplyModel {
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