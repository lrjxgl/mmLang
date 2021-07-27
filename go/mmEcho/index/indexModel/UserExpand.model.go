package indexModel

import (
	"app/config"
)

type UserExpandModel struct {
	Uid      uint   `gorm:"primaryKey";json:"uid"`
	Content      string   `json:"content"` 

}

func (UserExpandModel) TableName() string {
	return config.TablePre + "user_expand"
}

func UserExpandList(list []UserExpandModel) []UserExpandModel {
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