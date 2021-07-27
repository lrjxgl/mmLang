package indexModel

import (
	"app/config"
)

type ExpressFeeModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Fmoney      float64   `json:"fmoney"` 
	Amoney      float64   `json:"amoney"` 

}

func (ExpressFeeModel) TableName() string {
	return config.TablePre + "express_fee"
}

func ExpressFeeList(list []ExpressFeeModel) []ExpressFeeModel {
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