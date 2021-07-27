package indexModel

import (
	"app/config"
)

type UserLastaddrModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Nickname      string   `json:"nickname"` 
	Telephone      string   `json:"telephone"` 
	Address      string   `json:"address"` 

}

func (UserLastaddrModel) TableName() string {
	return config.TablePre + "user_lastaddr"
}

func UserLastaddrList(list []UserLastaddrModel) []UserLastaddrModel {
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