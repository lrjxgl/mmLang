package indexModel

import (
	"app/config"
)

type InviteAccountModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Money      float64   `json:"money"` 
	Income      float64   `json:"income"` 
	Dateline      uint   `json:"dateline"` 

}

func (InviteAccountModel) TableName() string {
	return config.TablePre + "invite_account"
}

func InviteAccountList(list []InviteAccountModel) []InviteAccountModel {
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