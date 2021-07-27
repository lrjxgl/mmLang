package indexModel

import (
	"app/config"
)

type CouponUserModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Coupon_id      uint   `json:"coupon_id"` 
	Userid      uint   `json:"userid"` 
	Dateline      uint   `json:"dateline"` 
	Status      uint   `json:"status"` 
	Notice_time      uint   `json:"notice_time"` 

}

func (CouponUserModel) TableName() string {
	return config.TablePre + "coupon_user"
}

func CouponUserList(list []CouponUserModel) []CouponUserModel {
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