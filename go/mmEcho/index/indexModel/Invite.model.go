package indexModel

import (
	"app/config"
)

type InviteModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	In_userid      uint   `json:"in_userid"` 
	Dateline      uint   `json:"dateline"` 
	Ispay      uint   `json:"ispay"` 
	Bstatus      uint   `json:"bstatus"` 

}

func (InviteModel) TableName() string {
	return config.TablePre + "invite"
}

func InviteList(list []InviteModel) []InviteModel {
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