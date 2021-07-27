package indexModel

import (
	"app/config"
)

type DbsessionModel struct {
	Id      string   `gorm:"primaryKey";json:"id"`
	Data      string   `json:"data"` 
	Dateline      uint   `json:"dateline"` 

}

func (DbsessionModel) TableName() string {
	return config.TablePre + "dbsession"
}

func DbsessionList(list []DbsessionModel) []DbsessionModel {
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