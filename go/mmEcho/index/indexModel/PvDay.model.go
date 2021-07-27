package indexModel

import (
	"app/config"
)

type PvDayModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Ssid      string   `json:"ssid"` 
	Createtime      string   `json:"createtime"` 
	Url      string   `json:"url"` 
	M      string   `json:"m"` 
	A      string   `json:"a"` 

}

func (PvDayModel) TableName() string {
	return config.TablePre + "pv_day"
}

func PvDayList(list []PvDayModel) []PvDayModel {
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