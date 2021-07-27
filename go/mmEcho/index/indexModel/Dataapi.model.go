package indexModel

import (
	"app/config"
)

type DataapiModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Word      string   `json:"word"` 
	Dateline      uint   `json:"dateline"` 
	Type_id      uint   `json:"type_id"` 
	Equation      string   `json:"equation"` 
	Info      string   `json:"info"` 
	Content      string   `json:"content"` 
	Status      uint   `json:"status"` 

}

func (DataapiModel) TableName() string {
	return config.TablePre + "dataapi"
}

func DataapiList(list []DataapiModel) []DataapiModel {
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