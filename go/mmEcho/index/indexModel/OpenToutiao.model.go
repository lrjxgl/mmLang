package indexModel

import (
	"app/config"
)

type OpenToutiaoModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Createtime      string   `json:"createtime"` 
	Appid      string   `json:"appid"` 
	Appkey      string   `json:"appkey"` 
	Status      uint   `json:"status"` 
	Merchant_private_key      string   `json:"merchant_private_key"` 
	Alipay_public_key      string   `json:"alipay_public_key"` 
	Notify_url      string   `json:"notify_url"` 
	Return_url      string   `json:"return_url"` 
	Openlogin      uint   `json:"openlogin"` 

}

func (OpenToutiaoModel) TableName() string {
	return config.TablePre + "open_toutiao"
}

func OpenToutiaoList(list []OpenToutiaoModel) []OpenToutiaoModel {
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