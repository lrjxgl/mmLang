package indexModel

import (
	"app/config"
)

type SiteCityModel struct {
	Sc_id      uint   `gorm:"primaryKey";json:"sc_id"`
	Title      string   `json:"title"` 
	Cityid      uint   `json:"cityid"` 
	Lat      float64   `json:"lat"` 
	Lng      float64   `json:"lng"` 
	Orderindex      uint   `json:"orderindex"` 
	Status      uint   `json:"status"` 
	Pid      uint   `json:"pid"` 
	Siteid      uint   `json:"siteid"` 

}

func (SiteCityModel) TableName() string {
	return config.TablePre + "site_city"
}

func SiteCityList(list []SiteCityModel) []SiteCityModel {
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