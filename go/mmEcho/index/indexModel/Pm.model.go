package indexModel

import (
	"app/config"
)

type PmModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Type_id      uint   `json:"type_id"` 
	T_userid      uint   `json:"t_userid"` 
	Status      uint   `json:"status"` 
	Content      string   `json:"content"` 
	Createtime      string   `json:"createtime"` 

}

func (PmModel) TableName() string {
	return config.TablePre + "pm"
}

func PmList(list []PmModel) []PmModel {
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