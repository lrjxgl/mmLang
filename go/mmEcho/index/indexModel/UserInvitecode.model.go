package indexModel

import (
	"app/config"
)

type UserInvitecodeModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Icode      string   `json:"icode"` 

}

func (UserInvitecodeModel) TableName() string {
	return config.TablePre + "user_invitecode"
}

func UserInvitecodeList(list []UserInvitecodeModel) []UserInvitecodeModel {
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