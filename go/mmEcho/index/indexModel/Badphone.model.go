package indexModel

import (
	"app/config"
)

type BadphoneModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Telephone      string   `json:"telephone"` 
	Createtime      string   `json:"createtime"` 
	Content      string   `json:"content"` 

}

func (BadphoneModel) TableName() string {
	return config.TablePre + "badphone"
}

func BadphoneList(list []BadphoneModel) []BadphoneModel {
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