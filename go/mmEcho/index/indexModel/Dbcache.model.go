package indexModel

import (
	"app/config"
)

type DbcacheModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	K      string   `json:"k"` 
	V      string   `json:"v"` 
	Expire      uint   `json:"expire"` 

}

func (DbcacheModel) TableName() string {
	return config.TablePre + "dbcache"
}

func DbcacheList(list []DbcacheModel) []DbcacheModel {
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