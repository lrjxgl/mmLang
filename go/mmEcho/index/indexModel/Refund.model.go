package indexModel

import (
	"app/config"
)

type RefundModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Shopid      uint   `json:"shopid"` 
	Paytype      string   `json:"paytype"` 
	Createtime      string   `json:"createtime"` 
	Money      float64   `json:"money"` 
	Recharge_orderno      string   `json:"recharge_orderno"` 
	Recharge_pay_orderno      string   `json:"recharge_pay_orderno"` 
	Recharge_id      uint   `json:"recharge_id"` 
	Content      string   `json:"content"` 
	Odata      string   `json:"odata"` 
	Status      uint   `json:"status"` 

}

func (RefundModel) TableName() string {
	return config.TablePre + "refund"
}

func RefundList(list []RefundModel) []RefundModel {
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