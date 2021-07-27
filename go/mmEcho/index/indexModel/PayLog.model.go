package indexModel

import (
	"app/config"
)

type PayLogModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Dateline      uint   `json:"dateline"` 
	Userid      uint   `json:"userid"` 
	Type_id      uint   `json:"type_id"` 
	Content      string   `json:"content"` 
	Isdelete      uint   `json:"isdelete"` 
	Money      float64   `json:"money"` 
	Ispay      uint   `json:"ispay"` 

}

func (PayLogModel) TableName() string {
	return config.TablePre + "pay_log"
}

func PayLogList(list []PayLogModel) []PayLogModel {
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