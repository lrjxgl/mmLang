package indexModel

import (
	"app/config"
)

type AttachModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Dateline      uint   `json:"dateline"` 
	Status      uint   `json:"status"` 
	Userid      uint   `json:"userid"` 
	File_name      string   `json:"file_name"` 
	File_url      string   `json:"file_url"` 
	File_size      float64   `json:"file_size"` 
	File_type      string   `json:"file_type"` 
	File_group      string   `json:"file_group"` 

}

func (AttachModel) TableName() string {
	return config.TablePre + "attach"
}

func AttachList(list []AttachModel) []AttachModel {
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