package indexModel

import (
	"app/config"
)

type TableFieldsModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Tableid      uint   `json:"tableid"` 
	Title      string   `json:"title"` 
	Fieldname      string   `json:"fieldname"` 
	Fieldtype      string   `json:"fieldtype"` 
	Description      string   `json:"description"` 
	Status      uint   `json:"status"` 
	Orderindex      uint   `json:"orderindex"` 
	Islist      uint   `json:"islist"` 
	Optionlist      string   `json:"optionlist"` 

}

func (TableFieldsModel) TableName() string {
	return config.TablePre + "table_fields"
}

func TableFieldsList(list []TableFieldsModel) []TableFieldsModel {
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