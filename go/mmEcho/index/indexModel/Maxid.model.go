package indexModel

import (
	"app/config"
)

type MaxidModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	T      uint   `json:"t"` 

}

func (MaxidModel) TableName() string {
	return config.TablePre + "maxid"
}

func MaxidList(list []MaxidModel) []MaxidModel {
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