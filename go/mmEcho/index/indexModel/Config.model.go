package indexModel

import (
	"app/config"
)

type ConfigModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Siteid      uint   `json:"siteid"` 
	Phone_on      uint   `json:"phone_on"` 
	Phone_reg      uint   `json:"phone_reg"` 
	Phone_user      string   `json:"phone_user"` 
	Phone_pwd      string   `json:"phone_pwd"` 
	Phone_num      string   `json:"phone_num"` 
	Smtphost      string   `json:"smtphost"` 
	Smtpport      uint   `json:"smtpport"` 
	Smtpemail      string   `json:"smtpemail"` 
	Smtpuser      string   `json:"smtpuser"` 
	Smtppwd      string   `json:"smtppwd"` 
	Water_on      uint   `json:"water_on"` 
	Water_type      uint   `json:"water_type"` 
	Water_pos      uint   `json:"water_pos"` 
	Water_str      string   `json:"water_str"` 
	Water_img      string   `json:"water_img"` 
	Water_size      uint   `json:"water_size"` 
	Rewrite_on      uint   `json:"rewrite_on"` 
	Spread_on      uint   `json:"spread_on"` 
	Spread_discount      uint   `json:"spread_discount"` 
	Spread_money      uint   `json:"spread_money"` 
	Spread_type      uint   `json:"spread_type"` 
	Grade_on      uint   `json:"grade_on"` 
	Hotsearch      string   `json:"hotsearch"` 
	Kf_qq      string   `json:"kf_qq"` 
	Kf_ww      string   `json:"kf_ww"` 
	Sms_qianming      string   `json:"sms_qianming"` 
	Sms_type      string   `json:"sms_type"` 
	S_bank_name      string   `json:"s_bank_name"` 
	S_bank_user      string   `json:"s_bank_user"` 
	S_bank_num      string   `json:"s_bank_num"` 

}

func (ConfigModel) TableName() string {
	return config.TablePre + "config"
}

func ConfigList(list []ConfigModel) []ConfigModel {
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