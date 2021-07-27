package indexModel

import (
	"app/config"
)

type FriendModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Touserid      uint   `json:"touserid"` 
	Dateline      uint   `json:"dateline"` 

}

func (FriendModel) TableName() string {
	return config.TablePre + "friend"
}

func FriendList(list []FriendModel) []FriendModel {
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