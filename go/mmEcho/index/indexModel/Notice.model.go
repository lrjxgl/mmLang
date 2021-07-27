package indexModel

import (
	"app/config"
)

type NoticeModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Dateline      uint   `json:"dateline"` 
	Type_id      uint   `json:"type_id"` 
	Status      uint   `json:"status"` 
	Userid      uint   `json:"userid"` 
	Content      string   `json:"content"` 
	Linkurl      string   `json:"linkurl"` 
	Title      string   `json:"title"` 

}

func (NoticeModel) TableName() string {
	return config.TablePre + "notice"
}

func NoticeList(list []NoticeModel) []NoticeModel {
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