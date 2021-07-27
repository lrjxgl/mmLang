package indexModel

import (
	"app/config"
)

type CheckinIndexModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Type_id      uint   `json:"type_id"` 
	Userid      uint   `json:"userid"` 
	Dateline      uint   `json:"dateline"` 
	Grade      uint   `json:"grade"` 
	Last_day      uint   `json:"last_day"` 
	Last_ip      string   `json:"last_ip"` 
	All_times      uint   `json:"all_times"` 
	Is_valid      uint   `json:"is_valid"` 
	Gold      uint   `json:"gold"` 
	Days      uint   `json:"days"` 
	Siteid      uint   `json:"siteid"` 

}

func (CheckinIndexModel) TableName() string {
	return config.TablePre + "checkin_index"
}

func CheckinIndexList(list []CheckinIndexModel) []CheckinIndexModel {
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