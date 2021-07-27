package indexModel

import (
	"app/config"
)

type RechargeModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Siteid      uint   `json:"siteid"` 
	Userid      uint   `json:"userid"` 
	Money      float64   `json:"money"` 
	Pay_type      string   `json:"pay_type"` 
	Openid      string   `json:"openid"` 
	Pay_orderno      string   `json:"pay_orderno"` 
	Type_id      uint   `json:"type_id"` 
	Dateline      uint   `json:"dateline"` 
	Status      uint   `json:"status"` 
	Orderno      string   `json:"orderno"` 
	Orderinfo      string   `json:"orderinfo"` 
	Tablename      string   `json:"tablename"` 
	Orderdata      string   `json:"orderdata"` 

}

func (RechargeModel) TableName() string {
	return config.TablePre + "recharge"
}

func RechargeList(list []RechargeModel) []RechargeModel {
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