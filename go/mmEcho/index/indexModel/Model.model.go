package indexModel

import (
	"app/config"
)

type ModelModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Tablename      string   `json:"tablename"` 
	Cat_tpl      string   `json:"cat_tpl"` 
	List_tpl      string   `json:"list_tpl"` 
	Show_tpl      string   `json:"show_tpl"` 
	Module      string   `json:"module"` 
	Status      uint   `json:"status"` 
	Data      string   `json:"data"` 

}

func (ModelModel) TableName() string {
	return config.TablePre + "model"
}

func ModelList(list []ModelModel) []ModelModel {
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