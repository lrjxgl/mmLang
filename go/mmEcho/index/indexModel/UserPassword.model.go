package indexModel

import (
	"app/config"
)

type UserPasswordModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Password      string   `json:"password"` 
	Paypwd      string   `json:"paypwd"` 
	Salt      string   `json:"salt"` 

}

func (UserPasswordModel) TableName() string {
	return config.TablePre + "user_password"
}

func UserPasswordList(list []UserPasswordModel) []UserPasswordModel {
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