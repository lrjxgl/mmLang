package indexModel

import (
	"app/config"
)

type PvModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Ssid      string   `json:"ssid"` 
	Url      string   `json:"url"` 
	M      string   `json:"m"` 
	A      string   `json:"a"` 
	Createtime      string   `json:"createtime"` 
	Createhour      uint   `json:"createhour"` 
	Createweek      uint   `json:"createweek"` 
	Ip      string   `json:"ip"` 
	Isphone      uint   `json:"isphone"` 
	User_agent      string   `json:"user_agent"` 

}

func (PvModel) TableName() string {
	return config.TablePre + "pv"
}

func PvList(list []PvModel) []PvModel {
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