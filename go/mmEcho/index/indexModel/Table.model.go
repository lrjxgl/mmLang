package indexModel

import (
	"app/config"
)

type TableModel struct {
	Tableid      uint   `gorm:"primaryKey";json:"tableid"`
	Title      string   `json:"title"` 
	Tablename      string   `json:"tablename"` 
	Description      string   `json:"description"` 
	Status      uint   `json:"status"` 
	IsLogin      uint   `json:"isLogin"` 
	IsList      uint   `json:"isList"` 
	ShowTpl      string   `json:"showTpl"` 
	ListTpl      string   `json:"listTpl"` 
	AddTpl      string   `json:"addTpl"` 

}

func (TableModel) TableName() string {
	return config.TablePre + "table"
}

func TableList(list []TableModel) []TableModel {
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