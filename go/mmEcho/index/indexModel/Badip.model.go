package indexModel

import (
	"app/config"
)

type BadipModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Ip      string   `json:"ip"` 
	Createtime      string   `json:"createtime"` 
	Content      string   `json:"content"` 

}

func (BadipModel) TableName() string {
	return config.TablePre + "badip"
}

func BadipList(list []BadipModel) []BadipModel {
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