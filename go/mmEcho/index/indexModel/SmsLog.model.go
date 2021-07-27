package indexModel

import (
	"app/config"
)

type SmsLogModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Telephone      string   `json:"telephone"` 
	Userid      uint   `json:"userid"` 
	Dateline      uint   `json:"dateline"` 
	Content      string   `json:"content"` 
	Status      uint   `json:"status"` 

}

func (SmsLogModel) TableName() string {
	return config.TablePre + "sms_log"
}

func SmsLogList(list []SmsLogModel) []SmsLogModel {
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