package indexModel

import (
	"app/config"
)

type OpenQqModel struct {
	Id      uint   `gorm:"primaryKey";json:"id"`
	Title      string   `json:"title"` 
	Apptype      string   `json:"apptype"` 
	Createtime      string   `json:"createtime"` 
	Appid      string   `json:"appid"` 
	Appkey      string   `json:"appkey"` 
	Status      uint   `json:"status"` 
	Openlogin      uint   `json:"openlogin"` 

}

func (OpenQqModel) TableName() string {
	return config.TablePre + "open_qq"
}

func OpenQqList(list []OpenQqModel) []OpenQqModel {
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