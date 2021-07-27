package indexModel

import (
	"app/config"
)

type InviteAccountLogModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Cuserid      uint   `json:"cuserid"` 
	Money      float64   `json:"money"` 
	Dateline      uint   `json:"dateline"` 
	Content      string   `json:"content"` 

}

func (InviteAccountLogModel) TableName() string {
	return config.TablePre + "invite_account_log"
}

func InviteAccountLogList(list []InviteAccountLogModel) []InviteAccountLogModel {
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