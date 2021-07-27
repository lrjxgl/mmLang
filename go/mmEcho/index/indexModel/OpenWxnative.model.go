package indexModel

import (
	"app/config"
)

type OpenWxnativeModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Createtime      string   `json:"createtime"` 
	Appid      string   `json:"appid"` 
	Appkey      string   `json:"appkey"` 
	Status      uint   `json:"status"` 
	Sslcert_path      string   `json:"sslcert_path"` 
	Sslkey_path      string   `json:"sslkey_path"` 
	Mchid      string   `json:"mchid"` 
	Mchkey      string   `json:"mchkey"` 
	Openlogin      uint   `json:"openlogin"` 

}

func (OpenWxnativeModel) TableName() string {
	return config.TablePre + "open_wxnative"
}

func OpenWxnativeList(list []OpenWxnativeModel) []OpenWxnativeModel {
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