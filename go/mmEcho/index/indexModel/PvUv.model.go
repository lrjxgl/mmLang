package indexModel

import (
	"app/config"
)

type PvUvModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Ssid      string   `json:"ssid"` 
	Createtime      string   `json:"createtime"` 

}

func (PvUvModel) TableName() string {
	return config.TablePre + "pv_uv"
}

func PvUvList(list []PvUvModel) []PvUvModel {
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