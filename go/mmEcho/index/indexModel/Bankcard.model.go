package indexModel

import (
	"app/config"
)

type BankcardModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Status      uint   `json:"status"` 
	Yhk_name      string   `json:"yhk_name"` 
	Yhk_haoma      string   `json:"yhk_haoma"` 
	Yhk_huming      string   `json:"yhk_huming"` 
	Telephone      string   `json:"telephone"` 
	Yhk_address      string   `json:"yhk_address"` 
	Paytype      string   `json:"paytype"` 

}

func (BankcardModel) TableName() string {
	return config.TablePre + "bankcard"
}

func BankcardList(list []BankcardModel) []BankcardModel {
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