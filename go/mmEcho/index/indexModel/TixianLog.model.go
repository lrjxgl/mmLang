package indexModel

import (
	"app/config"
)

type TixianLogModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Siteid      uint   `json:"siteid"` 
	Dateline      uint   `json:"dateline"` 
	Order_id      uint   `json:"order_id"` 
	Content      string   `json:"content"` 
	Admin_id      uint   `json:"admin_id"` 

}

func (TixianLogModel) TableName() string {
	return config.TablePre + "tixian_log"
}

func TixianLogList(list []TixianLogModel) []TixianLogModel {
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