package indexModel

import (
	"app/config"
)

type TableDataModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Tableid      uint   `json:"tableid"` 
	Catid      uint   `json:"catid"` 
	Objectid      uint   `json:"objectid"` 
	Userid      uint   `json:"userid"` 
	Createtime      string   `json:"createtime"` 
	Status      uint   `json:"status"` 
	Comment_num      uint   `json:"comment_num"` 
	View_num      uint   `json:"view_num"` 
	Content      string   `json:"content"` 

}

func (TableDataModel) TableName() string {
	return config.TablePre + "table_data"
}

func TableDataList(list []TableDataModel) []TableDataModel {
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