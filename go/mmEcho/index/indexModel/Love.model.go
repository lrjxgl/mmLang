package indexModel

import (
	"app/config"
)

type LoveModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Tablename      string   `json:"tablename"` 
	Objectid      uint   `json:"objectid"` 
	Userid      uint   `json:"userid"` 

}

func (LoveModel) TableName() string {
	return config.TablePre + "love"
}

func LoveList(list []LoveModel) []LoveModel {
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