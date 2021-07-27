package indexModel

import (
	"app/config"
)

type OpenAlipayAppModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Createtime      string   `json:"createtime"` 
	Appid      string   `json:"appid"` 
	Appkey      string   `json:"appkey"` 
	Status      uint   `json:"status"` 
	Appcert_path      string   `json:"appcert_path"` 
	Alicert_path      string   `json:"alicert_path"` 
	Rootcert_path      string   `json:"rootcert_path"` 
	Notify_url      string   `json:"notify_url"` 
	Return_url      string   `json:"return_url"` 
	Openlogin      uint   `json:"openlogin"` 
	Merchant_private_key      string   `json:"merchant_private_key"` 

}

func (OpenAlipayAppModel) TableName() string {
	return config.TablePre + "open_alipay_app"
}

func OpenAlipayAppList(list []OpenAlipayAppModel) []OpenAlipayAppModel {
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