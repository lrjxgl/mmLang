package indexModel

import (
	"app/config"
)

type PvIpModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Ip      string   `json:"ip"` 
	Createtime      string   `json:"createtime"` 

}

func (PvIpModel) TableName() string {
	return config.TablePre + "pv_ip"
}

func PvIpList(list []PvIpModel) []PvIpModel {
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