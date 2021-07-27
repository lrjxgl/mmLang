package indexModel

import (
	"app/config"
)

type ModelIndexModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Tablename      string   `json:"tablename"` 

}

func (ModelIndexModel) TableName() string {
	return config.TablePre + "model_index"
}

func ModelIndexList(list []ModelIndexModel) []ModelIndexModel {
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