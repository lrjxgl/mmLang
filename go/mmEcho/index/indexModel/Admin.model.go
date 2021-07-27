package indexModel

import (
	"app/config"
)

type AdminModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Username      string   `json:"username"` 
	Password      string   `json:"password"` 
	Salt      uint   `json:"salt"` 
	Group_id      uint   `json:"group_id"` 
	Isfounder      uint   `json:"isfounder"` 

}

func (AdminModel) TableName() string {
	return config.TablePre + "admin"
}

func AdminList(list []AdminModel) []AdminModel {
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