package indexModel

import (
	"app/config"
)

type CouponModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Money      float64   `json:"money"` 
	Amount      uint   `json:"amount"` 
	Etime      string   `json:"etime"` 
	Dateline      uint   `json:"dateline"` 
	Status      uint   `json:"status"` 
	Get_num      uint   `json:"get_num"` 
	Use_num      uint   `json:"use_num"` 
	Lower_money      float64   `json:"lower_money"` 
	Imgurl      string   `json:"imgurl"` 
	Limit_num      uint   `json:"limit_num"` 

}

func (CouponModel) TableName() string {
	return config.TablePre + "coupon"
}

func CouponList(list []CouponModel) []CouponModel {
	slen := len(list)
	if slen == 0 {
		return list
	}

	for i := 0; i < slen; i++ {
		m := list[i]
		m.Imgurl = config.Image_site(m.Imgurl)
		list[i] = m
	}
	return list
}