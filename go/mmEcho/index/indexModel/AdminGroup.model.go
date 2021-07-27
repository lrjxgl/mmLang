package indexModel

import (
	"app/config"
)

type AdminGroupModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Orderindex      uint   `json:"orderindex"` 
	Content      string   `json:"content"` 

}

func (AdminGroupModel) TableName() string {
	return config.TablePre + "admin_group"
}

func AdminGroupList(list []AdminGroupModel) []AdminGroupModel {
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