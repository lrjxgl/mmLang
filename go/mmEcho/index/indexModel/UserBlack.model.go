package indexModel

import (
	"app/config"
)

type UserBlackModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Buserid      uint   `json:"buserid"` 

}

func (UserBlackModel) TableName() string {
	return config.TablePre + "user_black"
}

func UserBlackList(list []UserBlackModel) []UserBlackModel {
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