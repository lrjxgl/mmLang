package indexModel

import (
	"app/config"
)

type SgpayModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Userid      uint   `json:"userid"` 
	Yhk_name      string   `json:"yhk_name"` 
	Yhk_num      string   `json:"yhk_num"` 
	Yhk_user      string   `json:"yhk_user"` 
	Money      float64   `json:"money"` 
	Createtime      string   `json:"createtime"` 
	Status      uint   `json:"status"` 

}

func (SgpayModel) TableName() string {
	return config.TablePre + "sgpay"
}

func SgpayList(list []SgpayModel) []SgpayModel {
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