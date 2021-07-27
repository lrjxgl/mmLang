package indexModel

import (
	"app/config"
)

type GoldLogModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Dateline      uint   `json:"dateline"` 
	Userid      uint   `json:"userid"` 
	Type_id      uint   `json:"type_id"` 
	Ispay      uint   `json:"ispay"` 
	Content      string   `json:"content"` 
	Isdelete      uint   `json:"isdelete"` 
	Money      float64   `json:"money"` 

}

func (GoldLogModel) TableName() string {
	return config.TablePre + "gold_log"
}

func GoldLogList(list []GoldLogModel) []GoldLogModel {
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