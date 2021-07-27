package indexModel

import (
	"app/config"
)

type PmIndexModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	T_userid      uint   `json:"t_userid"` 
	Pm_num      uint   `json:"pm_num"` 
	Type_id      uint   `json:"type_id"` 
	Status      uint   `json:"status"` 
	Content      string   `json:"content"` 
	Createtime      string   `json:"createtime"` 

}

func (PmIndexModel) TableName() string {
	return config.TablePre + "pm_index"
}

func PmIndexList(list []PmIndexModel) []PmIndexModel {
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