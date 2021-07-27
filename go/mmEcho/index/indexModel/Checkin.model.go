package indexModel

import (
	"app/config"
)

type CheckinModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Siteid      uint   `json:"siteid"` 
	Userid      uint   `json:"userid"` 
	Day      uint   `json:"day"` 
	Dateline      uint   `json:"dateline"` 
	Mood      uint   `json:"mood"` 
	Content      string   `json:"content"` 
	Type_id      uint   `json:"type_id"` 
	Ip      string   `json:"ip"` 

}

func (CheckinModel) TableName() string {
	return config.TablePre + "checkin"
}

func CheckinList(list []CheckinModel) []CheckinModel {
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