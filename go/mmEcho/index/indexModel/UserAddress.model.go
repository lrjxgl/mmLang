package indexModel

import (
	"app/config"
)

type UserAddressModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Address      string   `json:"address"` 
	Telephone      string   `json:"telephone"` 
	Truename      string   `json:"truename"` 
	Status      uint   `json:"status"` 
	Zip_code      string   `json:"zip_code"` 
	Lastid      uint   `json:"lastid"` 
	Province_id      uint   `json:"province_id"` 
	City_id      uint   `json:"city_id"` 
	Town_id      uint   `json:"town_id"` 
	Isdefault      uint   `json:"isdefault"` 
	Dateline      uint   `json:"dateline"` 
	Pct_address      string   `json:"pct_address"` 
	Lat      float64   `json:"lat"` 
	Lng      float64   `json:"lng"` 

}

func (UserAddressModel) TableName() string {
	return config.TablePre + "user_address"
}

func UserAddressList(list []UserAddressModel) []UserAddressModel {
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