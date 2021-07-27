package indexModel

import (
	"app/config"
)

type ExpressFeeCityModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Pid      uint   `json:"pid"` 
	Areaid      uint   `json:"areaid"` 

}

func (ExpressFeeCityModel) TableName() string {
	return config.TablePre + "express_fee_city"
}

func ExpressFeeCityList(list []ExpressFeeCityModel) []ExpressFeeCityModel {
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