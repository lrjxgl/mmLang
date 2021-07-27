package indexModel

import (
	"app/config"
)

type ModuleModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Tablename      string   `json:"tablename"` 
	Data      string   `json:"data"` 
	Cat_tpl      string   `json:"cat_tpl"` 
	List_tpl      string   `json:"list_tpl"` 
	Show_tpl      string   `json:"show_tpl"` 
	Module      string   `json:"module"` 

}

func (ModuleModel) TableName() string {
	return config.TablePre + "module"
}

func ModuleList(list []ModuleModel) []ModuleModel {
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