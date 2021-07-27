package indexModel

import (
	"app/config"
)

type TixianModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Objectid      uint   `json:"objectid"` 
	K      string   `json:"k"` 
	Money      float64   `json:"money"` 
	Status      uint   `json:"status"` 
	Dateline      uint   `json:"dateline"` 
	Info      string   `json:"info"` 
	Reply      string   `json:"reply"` 
	Redateline      uint   `json:"redateline"` 
	Siteid      uint   `json:"siteid"` 
	Yhk_name      string   `json:"yhk_name"` 
	Yhk_haoma      string   `json:"yhk_haoma"` 
	Yhk_huming      string   `json:"yhk_huming"` 
	Telephone      string   `json:"telephone"` 
	Yhk_address      string   `json:"yhk_address"` 
	Adminid      uint   `json:"adminid"` 
	Paytype      string   `json:"paytype"` 
	Bankid      uint   `json:"bankid"` 

}

func (TixianModel) TableName() string {
	return config.TablePre + "tixian"
}

func TixianList(list []TixianModel) []TixianModel {
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