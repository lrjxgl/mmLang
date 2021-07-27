package indexModel

import (
	"app/config"
)

type PermissionModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	M      string   `json:"m"` 
	Access      string   `json:"access"` 
	Title      string   `json:"title"` 

}

func (PermissionModel) TableName() string {
	return config.TablePre + "permission"
}

func PermissionList(list []PermissionModel) []PermissionModel {
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