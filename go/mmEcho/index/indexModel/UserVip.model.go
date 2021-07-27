package indexModel

import (
	"app/config"
)

type UserVipModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Starttime      uint   `json:"starttime"` 
	Firsttime      uint   `json:"firsttime"` 
	Endtime      uint   `json:"endtime"` 
	Stype      uint   `json:"stype"` 
	Status      uint   `json:"status"` 

}

func (UserVipModel) TableName() string {
	return config.TablePre + "user_vip"
}

func UserVipList(list []UserVipModel) []UserVipModel {
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