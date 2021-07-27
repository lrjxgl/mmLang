package indexModel

import (
	"app/config"
)

type BlacklistModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Dateline      uint   `json:"dateline"` 

}

func (BlacklistModel) TableName() string {
	return config.TablePre + "blacklist"
}

func BlacklistList(list []BlacklistModel) []BlacklistModel {
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