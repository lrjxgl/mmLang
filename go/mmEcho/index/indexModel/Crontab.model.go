package indexModel

import (
	"app/config"
)

type CrontabModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Url      string   `json:"url"` 
	Crontime      string   `json:"crontime"` 
	Status      uint   `json:"status"` 

}

func (CrontabModel) TableName() string {
	return config.TablePre + "crontab"
}

func CrontabList(list []CrontabModel) []CrontabModel {
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