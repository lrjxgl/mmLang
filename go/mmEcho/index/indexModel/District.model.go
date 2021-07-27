package indexModel

import (
	"app/config"
)

type DistrictModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Name      string   `json:"name"` 
	Level      uint   `json:"level"` 
	Usetype      uint   `json:"usetype"` 
	Upid      uint   `json:"upid"` 
	Displayorder      uint   `json:"displayorder"` 

}

func (DistrictModel) TableName() string {
	return config.TablePre + "district"
}

func DistrictList(list []DistrictModel) []DistrictModel {
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