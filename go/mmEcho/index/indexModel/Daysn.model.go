package indexModel

import (
	"app/config"
)

type DaysnModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Bday      string   `json:"bday"` 
	Tablename      string   `json:"tablename"` 
	Shopid      uint   `json:"shopid"` 
	DaySn      uint   `json:"daySn"` 

}

func (DaysnModel) TableName() string {
	return config.TablePre + "daysn"
}

func DaysnList(list []DaysnModel) []DaysnModel {
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