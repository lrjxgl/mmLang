package indexModel

import (
	"app/config"
)

type CityModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Name      string   `json:"name"` 
	Lat      float64   `json:"lat"` 
	Lng      float64   `json:"lng"` 
	Status      uint   `json:"status"` 
	Isrecommend      uint   `json:"isrecommend"` 

}

func (CityModel) TableName() string {
	return config.TablePre + "city"
}

func CityList(list []CityModel) []CityModel {
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